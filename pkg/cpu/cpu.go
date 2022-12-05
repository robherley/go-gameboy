package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/mmu"
)

type CPU struct {
	Registers *Registers
	MMU       *mmu.MMU
	Interrupt *Interrupt
	Halted    bool
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html
func New(cart *cartridge.Cartridge) *CPU {
	interrupt := &Interrupt{
		MasterEnabled: true,
		EI:            MASTER_SET_NONE,
		DI:            MASTER_SET_NONE,
		enable:        0x0,
		flag:          0x0,
	}

	return &CPU{
		Registers: RegistersForDMG(cart),
		MMU:       mmu.New(cart, interrupt),
		Interrupt: interrupt,
		Halted:    false,
	}
}

func (cpu *CPU) Fetch8() byte {
	defer func() {
		cpu.Registers.PC++
	}()

	return cpu.MMU.Read8(cpu.Registers.PC)
}

func (cpu *CPU) Fetch16() uint16 {
	defer func() {
		cpu.Registers.PC += 2
	}()

	return cpu.MMU.Read16(cpu.Registers.PC)
}

func (cpu *CPU) StackPush8(data byte) {
	cpu.Registers.SP--
	cpu.MMU.Write8(cpu.Registers.SP, data)
}

func (cpu *CPU) StackPush16(data uint16) {
	cpu.StackPush8(bits.Hi(data))
	cpu.StackPush8(bits.Lo(data))
}

func (cpu *CPU) StackPop8() byte {
	val := cpu.MMU.Read8(cpu.Registers.SP)
	cpu.Registers.SP++
	return val
}

func (cpu *CPU) StackPop16() uint16 {
	lo := cpu.StackPop8()
	hi := cpu.StackPop8()

	return bits.To16(hi, lo)
}

func (cpu *CPU) NextInstruction() (byte, *Instruction) {
	opcode := cpu.Fetch8()
	isCB := opcode == 0xCB
	if isCB {
		// cb-prefixed instructions have opcode on next fetch
		opcode = cpu.Fetch8()
	}

	instruction := InstructionFromOPCode(Byte(opcode), isCB)
	if instruction == nil {
		panic(errs.NewUnknownOPCodeError(opcode))
	}

	return opcode, instruction
}

// For a given operand, resolve the value at it's symbol (without dereferencing)
func (cpu *CPU) Resolve(operand *Operand) uint16 {
	return operand.Symbol.Resolve(cpu)
}

// Set8 will set 8-bit data based on the operand's symbol
// Only valid for Data and Register operands, will panic otherwise
func (cpu *CPU) Set8(operand *Operand, val byte) {
	switch symbol := (operand).Symbol.(type) {
	case Address:
		addr := cpu.Fetch16()
		cpu.MMU.Write8(addr, val)
	case Register:
		if operand.Deref {
			addr := cpu.Registers.Get(symbol)
			cpu.MMU.Write8(addr, byte(val))
		} else {
			cpu.Registers.Set(symbol, uint16(val))
		}
	default:
		panic(errs.NewInvalidSetOperandError(symbol))
	}
}

// Set16 will set 16-bit data based on the operand's symbol
// Only valid for Data and Register operands, will panic otherwise
func (cpu *CPU) Set16(operand *Operand, val uint16) {
	switch symbol := operand.Symbol.(type) {
	case Address:
		val := cpu.Fetch16()
		addr := cpu.MMU.Read16(val)
		cpu.MMU.Write16(addr, val)
	case Data:
		addr := operand.Symbol.Resolve(cpu)
		cpu.MMU.Write16(addr, val)
	case Register:
		if operand.Deref {
			addr := cpu.Registers.Get(symbol)
			cpu.MMU.Write16(addr, val)
		} else {
			cpu.Registers.Set(symbol, val)
		}
	default:
		panic(errs.NewInvalidSetOperandError(symbol))
	}
}

func (cpu *CPU) HandleInterrupts() {
	// check if master flag should be enabled this cycle
	if cpu.Interrupt.EI != MASTER_SET_NONE {
		if cpu.Interrupt.EI == MASTER_SET_NOW {
			cpu.Interrupt.MasterEnabled = true
		}
		cpu.Interrupt.EI--
	}

	// check if master flag should be disabled this cycle
	if cpu.Interrupt.DI != MASTER_SET_NONE {
		if cpu.Interrupt.DI == MASTER_SET_NOW {
			cpu.Interrupt.MasterEnabled = true
		}
		cpu.Interrupt.DI--
	}

	if !cpu.Interrupt.MasterEnabled {
		return
	}

	for _, interrupt := range interrupts {
		addr := interruptsToAddress[interrupt]

		// only handle if *both* interrupt enable and interrupt flag are set
		if cpu.Interrupt.IsFlagged(interrupt) && cpu.Interrupt.IsEnabled(interrupt) {
			// 1. push program counter to stack
			cpu.StackPush16(cpu.Registers.PC)
			// 2. set program counter to mapped interrupt address
			cpu.Registers.PC = addr
			// 3. clear interrupt flag
			cpu.Interrupt.flag &= ^byte(interrupt)
			// 4. unhalt cpu
			cpu.Halted = false
			// 5. disable all interrupts
			cpu.Interrupt.MasterEnabled = false

			// only handle one interrupt at a time, priority is based on bit order
			return
		}
	}
}

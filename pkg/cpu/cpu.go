package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/interrupt"
	"github.com/robherley/go-gameboy/pkg/mmu"
	"github.com/robherley/go-gameboy/pkg/timer"
)

type CPU struct {
	Registers *Registers
	MMU       *mmu.MMU
	Interrupt *interrupt.Interrupt
	Timer     *timer.Timer
	Halted    bool
	Ticks     uint64
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html
func New(cart *cartridge.Cartridge) *CPU {
	inter := interrupt.New()
	time := timer.New(func() {
		inter.Flag |= byte(interrupt.TIMER)
	})

	return &CPU{
		Registers: RegistersForDMG(cart),
		Timer:     time,
		MMU: mmu.New(
			cart,
			inter,
			time,
		),
		Interrupt: inter,
		Halted:    false,
	}
}

func (cpu *CPU) EmulateCycles(cycles int) {
	n := cycles * 4
	for i := 0; i < n; i++ {
		cpu.Ticks++
		cpu.Timer.Tick()
	}
}

func (cpu *CPU) Read8(address uint16) byte {
	val := cpu.MMU.Read8(address)
	cpu.EmulateCycles(1)
	return val
}

func (cpu *CPU) Read16(address uint16) uint16 {
	val := cpu.MMU.Read16(address)
	cpu.EmulateCycles(1)
	return val
}

func (cpu *CPU) Write8(address uint16, data byte) {
	if address == interrupt.FLAG_ADDRESS || address == timer.TAC_ADDRESS || address == timer.TIMA_ADDRESS {
		func() {}()
	}

	cpu.EmulateCycles(1)
	cpu.MMU.Write8(address, data)
}

func (cpu *CPU) Write16(address uint16, data uint16) {
	cpu.EmulateCycles(2)
	cpu.MMU.Write16(address, data)
}

func (cpu *CPU) Fetch8() byte {
	val := cpu.Read8(cpu.Registers.PC)
	cpu.Registers.PC++

	return val
}

func (cpu *CPU) Fetch16() uint16 {
	val := cpu.Read16(cpu.Registers.PC)
	cpu.Registers.PC += 2

	return val
}

func (cpu *CPU) StackPush8(data byte) {
	cpu.Registers.SP--
	cpu.Write8(cpu.Registers.SP, data)
}

func (cpu *CPU) StackPush16(data uint16) {
	cpu.StackPush8(bits.Hi(data))
	cpu.StackPush8(bits.Lo(data))
}

func (cpu *CPU) StackPop8() byte {
	val := cpu.Read8(cpu.Registers.SP)
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

	instruction := InstructionFromOPCode(opcode, isCB)
	if instruction == nil {
		panic(errs.NewUnknownOPCodeError(opcode))
	}

	return opcode, instruction
}

// For a given operand, resolve the value at it's symbol, automatically dereferencing the source if possible.
// The Address, Data and Byte symbols will never be dereferenced
func (cpu *CPU) Get(operand *Operand) uint16 {
	val := operand.Symbol.Resolve(cpu)
	switch symbol := operand.Symbol.(type) {
	case Register:
		if operand.Deref {
			return uint16(cpu.Read8(val))
		}
		return val
	case Address, Data, Byte:
		// for byte and data, there is never a dereference
		// but, for address, the context matters on how it's being used
		//  ie: see LDH A,(a8) vs. LDH (a8),A
		//  the former sets A to the dereference address, the latter sets the non-dereferenced address to A
		//  so for address derefs, we'll handle those case by case
		return val
	default:
		panic(errs.NewInvalidGetOperandError(symbol))
	}
}

// For a given operand, set the value automatically dereferencing the destination if possible.
// The Address and Data symbols will never be dereferenced here
func (cpu *CPU) Set(operand *Operand, val uint16) {
	writeFunc := func(addr uint16, data uint16) {
		if operand.Is16() && !operand.Deref { // 16 bit load
			cpu.Write16(addr, data)
		} else { // 8 bit load
			cpu.Write8(addr, byte(data))
		}
	}

	switch symbol := operand.Symbol.(type) {
	case Register:
		if operand.Deref {
			addr := cpu.Registers.Get(symbol)
			writeFunc(addr, val)
		} else {
			cpu.Registers.Set(symbol, val)
		}
	case Address, Data:
		addr := cpu.Get(operand)
		writeFunc(addr, val)
	default:
		panic(errs.NewInvalidSetOperandError(symbol))
	}
}

func (cpu *CPU) HandleInterrupts() {
	// check if master flag should be enabled this cycle
	if cpu.Interrupt.EI != interrupt.MASTER_SET_NONE {
		if cpu.Interrupt.EI == interrupt.MASTER_SET_NOW {
			cpu.Interrupt.MasterEnabled = true
		}
		cpu.Interrupt.EI--
	}

	// check if master flag should be disabled this cycle
	if cpu.Interrupt.DI != interrupt.MASTER_SET_NONE {
		if cpu.Interrupt.DI == interrupt.MASTER_SET_NOW {
			cpu.Interrupt.MasterEnabled = false
		}
		cpu.Interrupt.DI--
	}

	// exit early if master is not enabled
	if !cpu.Interrupt.MasterEnabled {
		return
	}

	for _, interruptType := range interrupt.Types {
		addr := interrupt.TypeToAddress[interruptType]

		// only handle if *both* interrupt enable and interrupt flag are set
		if cpu.Interrupt.Triggered(interruptType) {
			// 1. push program counter to stack
			cpu.StackPush16(cpu.Registers.PC)
			// 2. set program counter to mapped interrupt address
			cpu.Registers.PC = addr
			// 3. clear interrupt flag for type
			cpu.Interrupt.Flag &= ^byte(interruptType)
			// 4. unhalt cpu
			cpu.Halted = false
			// 5. disable all interrupts
			cpu.Interrupt.MasterEnabled = false

			// only handle one interrupt at a time, priority is based on bit order
			return
		}
	}
}

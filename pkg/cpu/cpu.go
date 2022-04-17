package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/instructions"
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
		Registers: TempRegisters(cart),
		MMU:       mmu.New(cart, interrupt),
		Interrupt: interrupt,
		Halted:    false,
	}
}

func (c *CPU) Fetch8() byte {
	defer func() {
		c.Registers.PC++
	}()

	return c.MMU.Read8(c.Registers.PC)
}

func (c *CPU) Fetch16() uint16 {
	defer func() {
		c.Registers.PC += 2
	}()

	return c.MMU.Read16(c.Registers.PC)
}

func (c *CPU) StackPush8(data byte) {
	c.Registers.SP--
	c.MMU.Write8(c.Registers.SP, data)
}

func (c *CPU) StackPush16(data uint16) {
	c.StackPush8(bits.Hi(data))
	c.StackPush8(bits.Lo(data))
}

func (c *CPU) StackPop8() byte {
	val := c.MMU.Read8(c.Registers.SP)
	c.Registers.SP++
	return val
}

func (c *CPU) StackPop16() uint16 {
	lo := c.StackPop8()
	hi := c.StackPop8()

	return bits.To16(hi, lo)
}

func (c *CPU) NextInstruction() (byte, *instructions.Instruction) {
	opcode := c.Fetch8()
	isCBPrexied := opcode == 0xCB
	if isCBPrexied {
		// cb-prefixed instructions have opcode on next fetch
		opcode = c.Fetch8()
	}

	instruction := instructions.FromOPCode(opcode, isCBPrexied)
	if instruction == nil {
		panic(errs.NewUnknownOPCodeError(opcode))
	}

	return opcode, instruction
}

func (c *CPU) HandleInterrupts() {
	if c.Interrupt.EI != MASTER_SET_NONE {
		if c.Interrupt.EI == MASTER_SET_NOW {
			c.Interrupt.MasterEnabled = true
		}
		c.Interrupt.EI--
	}

	if c.Interrupt.DI != MASTER_SET_NONE {
		if c.Interrupt.DI == MASTER_SET_NOW {
			c.Interrupt.MasterEnabled = true
		}
		c.Interrupt.DI--
	}

	if !c.Interrupt.MasterEnabled {
		return
	}

	for i := range interrupts {
		it := interrupts[i]
		addr := interruptsToAddress[it]

		// only handle if *both* interrupt enable and interrupt flag are set
		if c.Interrupt.IsFlagged(it) && c.Interrupt.IsEnabled(it) {
			// 1. push program counter to stack
			c.StackPush16(c.Registers.PC)
			// 2. set program counter to mapped interrupt address
			c.Registers.PC = addr
			// 3. clear interrupt flag
			c.Interrupt.flag &= ^byte(it)
			// 4. unhalt cpu
			c.Halted = false
			// 5. disable all interrupts
			c.Interrupt.MasterEnabled = false

			// only handle one interrupt at a time, priority is based on bit order
			return
		}
	}
}

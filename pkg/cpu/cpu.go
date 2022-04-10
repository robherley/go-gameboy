package cpu

import (
	"github.com/robherley/go-gameboy/pkg/cartridge"
)

type CPU struct {
	Registers *Registers
	Cartridge *cartridge.Cartridge
	RAM       *RAM

	IsHalted    bool
	IME         bool
	EnablingIME bool
	IE          byte
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html
func New(cart *cartridge.Cartridge) *CPU {
	return &CPU{
		Registers: TempRegisters(cart),
		Cartridge: cart,
		RAM:       &RAM{},
		IME:       true,
	}
}

func (c *CPU) Fetch8() byte {
	defer func() {
		c.Registers.PC++
	}()

	return c.Read8(c.Registers.PC)
}

func (c *CPU) Fetch16() uint16 {
	defer func() {
		c.Registers.PC += 2
	}()

	return c.Read16(c.Registers.PC)
}

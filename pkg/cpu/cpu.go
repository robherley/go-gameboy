package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/instructions"
)

type CPU struct {
	Registers *Registers
	Cartridge *cartridge.Cartridge
	RAM       *RAM

	IME bool
	IE  byte
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html
func New(cart *cartridge.Cartridge) *CPU {
	return &CPU{
		Registers: RegistersForDMG(cart),
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

// ValueOf resolves the given instruction symbol's value based on CPU state
func (c *CPU) ValueOf(operand *instructions.Operand) uint16 {
	switch symbol := (*operand).Symbol.(type) {
	case instructions.Data:
		if operand.Is16() {
			return c.Fetch16()
		} else {
			return uint16(c.Fetch8())
		}
	case instructions.Register:
		val := c.Registers.Get(symbol)
		if operand.Deref {
			val = uint16(c.Read8(val))
		}
		return val
	case byte:
		return uint16(symbol)
	default:
		panic(fmt.Errorf("invalid operand type: %T", symbol))
	}
}

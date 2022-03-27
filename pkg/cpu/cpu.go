package cpu

import (
	"github.com/robherley/go-dmg/pkg/cartridge"
)

type CPU struct {
	*Registers
	*MMU

	// Int Master Enabled: enables/disables interrupts
	IME bool
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html
func New(cart *cartridge.Cartridge) *CPU {
	return &CPU{
		MMU:       &MMU{cart},
		Registers: RegistersForDMG(cart),
		IME:       true,
	}
}

func (c *CPU) Fetch8() byte {
	defer func() {
		c.PC++
	}()

	return c.Read8(c.PC)
}

func (c *CPU) Fetch16() uint16 {
	defer func() {
		c.PC += 2
	}()

	return c.Read16(c.PC)
}

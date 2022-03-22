package cpu

import (
	"github.com/robherley/go-dmg/pkg/cartridge"
)

type CPU struct {
	*Registers
	MMU *MMU
}

func New(cart *cartridge.Cartridge) *CPU {
	return &CPU{
		MMU: &MMU{
			Cartridge: cart,
		},
		Registers: &Registers{
			PC: 0x100,
		},
	}
}

func (c *CPU) Fetch8() byte {
	defer func() {
		c.PC++
	}()

	return c.MMU.read8(c.PC)
}

func (c *CPU) Fetch16() uint16 {
	lo := c.Fetch8()
	hi := c.Fetch8()

	return uint16(lo) | (uint16(hi) << 8)
}

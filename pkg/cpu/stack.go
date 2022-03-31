package cpu

import "github.com/robherley/go-dmg/internal/bits"

func (c *CPU) StackPush8(data byte) {
	c.Registers.SP--
	c.Write8(c.Registers.SP, data)
}

func (c *CPU) StackPush16(data uint16) {
	c.StackPush8(bits.Hi(data))
	c.StackPush8(bits.Lo(data))
}

func (c *CPU) StackPop8() byte {
	val := c.Read8(c.Registers.SP)
	c.Registers.SP++
	return val
}

func (c *CPU) StackPop16() uint16 {
	lo := c.StackPop8()
	hi := c.StackPop8()

	return bits.To16(hi, lo)
}

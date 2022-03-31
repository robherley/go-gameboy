package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/instructions"
)

// Resolves the given symbol's value from the CPU
func Resolve(c *CPU, operand *instructions.Operand) uint16 {
	var val uint16

	switch symbol := (*operand).Symbol.(type) {
	case instructions.Data:
		val = resolveData(c, &symbol)
	case instructions.Register:
		val = c.Registers.Get(symbol)
		if operand.Deref {
			val = uint16(c.Read8(val))
		}
	case byte:
		val = uint16(symbol)
	default:
		panic(fmt.Errorf("invalid operand type: %T", symbol))
	}

	return val
}

func resolveData(c *CPU, data *instructions.Data) uint16 {
	switch *data {
	case instructions.D8:
		return uint16(c.Fetch8())
	case instructions.A8:
		return 0xFF00 | uint16(c.Fetch8())
	case instructions.D16, instructions.A16:
		return c.Fetch16()
	default:
		panic(fmt.Errorf("invalid data: %v", data))
	}
}

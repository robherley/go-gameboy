package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/instructions"
)

// Resolves the given symbol's value from the CPU
func ResolveValue[T bits.Uintish](c *CPU, operand *instructions.Operand) T {
	switch symbol := (*operand).Symbol.(type) {
	case instructions.Data:
		return resolveData[T](c, &symbol)
	case instructions.Register:
		return resolveRegister[T](c, &symbol)
	case byte:
		return T(symbol)
	default:
		panic(fmt.Errorf("invalid operand: %v", symbol))
	}
}

// Resolves if a condition is true or false based on CPU flags
func ResolveCondition(c *CPU, cond instructions.Condition) bool {
	switch cond {
	case instructions.NZ:
		return !c.GetFlag(FlagZ)
	case instructions.Z:
		return c.GetFlag(FlagZ)
	case instructions.NC:
		return !c.GetFlag(FlagC)
	case instructions.Ca:
		return c.GetFlag(FlagC)
	default:
		panic(fmt.Errorf("invalid condition: %v", cond))
	}
}

func resolveData[T bits.Uintish](c *CPU, data *instructions.Data) T {
	switch *data {
	case instructions.D8:
		return T(c.Fetch8())
	case instructions.D16:
		return T(c.Fetch16())
	case instructions.A8:
		return T(0xFF00 | uint16(c.Fetch8()))
	case instructions.A16:
		return T(c.Fetch16())
	default:
		panic(fmt.Errorf("invalid data: %v", data))
	}
}

func resolveRegister[T bits.Uintish](c *CPU, reg *instructions.Register) T {
	switch *reg {
	case instructions.A:
		return T(c.A)
	case instructions.B:
		return T(c.B)
	case instructions.C:
		return T(c.C)
	case instructions.D:
		return T(c.D)
	case instructions.E:
		return T(c.E)
	case instructions.F:
		return T(c.F)
	case instructions.H:
		return T(c.H)
	case instructions.L:
		return T(c.L)
	case instructions.SP:
		return T(c.SP)
	case instructions.PC:
		return T(c.PC)
	case instructions.AF:
		return T(c.GetAF())
	case instructions.BC:
		return T(c.GetBC())
	case instructions.DE:
		return T(c.GetDE())
	case instructions.HL:
		return T(c.GetHL())
	default:
		panic(fmt.Errorf("invalid register: %v", reg))
	}
}

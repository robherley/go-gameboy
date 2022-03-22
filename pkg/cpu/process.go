package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	instr "github.com/robherley/go-dmg/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

func (c *CPU) Process(in *instr.Instruction) byte {
	switch in.Mnemonic {
	case instr.NOP:
		return c.nop(in)
	case instr.JP:
		return c.jp(in)
	}

	panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
}

func (c *CPU) CheckCondition(cond instr.Condition) bool {
	switch cond {
	case instr.NZ:
		return !c.GetFlag(FlagZ)
	case instr.Z:
		return c.GetFlag(FlagZ)
	case instr.NC:
		return !c.GetFlag(FlagC)
	case instr.Ca:
		return c.GetFlag(FlagC)
	}

	panic(fmt.Errorf("invalid condition: %v", cond))
}

func (c *CPU) resolver(operand instr.Operand) uint16 {
	switch symbol := operand.Symbol.(type) {
	case instr.Data:
		return c.resolveData(symbol)
	case instr.Register:
		return c.resolveRegister(symbol)
	case byte:
		return bits.To16(symbol, 0)
	default:
		panic(fmt.Errorf("invalid operand: %v", symbol))
	}
}

func (c *CPU) resolveData(data instr.Data) uint16 {
	switch data {
	case instr.D8:
		return uint16(c.Fetch8())
	case instr.D16:
	case instr.A16:
		return c.Fetch16()
	case instr.A8:
		return 0xFF00 | uint16(c.Fetch8())
	}

	panic(fmt.Errorf("invalid data: %v", data))
}

func (c *CPU) resolveRegister(reg instr.Register) uint16 {
	switch reg {
	case instr.A:
		return uint16(c.A)
	case instr.B:
		return uint16(c.B)
	case instr.C:
		return uint16(c.C)
	case instr.D:
		return uint16(c.D)
	case instr.E:
		return uint16(c.E)
	case instr.F:
		return uint16(c.F)
	case instr.H:
		return uint16(c.H)
	case instr.L:
		return uint16(c.L)
	case instr.SP:
		return c.SP
	case instr.PC:
		return c.PC
	}

	panic(fmt.Errorf("invalid register: %v", reg))
}

func (c *CPU) nop(in *instr.Instruction) byte {
	return 4
}

func (c *CPU) jp(in *instr.Instruction) byte {
	// check if conditional jump
	if len(in.Operands) > 1 {
		cond, ok := in.Operands[0].Symbol.(instr.Condition)
		if !ok {
			panic(fmt.Errorf("JP must have <condition> <operand> for > 1 operand, got: %v", in.Operands[0].Symbol))
		}
		if c.CheckCondition(cond) {
			// condition passed, so jump to resolved value
			c.PC = c.resolver(in.Operands[1])
		}
	} else {
		// doesn't have condition, resolve the value
		c.PC = c.resolver(in.Operands[0])
	}

	return 4
}

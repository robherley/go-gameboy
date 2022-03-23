package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

// Process an instruction for a given mnemonic, returns number of ticks
func (c *CPU) Process(in *instructions.Instruction) byte {
	switch in.Mnemonic {
	case instructions.NOP:
		return c.NOP(in)
	case instructions.JP:
		return c.JP(in)
	case instructions.DI:
		return c.DI(in)
	case instructions.EI:
		return c.EI(in)
	case instructions.XOR:
		return c.XOR(in)
	default:
		panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
	}
}

// NOP: No operation
func (c *CPU) NOP(in *instructions.Instruction) byte {
	return 4
}

// JP: jump to address
func (c *CPU) JP(in *instructions.Instruction) byte {
	// check if conditional jump
	if len(in.Operands) > 1 {
		cond, ok := in.Operands[0].Symbol.(instructions.Condition)
		if !ok {
			panic(fmt.Errorf("JP must have <condition> <operand> for > 1 operand, got: %v", in.Operands[0].Symbol))
		}
		if ResolveCondition(c, cond) {
			// condition passed, so jump to resolved value
			c.PC = ResolveValue[uint16](c, &in.Operands[1])
		}
	} else {
		// doesn't have condition, resolve the value
		c.PC = ResolveValue[uint16](c, &in.Operands[0])
	}

	return 4
}

// DI: disables interrupts
func (c *CPU) DI(in *instructions.Instruction) byte {
	c.IME = false
	return 4
}

// EI: enables interrupts
func (c *CPU) EI(in *instructions.Instruction) byte {
	c.IME = true
	return 4
}

// XOR: logical exclusive OR with register A
func (c *CPU) XOR(in *instructions.Instruction) byte {
	if in.Operands[0].Size() == 2 {
		value := ResolveValue[uint16](c, &in.Operands[0])
		c.A ^= bits.Lo(value)
	} else {
		value := ResolveValue[byte](c, &in.Operands[0])
		c.A ^= value
	}

	// set zero flag if result is zero
	if c.A == 0 {
		c.SetFlag(FlagZ)
	}

	// reset other flags
	c.ClearFlag(FlagN)
	c.ClearFlag(FlagH)
	c.ClearFlag(FlagC)

	return 4
}

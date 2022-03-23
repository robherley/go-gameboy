package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

func (c *CPU) Process(in *instructions.Instruction) byte {
	switch in.Mnemonic {
	case instructions.NOP:
		return c.nop(in)
	case instructions.JP:
		return c.jp(in)
	}

	panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
}

func (c *CPU) nop(in *instructions.Instruction) byte {
	return 4
}

func (c *CPU) jp(in *instructions.Instruction) byte {
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

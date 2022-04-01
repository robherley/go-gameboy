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
	case instructions.LD:
		return c.LD(in)
	case instructions.LDH:
		return c.LDH(in)
	default:
		panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
	}
}

// NOP: No operation
func (c *CPU) NOP(in *instructions.Instruction) byte {
	return 4
}

// INC: increment register
func (c *CPU) INC(in *instructions.Instruction) byte {
	reg, _ := in.Operands[0].Symbol.(instructions.Register)

	val := c.Registers.Get(reg)
	c.Registers.Set(reg, val+1)

	return 4
}

// DEC: decrement register
func (c *CPU) DEC(in *instructions.Instruction) byte {
	reg, _ := in.Operands[0].Symbol.(instructions.Register)

	val := c.Registers.Get(reg)
	c.Registers.Set(reg, val-1)

	return 4
}

// JP: jump to address
func (c *CPU) JP(in *instructions.Instruction) byte {
	// check if conditional jump
	if len(in.Operands) > 1 {
		condition, _ := in.Operands[0].Symbol.(instructions.Condition)
		if c.Registers.IsCondition(condition) {
			// condition passed, so jump to resolved value
			c.Registers.PC = c.ValueOf(&in.Operands[1])
		}
	} else {
		// doesn't have condition, resolve the value
		c.Registers.PC = c.ValueOf(&in.Operands[0])
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
	value := c.ValueOf(&in.Operands[0])
	c.Registers.A ^= bits.Lo(value)

	// set zero flag if result is zero
	if c.Registers.A == 0 {
		c.Registers.SetFlag(FlagZ)
	}

	// reset other flags
	c.Registers.ClearFlag(FlagN)
	c.Registers.ClearFlag(FlagH)
	c.Registers.ClearFlag(FlagC)

	return 4
}

// LD: puts values from one operand into another
func (c *CPU) LD(in *instructions.Instruction) byte {
	numOps := len(in.Operands)

	// special case instruction for 0xF8
	if numOps == 3 {
		r8 := c.ValueOf(&in.Operands[2])

		// half carry (4 bits)
		setH := (c.Registers.SP&0xF)+(r8&0xF) > 0xF
		if setH {
			c.Registers.SetFlag(FlagH)
		}

		// carry (8 bits)
		setC := (c.Registers.SP&0xFF)+(r8&0xFF) > 0xFF
		if setC {
			c.Registers.SetFlag(FlagC)
		}

		// reset other flags
		c.Registers.ClearFlag(FlagZ)
		c.Registers.ClearFlag(FlagN)

		c.Registers.SetHL(c.Registers.SP + r8)

		return 4
	}

	dst := &in.Operands[0]
	src := &in.Operands[1]

	srcData := c.ValueOf(src)

	if dst.IsData() || dst.Deref {
		// if destination is data or dereference, we're writing to the address
		addr := c.ValueOf(dst)
		if src.Is16() {
			c.Write16(addr, srcData)
		} else {
			c.Write8(addr, byte(srcData))
		}
	} else if dst.IsRegister() {
		// if register to register, just write to the register
		c.Registers.Set(dst.Symbol.(instructions.Register), srcData)
	}

	// check if any HL+ or HL-, and adjust
	for i := range in.Operands {
		if in.Operands[i].Symbol != instructions.HL {
			continue
		}

		hl := c.Registers.Get(instructions.HL)

		if in.Operands[i].Inc {
			c.Registers.Set(instructions.HL, hl+1)
		}

		if in.Operands[i].Dec {
			c.Registers.Set(instructions.HL, hl-1)
		}
	}

	return 4
}

// LDH: loads/sets A from 8-bit signed data
func (c *CPU) LDH(in *instructions.Instruction) byte {
	first := in.Operands[0].Symbol
	second := in.Operands[1].Symbol

	if first == instructions.A && second == instructions.A8 {
		// LDH A (a8), alternate mnemonic is LD A,($FF00+a8)
		a8 := c.ValueOf(&in.Operands[1])
		c.Registers.A = c.Read8(0xFF00 | a8)

	} else if first == instructions.A8 && second == instructions.A {
		// LDH (a8) A, alternate mnemonic is LD ($FF00+a8),A
		a8 := c.ValueOf(&in.Operands[0])
		c.Write8(0xFF00|a8, c.Registers.A)
	}

	return 4
}

// POP: pops a two byte value off the stack
func (c *CPU) POP(in *instructions.Instruction) byte {
	val := c.StackPop16()

	// special case for AF, protect last nibble for flags
	if in.Operands[0].Symbol == instructions.AF {
		c.Registers.SetAF(val & 0xFFF0)
	} else {
		c.Registers.SetAF(val)
	}

	return 4
}

// PUSH: pushes a two byte value on the stacks
func (c *CPU) PUSH(in *instructions.Instruction) byte {
	val := c.ValueOf(&in.Operands[0])
	c.StackPush16(val)

	return 4
}

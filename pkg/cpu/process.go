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
	reg, ok := in.Operands[0].Symbol.(instructions.Register)
	if !ok {
		panic(fmt.Errorf("INC: must have register, got %s", in.Operands[0].Symbol))
	}

	val := GetRegister(c.Registers, reg)
	SetRegister(c.Registers, reg, val+1)

	return 4
}

// DEC: decrement register
func (c *CPU) DEC(in *instructions.Instruction) byte {
	reg, ok := in.Operands[0].Symbol.(instructions.Register)
	if !ok {
		panic(fmt.Errorf("DEC: must have register, got %s", in.Operands[0].Symbol))
	}

	val := GetRegister(c.Registers, reg)
	SetRegister(c.Registers, reg, val-1)

	return 4
}

// JP: jump to address
func (c *CPU) JP(in *instructions.Instruction) byte {
	// check if conditional jump
	if len(in.Operands) > 1 {
		cond, ok := in.Operands[0].Symbol.(instructions.Condition)
		if !ok {
			panic(fmt.Errorf("JP: must have <condition> <operand> for > 1 operand, got: %v", in.Operands[0].Symbol))
		}
		if c.IsCondition(cond) {
			// condition passed, so jump to resolved value
			c.PC = Resolve(c, &in.Operands[1])
		}
	} else {
		// doesn't have condition, resolve the value
		c.PC = Resolve(c, &in.Operands[0])
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
	value := Resolve(c, &in.Operands[0])
	c.A ^= bits.Lo(value)

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

// LD: puts values from one operand into another
func (c *CPU) LD(in *instructions.Instruction) byte {
	numOps := len(in.Operands)

	if (numOps != 2) && (numOps != 3) {
		panic(fmt.Errorf("LD: must have 2-3 operands, got: %d", numOps))
	}

	// special instruction for 0xF8
	if numOps == 3 {
		r8 := Resolve(c, &in.Operands[2])

		// half carry
		setH := (c.SP&0x000F)+(r8&0x000F) > 0x000F
		if setH {
			c.SetFlag(FlagH)
		}

		// carry
		setC := (c.SP&0x00FF)+(r8&0x00FF) > 0x00FF
		if setC {
			c.SetFlag(FlagC)
		}

		// reset other flags
		c.ClearFlag(FlagZ)
		c.ClearFlag(FlagN)
	}

	dst := &in.Operands[0]
	src := &in.Operands[1]

	srcData := Resolve(c, src)

	if dst.IsData() || dst.Deref {
		// if destination is data or dereference, we're writing to the address of the register
		addr := Resolve(c, dst)
		if src.Is16() {
			c.Write16(addr, srcData)
		} else {
			c.Write8(addr, byte(srcData))
		}
	} else if dst.IsRegister() {
		// if register to register, just copy the value
		SetRegister(c.Registers, dst.Symbol.(instructions.Register), srcData)
	} else {
		// unknown state
		panic(fmt.Errorf("LD: invalid symbol type, got: %T", dst.Symbol))
	}

	// check if any HL+ or HL-, and adjust
	for i := range in.Operands {
		if in.Operands[i].Symbol != instructions.HL {
			continue
		}

		hl := GetRegister(c.Registers, instructions.HL)

		if in.Operands[i].Inc {
			SetRegister(c.Registers, instructions.HL, hl+1)
		}

		if in.Operands[i].Dec {
			SetRegister(c.Registers, instructions.HL, hl-1)
		}
	}

	return 4
}

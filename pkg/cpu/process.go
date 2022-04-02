package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

// Process an instruction for a given mnemonic, returns number of ticks
func (c *CPU) Process(in *instructions.Instruction) byte {
	ops := in.Operands

	switch in.Mnemonic {
	case instructions.NOP:
		return c.NOP(ops)
	case instructions.JP:
		return c.JP(ops)
	case instructions.JR:
		return c.JR(ops)
	case instructions.CALL:
		return c.CALL(ops)
	case instructions.RET:
		return c.RET(ops)
	case instructions.RETI:
		return c.RETI(ops)
	case instructions.DI:
		return c.DI(ops)
	case instructions.EI:
		return c.EI(ops)
	case instructions.XOR:
		return c.XOR(ops)
	case instructions.LD:
		return c.LD(ops)
	case instructions.LDH:
		return c.LDH(ops)
	default:
		panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
	}
}

// jumper is a helper method for jump operations (JP, JR, CALL, etc)
func (c *CPU) jumper(ops []instructions.Operand, relative, pushPC bool) byte {
	var addr uint16

	// check if conditional jump
	if len(ops) > 1 {
		condition, _ := ops[0].Symbol.(instructions.Condition)
		if !c.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return 4
		}
		addr = c.ValueOf(&ops[1])
	} else {
		// doesn't have condition, resolve the value
		addr = c.ValueOf(&ops[0])
	}

	// push program counter, used for CALL
	if pushPC {
		c.StackPush16(c.Registers.PC)
	}

	// cast and add signed data for relative jump, used for JR
	if relative {
		rel := int8(addr & 0xFF)
		addr = c.Registers.PC + uint16(rel)
	}

	c.Registers.PC = addr

	return 4
}

// NOP: No operation
func (c *CPU) NOP(ops []instructions.Operand) byte {
	return 4
}

// INC: increment register
func (c *CPU) INC(ops []instructions.Operand) byte {
	reg, _ := ops[0].Symbol.(instructions.Register)

	val := c.Registers.Get(reg)
	c.Registers.Set(reg, val+1)

	return 4
}

// DEC: decrement register
func (c *CPU) DEC(ops []instructions.Operand) byte {
	reg, _ := ops[0].Symbol.(instructions.Register)

	val := c.Registers.Get(reg)
	c.Registers.Set(reg, val-1)

	return 4
}

// JP: jump to address (and check condition)
func (c *CPU) JP(ops []instructions.Operand) byte {
	// call jumper helper
	// not relative
	// don't push PC to stack
	return c.jumper(ops, false, false)
}

// JR: jump to relative address (and check condition)
func (c *CPU) JR(ops []instructions.Operand) byte {
	// call jumper helper
	// relative
	// don't push PC to stack
	return c.jumper(ops, true, false)
}

// CALL: push address of next instruction onto stack (and check condition)
func (c *CPU) CALL(ops []instructions.Operand) byte {
	// call jumper helper
	// not relative
	// push PC to stack
	return c.jumper(ops, false, true)
}

// RET: pop two bytes from stack & jump to that address (and check condition)
func (c *CPU) RET(ops []instructions.Operand) byte {
	addr := c.StackPop16()

	if len(ops) > 0 {
		condition, _ := ops[0].Symbol.(instructions.Condition)
		if !c.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return 4
		}
	}

	c.Registers.PC = addr

	return 4
}

// RETI: pop two bytes from stack & jump to that address then enable interrupts
func (c *CPU) RETI(ops []instructions.Operand) byte {
	// just a RET
	c.RET(ops)

	// and re-enable interrupts
	c.IME = true

	return 4
}

// DI: disables interrupts
func (c *CPU) DI(ops []instructions.Operand) byte {
	c.IME = false
	return 4
}

// EI: enables interrupts
func (c *CPU) EI(ops []instructions.Operand) byte {
	c.IME = true
	return 4
}

// XOR: logical exclusive OR with register A
func (c *CPU) XOR(ops []instructions.Operand) byte {
	value := c.ValueOf(&ops[0])
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
func (c *CPU) LD(ops []instructions.Operand) byte {
	numOps := len(ops)

	// special case instruction for 0xF8
	if numOps == 3 {
		r8 := c.ValueOf(&ops[2])

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

	dst := &ops[0]
	src := &ops[1]

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
	for i := range ops {
		if ops[i].Symbol != instructions.HL {
			continue
		}

		hl := c.Registers.Get(instructions.HL)

		if ops[i].Inc {
			c.Registers.Set(instructions.HL, hl+1)
		}

		if ops[i].Dec {
			c.Registers.Set(instructions.HL, hl-1)
		}
	}

	return 4
}

// LDH: loads/sets A from 8-bit signed data
func (c *CPU) LDH(ops []instructions.Operand) byte {
	first := ops[0].Symbol
	second := ops[1].Symbol

	if first == instructions.A && second == instructions.A8 {
		// LDH A (a8), alternate mnemonic is LD A,($FF00+a8)
		a8 := c.ValueOf(&ops[1])
		c.Registers.A = c.Read8(0xFF00 | a8)

	} else if first == instructions.A8 && second == instructions.A {
		// LDH (a8) A, alternate mnemonic is LD ($FF00+a8),A
		a8 := c.ValueOf(&ops[0])
		c.Write8(0xFF00|a8, c.Registers.A)
	}

	return 4
}

// POP: pops a two byte value off the stack
func (c *CPU) POP(ops []instructions.Operand) byte {
	val := c.StackPop16()

	// special case for AF, protect last nibble for flags
	if ops[0].Symbol == instructions.AF {
		c.Registers.SetAF(val & 0xFFF0)
	} else {
		reg, _ := ops[0].Symbol.(instructions.Register)
		c.Registers.Set(reg, val)
	}

	return 4
}

// PUSH: pushes a two byte value on the stacks
func (c *CPU) PUSH(ops []instructions.Operand) byte {
	val := c.ValueOf(&ops[0])
	c.StackPush16(val)

	return 4
}

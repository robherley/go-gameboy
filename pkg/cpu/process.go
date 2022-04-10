package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

// Process an instruction for a given mnemonic, returns number of ticks
func (c *CPU) Process(in *instructions.Instruction) error {
	var proc func([]instructions.Operand) error

	switch in.Mnemonic {
	case instructions.NOP:
		proc = c.NOP
	case instructions.JP:
		proc = c.JP
	case instructions.JR:
		proc = c.JR
	case instructions.CALL:
		proc = c.CALL
	case instructions.RST:
		proc = c.CALL
	case instructions.RET:
		proc = c.RET
	case instructions.RETI:
		proc = c.RETI
	case instructions.DI:
		proc = c.DI
	case instructions.EI:
		proc = c.EI
	case instructions.LD:
		proc = c.LD
	case instructions.LDH:
		proc = c.LDH
	case instructions.POP:
		proc = c.POP
	case instructions.PUSH:
		proc = c.PUSH
	case instructions.INC:
		proc = c.INC
	case instructions.DEC:
		proc = c.DEC
	case instructions.ADD:
		proc = c.ADD
	case instructions.ADC:
		proc = c.ADC
	case instructions.SUB:
		proc = c.SUB
	case instructions.SBC:
		proc = c.SBC
	case instructions.AND:
		proc = c.AND
	case instructions.OR:
		proc = c.OR
	case instructions.XOR:
		proc = c.XOR
	case instructions.CP:
		proc = c.CP
	case instructions.BIT:
		proc = c.BIT
	case instructions.RES:
		proc = c.RES
	case instructions.SET:
		proc = c.SET
	case instructions.RLC:
		proc = c.RLC
	case instructions.RL:
		proc = c.RL
	case instructions.RRC:
		proc = c.RRC
	case instructions.RR:
		proc = c.RR
	default:
		panic(fmt.Errorf("instruction not implemented: %s", in.Mnemonic))
	}

	return proc(in.Operands)
}

// valueOf: resolve the given instruction symbol's value based on CPU state
func (c *CPU) valueOf(operand *instructions.Operand) uint16 {
	switch symbol := (*operand).Symbol.(type) {
	case instructions.Data:
		if operand.Is16() {
			return c.Fetch16()
		}
		if operand.Symbol == instructions.R8 {
			// R8 is signed, convert it to int8 first
			return uint16(int8(c.Fetch8()))
		}
		return uint16(c.Fetch8())
	case instructions.Register:
		val := c.Registers.Get(symbol)
		if operand.Deref {
			val = uint16(c.Read8(val))
		}
		return val
	case byte:
		return uint16(symbol)
	default:
		panic(fmt.Errorf("invalid operand type: %T", symbol))
	}
}

// setRegOrAddr: set the register value or dereference the pointer address and set
func (c *CPU) setRegOrAddr(operand *instructions.Operand, value byte) {
	reg := operand.Symbol.(instructions.Register)
	regVal := c.Registers.Get(reg)

	// set address if deref
	if operand.Deref {
		c.Write8(regVal, value)
	} else { // otherwise set register
		c.Registers.Set(reg, uint16(value))
	}
}

// jumper: helper for jump operations (JP, JR, CALL, RST, etc)
func (c *CPU) jumper(mnemonic instructions.Mnemonic, ops []instructions.Operand) error {
	var addr uint16
	switch mnemonic {
	case instructions.RET, instructions.RETI:
		// RET/RETI gets jump value from stack
		addr = c.StackPop16()
	case instructions.JR:
		// relative jump, add to PC
		val := c.valueOf(&ops[len(ops)-1])
		addr = c.Registers.PC + val
	default:
		// otherwise get jump value from last operand
		addr = c.valueOf(&ops[len(ops)-1])
	}

	// check if has conditional
	// note: important to do this _after_ parameter read so PC is correct
	if len(ops) > 1 {
		condition := ops[0].Symbol.(instructions.Condition)
		if !c.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return nil
		}
	}

	// push program counter to stack, used for CALL & RST
	if mnemonic == instructions.CALL || mnemonic == instructions.RST {
		c.StackPush16(c.Registers.PC)
	}

	c.Registers.PC = addr

	return nil
}

// setRotateShiftFlags: helper to set flags for 0xCB rotate/shift func
func (c *CPU) setRotateShiftFlags(newVal byte, isCarry bool) {
	c.Registers.SetFlag(FlagZ, newVal == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, isCarry)
}

// NOP: No operation
func (c *CPU) NOP(ops []instructions.Operand) error {
	return nil
}

// INC: increment register
func (c *CPU) INC(ops []instructions.Operand) error {
	var result uint16

	if ops[0].Deref {
		// special case for instruction 0x34
		addr := c.valueOf(&ops[0])
		result = uint16(c.Read8(addr)) + 1
		c.Write8(addr, byte(result))
	} else {
		reg := ops[0].Symbol.(instructions.Register)
		result = c.Registers.Get(reg) + 1
		c.Registers.Set(reg, result)
	}

	if ops[0].Is16() && !ops[0].Deref {
		// if the parametes is 16 bit (and not a dereference)
		// then no flags get set (see instructions 0x03, 0x13, etc)
		return nil
	}

	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, (result&0xF) == 0)

	return nil
}

// DEC: decrement register
func (c *CPU) DEC(ops []instructions.Operand) error {
	var result uint16

	if ops[0].Deref {
		// special case for instruction 0x35
		addr := c.valueOf(&ops[0])
		result = uint16(c.Read8(addr)) - 1
		c.Write8(addr, byte(result))
	} else {
		reg := ops[0].Symbol.(instructions.Register)
		result = c.Registers.Get(reg) - 1
		c.Registers.Set(reg, result)
	}

	if ops[0].Is16() && !ops[0].Deref {
		// if the parametes is 16 bit (and not a dereference)
		// then no flags get set (see instructions 0x0B, 0x1B, etc)
		return nil
	}

	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (result&0xF) == 0xF)

	return nil
}

// JP: jump to address (and check condition)
func (c *CPU) JP(ops []instructions.Operand) error {
	return c.jumper(instructions.JP, ops)
}

// JR: jump to relative address (and check condition)
func (c *CPU) JR(ops []instructions.Operand) error {
	return c.jumper(instructions.JR, ops)
}

// CALL: push address of next instruction onto stack (and check condition)
func (c *CPU) CALL(ops []instructions.Operand) error {
	return c.jumper(instructions.CALL, ops)
}

// RST: push address on to stack, jump to n
func (c *CPU) RST(ops []instructions.Operand) error {
	return c.jumper(instructions.RST, ops)
}

// RET: pop two bytes from stack & jump to that address (and check condition)
func (c *CPU) RET(ops []instructions.Operand) error {
	return c.jumper(instructions.RET, ops)
}

// RETI: pop two bytes from stack & jump to that address then enable interrupts
func (c *CPU) RETI(ops []instructions.Operand) error {
	v := c.jumper(instructions.RETI, ops)
	c.IME = true
	return v
}

// DI: disables interrupts
func (c *CPU) DI(ops []instructions.Operand) error {
	c.IME = false
	return nil
}

// EI: enables interrupts
func (c *CPU) EI(ops []instructions.Operand) error {
	c.IME = true
	return nil
}

// LD: puts values from one operand into another
func (c *CPU) LD(ops []instructions.Operand) error {
	numOps := len(ops)

	// special case instruction for 0xF8
	if numOps == 3 {
		r8 := c.valueOf(&ops[2])

		// half carry (nibble)
		c.Registers.SetFlag(FlagH, (c.Registers.SP&0xF)+(r8&0xF) > 0xF)
		// carry (byte)
		c.Registers.SetFlag(FlagC, (c.Registers.SP&0xFF)+(r8&0xFF) > 0xFF)
		// reset other flags
		c.Registers.SetFlag(FlagZ, false)
		c.Registers.SetFlag(FlagN, false)

		c.Registers.SetHL(c.Registers.SP + r8)

		return nil
	}

	dst := &ops[0]
	src := &ops[1]

	srcData := c.valueOf(src)

	if dst.IsData() || dst.Deref {
		// if destination is data or dereference, we're writing to the address
		addr := c.valueOf(dst)
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

	return nil
}

// LDH: loads/sets A from 8-bit signed data
func (c *CPU) LDH(ops []instructions.Operand) error {
	first := ops[0].Symbol
	second := ops[1].Symbol

	if first == instructions.A && second == instructions.A8 {
		// LDH A (a8), alternate mnemonic is LD A,($FF00+a8)
		a8 := c.valueOf(&ops[1])
		c.Registers.A = c.Read8(0xFF00 | a8)

	} else if first == instructions.A8 && second == instructions.A {
		// LDH (a8) A, alternate mnemonic is LD ($FF00+a8),A
		a8 := c.valueOf(&ops[0])
		c.Write8(0xFF00|a8, c.Registers.A)
	}

	return nil
}

// POP: pops a two byte value off the stack
func (c *CPU) POP(ops []instructions.Operand) error {
	val := c.StackPop16()

	// special case for AF, protect last nibble for flags
	if ops[0].Symbol == instructions.AF {
		c.Registers.SetAF(val & 0xFFF0)
	} else {
		reg := ops[0].Symbol.(instructions.Register)
		c.Registers.Set(reg, val)
	}

	return nil
}

// PUSH: pushes a two byte value on the stack
func (c *CPU) PUSH(ops []instructions.Operand) error {
	val := c.valueOf(&ops[0])
	c.StackPush16(val)

	return nil
}

// ADD: Add a value to another value
func (c *CPU) ADD(ops []instructions.Operand) error {
	valA := c.valueOf(&ops[0])
	valB := c.valueOf(&ops[1])

	reg := ops[0].Symbol.(instructions.Register)
	sum := valA + valB

	c.Registers.Set(reg, sum)
	c.Registers.SetFlag(FlagN, false)

	// special case for 0xE8, adding n to stack pointer
	if ops[0].Symbol == instructions.SP {
		c.Registers.SetFlag(FlagZ, false)
		c.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF) > 0xF)
		c.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF) > 0xFF)
	} else if ops[0].Is16() { // 16bit add
		c.Registers.SetFlag(FlagH, (valA&0xFFF)+(valB&0xFFF) > 0xFFF)
		c.Registers.SetFlag(FlagH, (uint32(valA)&0xFFFF)+(uint32(valB)&0xFFFF) > 0xFFFF)
	} else { // 8bit add
		c.Registers.SetFlag(FlagZ, sum == 0)
		c.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF) > 0xF)
		c.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF) > 0xFF)
	}

	return nil
}

// ADC: Add with carry
func (c *CPU) ADC(ops []instructions.Operand) error {
	valA := c.valueOf(&ops[0])
	valB := c.valueOf(&ops[1])

	var carry uint16
	if c.Registers.GetFlag(FlagC) {
		carry = 1
	}

	sum := valA + valB + carry

	c.Registers.SetFlag(FlagZ, sum == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF)+carry > 0xF)
	c.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF)+carry > 0xFF)

	return nil
}

// SUB: Subtract a value from another value
func (c *CPU) SUB(ops []instructions.Operand) error {
	valA := uint16(c.Registers.A)
	valB := c.valueOf(&ops[0])
	diff := valA - valB

	c.Registers.Set(instructions.A, diff)
	c.Registers.SetFlag(FlagZ, diff == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	c.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))

	return nil
}

// SBC: Subtract a value (with carry flag) from another value
func (c *CPU) SBC(ops []instructions.Operand) error {
	valA := uint16(c.Registers.A)
	valB := c.valueOf(&ops[0])

	var carry uint16
	if c.Registers.GetFlag(FlagC) {
		carry = 1
	}

	diff := valA - valB + carry

	c.Registers.Set(instructions.A, diff)
	c.Registers.SetFlag(FlagZ, diff == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF)+carry)
	c.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF)+carry)

	return nil
}

// AND: logical AND with register A
func (c *CPU) AND(ops []instructions.Operand) error {
	val := c.valueOf(&ops[0])
	c.Registers.A &= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, true)
	c.Registers.SetFlag(FlagC, false)

	return nil
}

// OR: logical OR with register A
func (c *CPU) OR(ops []instructions.Operand) error {
	val := c.valueOf(&ops[0])
	c.Registers.A |= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)

	return nil
}

// XOR: logical exclusive OR with register A
func (c *CPU) XOR(ops []instructions.Operand) error {
	val := c.valueOf(&ops[0])
	c.Registers.A ^= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)

	return nil
}

// CP: compare with A (subtraction without setting result)
func (c *CPU) CP(ops []instructions.Operand) error {
	valA := uint16(c.Registers.A)
	valB := c.valueOf(&ops[0])
	diff := valA - valB

	c.Registers.Set(instructions.A, diff)
	c.Registers.SetFlag(FlagZ, diff == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	c.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))

	return nil
}

// BIT: (cb-prefixed) test bit in a register
func (c *CPU) BIT(ops []instructions.Operand) error {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	// will return t/f for nth bit in val
	isSet := bits.GetNBit(byte(val), byte(bit))

	c.Registers.SetFlag(FlagZ, !isSet)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, true)
	// carry flag not affected

	return nil
}

// RES: (cb-prefixed) reset bit b in a register
func (c *CPU) RES(ops []instructions.Operand) error {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	newVal := bits.ClearNBit(byte(val), byte(bit))

	// set address if deref
	if ops[1].Deref {
		c.Write8(val, byte(newVal))
	} else { // otherwise set register
		reg := ops[1].Symbol.(instructions.Register)
		c.Registers.Set(reg, uint16(newVal))
	}

	c.setRegOrAddr(&ops[1], newVal)
	// no flags affected

	return nil
}

// SET: (cb-prefixed) set bit b in a register
func (c *CPU) SET(ops []instructions.Operand) error {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	newVal := bits.SetNBit(byte(val), byte(bit))

	c.setRegOrAddr(&ops[1], newVal)
	// no flags affected

	return nil
}

// RLC: (cb-prefixed) rotate left, old bit 7 to carry
func (c *CPU) RLC(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 7)
	newVal := val << 1
	if isCarry {
		newVal |= 1
	}

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// RL: (cb-prefixed) rotate left through carry
func (c *CPU) RL(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 7)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	newVal := val << 1
	if isCarryFlagSet {
		newVal |= 1
	}

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// RRC: (cb-prefixed) rotate right, old bit 0 to carry
func (c *CPU) RRC(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 0)
	newVal := val >> 1
	if isCarry {
		newVal |= (1 << 7)
	}

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// RR: (cb-prefixed) rotate right through carry
func (c *CPU) RR(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 0)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	newVal := val >> 1
	if isCarryFlagSet {
		newVal |= (1 << 7)
	}

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// SLA: (cb-prefixed) shift left into carry. LSB of n set to 0
func (c *CPU) SLA(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	newVal := val << 1
	isCarry := bits.GetNBit(byte(val), 7)

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// SRA: (cb-prefixed) shift right into carry. MSB unaffected
func (c *CPU) SRA(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	newVal := (val >> 1) | (val & (1 << 7))
	isCarry := bits.GetNBit(byte(val), 0)

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// SRL: (cb-prefixed) shift right into carry. MSB set to 0
func (c *CPU) SRL(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	newVal := val >> 1
	isCarry := bits.GetNBit(byte(val), 0)

	c.setRegOrAddr(&ops[0], newVal)
	c.setRotateShiftFlags(newVal, isCarry)

	return nil
}

// SWAP: (cb-prefixed) swap upper & lower nibbles
func (c *CPU) SWAP(ops []instructions.Operand) error {
	val := byte(c.valueOf(&ops[0]))

	loNibs := val & 0x0F
	hiNibs := val & 0xF0

	newVal := (loNibs << 4) | (hiNibs >> 4)

	c.setRegOrAddr(&ops[0], newVal)
	c.Registers.SetFlag(FlagZ, newVal == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)

	return nil
}

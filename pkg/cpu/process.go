package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/instructions"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

// Process an instruction for a given mnemonic
func (c *CPU) Process(in *instructions.Instruction) {
	var proc func([]instructions.Operand)

	switch in.Mnemonic {
	case instructions.NOP:
		proc = c.NOP
	case instructions.STOP:
		proc = c.STOP
	case instructions.HALT:
		proc = c.HALT
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
	case instructions.RLCA:
		proc = c.RLCA
	case instructions.RRCA:
		proc = c.RRCA
	case instructions.RLA:
		proc = c.RLA
	case instructions.RRA:
		proc = c.RRA
	case instructions.DAA:
		proc = c.DAA
	case instructions.CPL:
		proc = c.CPL
	case instructions.CCF:
		proc = c.CCF
	case instructions.SCF:
		proc = c.SCF
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
		panic(errs.NewInvalidMnemonicError(string(in.Mnemonic)))
	}

	proc(in.Operands)
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
		panic(errs.NewInvalidOperandError(symbol))
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
func (c *CPU) jumper(mnemonic instructions.Mnemonic, ops []instructions.Operand) {
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
			return
		}
	}

	// push program counter to stack, used for CALL & RST
	if mnemonic == instructions.CALL || mnemonic == instructions.RST {
		c.StackPush16(c.Registers.PC)
	}

	c.Registers.PC = addr
}

// setRotateShiftFlags: helper to set flags for rotate/shift funcs
func (c *CPU) setRotateShiftFlags(result byte, isCarry bool) {
	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, isCarry)
}

// NOP: No operation
func (c *CPU) NOP(ops []instructions.Operand) {}

// STOP: halts CPU and display until button pressed
func (c *CPU) STOP(ops []instructions.Operand) {
	// TODO: figure out how this should actually behave
	panic(errs.NotImplementedError)
}

// HALT: power down CPU until an interrupt occurs
func (c *CPU) HALT(ops []instructions.Operand) {
	c.IsHalted = true
}

// INC: increment register
func (c *CPU) INC(ops []instructions.Operand) {
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
		return
	}

	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, (result&0xF) == 0)
}

// DEC: decrement register
func (c *CPU) DEC(ops []instructions.Operand) {
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
		return
	}

	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (result&0xF) == 0xF)
}

// JP: jump to address (and check condition)
func (c *CPU) JP(ops []instructions.Operand) {
	c.jumper(instructions.JP, ops)
}

// JR: jump to relative address (and check condition)
func (c *CPU) JR(ops []instructions.Operand) {
	c.jumper(instructions.JR, ops)
}

// CALL: push address of next instruction onto stack (and check condition)
func (c *CPU) CALL(ops []instructions.Operand) {
	c.jumper(instructions.CALL, ops)
}

// RST: push address on to stack, jump to n
func (c *CPU) RST(ops []instructions.Operand) {
	c.jumper(instructions.RST, ops)
}

// RET: pop two bytes from stack & jump to that address (and check condition)
func (c *CPU) RET(ops []instructions.Operand) {
	c.jumper(instructions.RET, ops)
}

// RETI: pop two bytes from stack & jump to that address then enable interrupts
func (c *CPU) RETI(ops []instructions.Operand) {
	c.jumper(instructions.RETI, ops)
	c.IME = true
}

// DI: disables interrupts
func (c *CPU) DI(ops []instructions.Operand) {
	c.IME = false
}

// EI: enables interrupts
func (c *CPU) EI(ops []instructions.Operand) {
	c.EnablingIME = true
}

// LD: puts values from one operand into another
func (c *CPU) LD(ops []instructions.Operand) {
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

		return
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
}

// LDH: loads/sets A from 8-bit signed data
func (c *CPU) LDH(ops []instructions.Operand) {
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
}

// POP: pops a two byte value off the stack
func (c *CPU) POP(ops []instructions.Operand) {
	val := c.StackPop16()

	// special case for AF, protect last nibble for flags
	if ops[0].Symbol == instructions.AF {
		c.Registers.SetAF(val & 0xFFF0)
	} else {
		reg := ops[0].Symbol.(instructions.Register)
		c.Registers.Set(reg, val)
	}
}

// PUSH: pushes a two byte value on the stack
func (c *CPU) PUSH(ops []instructions.Operand) {
	val := c.valueOf(&ops[0])
	c.StackPush16(val)
}

// ADD: Add a value to another value
func (c *CPU) ADD(ops []instructions.Operand) {
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
}

// ADC: Add with carry
func (c *CPU) ADC(ops []instructions.Operand) {
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
}

// SUB: Subtract a value from another value
func (c *CPU) SUB(ops []instructions.Operand) {
	valA := uint16(c.Registers.A)
	valB := c.valueOf(&ops[0])
	diff := valA - valB

	c.Registers.Set(instructions.A, diff)
	c.Registers.SetFlag(FlagZ, diff == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	c.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))
}

// SBC: Subtract a value (with carry flag) from another value
func (c *CPU) SBC(ops []instructions.Operand) {
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
}

// AND: logical AND with register A
func (c *CPU) AND(ops []instructions.Operand) {
	val := c.valueOf(&ops[0])
	c.Registers.A &= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, true)
	c.Registers.SetFlag(FlagC, false)
}

// OR: logical OR with register A
func (c *CPU) OR(ops []instructions.Operand) {
	val := c.valueOf(&ops[0])
	c.Registers.A |= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)
}

// XOR: logical exclusive OR with register A
func (c *CPU) XOR(ops []instructions.Operand) {
	val := c.valueOf(&ops[0])
	c.Registers.A ^= bits.Lo(val)

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)
}

// CP: compare with A (subtraction without setting result)
func (c *CPU) CP(ops []instructions.Operand) {
	valA := uint16(c.Registers.A)
	valB := c.valueOf(&ops[0])
	diff := valA - valB

	c.Registers.Set(instructions.A, diff)
	c.Registers.SetFlag(FlagZ, diff == 0)
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	c.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))
}

// RLCA: rotate A left. Old bit 7 to carry flag
func (c *CPU) RLCA(ops []instructions.Operand) {
	val := c.Registers.A

	isCarry := bits.GetNBit(val, 7)
	result := val << 1
	if isCarry {
		result |= 1
	}

	c.Registers.A = result
	c.setRotateShiftFlags(result, isCarry)
}

// RLCA: rotate A right. Old bit 0 to carry flag
func (c *CPU) RRCA(ops []instructions.Operand) {
	val := c.Registers.A

	isCarry := bits.GetNBit(val, 0)
	result := val >> 1
	if isCarry {
		result |= (1 << 7)
	}

	c.Registers.A = result
	c.setRotateShiftFlags(result, isCarry)
}

// RLA: rotate A left through carry flag
func (c *CPU) RLA(ops []instructions.Operand) {
	val := c.Registers.A

	isCarry := bits.GetNBit(val, 7)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	result := val << 1
	if isCarryFlagSet {
		result |= 1
	}

	c.Registers.A = result
	c.setRotateShiftFlags(result, isCarry)
}

// RLA: rotate A right through carry flag
func (c *CPU) RRA(ops []instructions.Operand) {
	val := c.Registers.A

	isCarry := bits.GetNBit(val, 0)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	result := val >> 1
	if isCarryFlagSet {
		result |= (1 << 7)
	}

	c.Registers.A = result
	c.setRotateShiftFlags(result, isCarry)
}

// DAA: decimal adjust register A
func (c *CPU) DAA(ops []instructions.Operand) {
	// this instruction is really weird, used for binary coded decimals (why is this a thing)
	// tldr: 16 - 10 = 6 👍
	// great explanation here: https://ehaskins.com/2018-01-30%20Z80%20DAA/
	// implementation borrowed from: https://github.com/mvdnes/rboy/blob/d14b07ce600cdba80754873a9ca185e1513f07c5/src/cpu.rs#L791-L807

	var adjust byte

	// if half carry
	if c.Registers.GetFlag(FlagH) {
		adjust |= 0x6
	}

	// if full carry
	if c.Registers.GetFlag(FlagC) {
		adjust |= 0x60
	}

	// if addition, there's some extra checks
	if !c.Registers.GetFlag(FlagN) {
		// ten's place
		if c.Registers.A&0xF > 0x9 {
			adjust |= 0x6
		}

		// hundredth's place
		if c.Registers.A > 0x99 {
			adjust |= 0x60
		}

		c.Registers.A += adjust
	} else {
		c.Registers.A -= adjust
	}

	c.Registers.SetFlag(FlagZ, c.Registers.A == 0)
	// flag N not affected
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, adjust >= 0x60)
}

// CPL: complement A register
func (c *CPU) CPL(ops []instructions.Operand) {
	c.Registers.A = ^c.Registers.A

	// flag Z not affected
	c.Registers.SetFlag(FlagN, true)
	c.Registers.SetFlag(FlagH, true)
	// flag C not affected
}

// SCF: set carry flag
func (c *CPU) SCF(ops []instructions.Operand) {
	// flag z not affected
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, true)
}

// CCF: complement carry flag
func (c *CPU) CCF(ops []instructions.Operand) {
	// flag z not affected
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, !c.Registers.GetFlag(FlagC))
}

// BIT: (cb-prefixed) test bit in a register
func (c *CPU) BIT(ops []instructions.Operand) {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	// will return t/f for nth bit in val
	isSet := bits.GetNBit(byte(val), byte(bit))

	c.Registers.SetFlag(FlagZ, !isSet)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, true)
	// carry flag not affected
}

// RES: (cb-prefixed) reset bit b in a register
func (c *CPU) RES(ops []instructions.Operand) {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	result := bits.ClearNBit(byte(val), byte(bit))

	// set address if deref
	if ops[1].Deref {
		c.Write8(val, byte(result))
	} else { // otherwise set register
		reg := ops[1].Symbol.(instructions.Register)
		c.Registers.Set(reg, uint16(result))
	}

	c.setRegOrAddr(&ops[1], result)
	// no flags affected
}

// SET: (cb-prefixed) set bit b in a register
func (c *CPU) SET(ops []instructions.Operand) {
	bit := c.valueOf(&ops[0])
	val := c.valueOf(&ops[1])

	result := bits.SetNBit(byte(val), byte(bit))

	c.setRegOrAddr(&ops[1], result)
	// no flags affected
}

// RLC: (cb-prefixed) rotate left, old bit 7 to carry
func (c *CPU) RLC(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 7)
	result := val << 1
	if isCarry {
		result |= 1
	}

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// RL: (cb-prefixed) rotate left through carry
func (c *CPU) RL(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 7)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	result := val << 1
	if isCarryFlagSet {
		result |= 1
	}

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// RRC: (cb-prefixed) rotate right, old bit 0 to carry
func (c *CPU) RRC(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 0)
	result := val >> 1
	if isCarry {
		result |= (1 << 7)
	}

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// RR: (cb-prefixed) rotate right through carry
func (c *CPU) RR(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	isCarry := bits.GetNBit(byte(val), 0)
	isCarryFlagSet := c.Registers.GetFlag(FlagC)
	result := val >> 1
	if isCarryFlagSet {
		result |= (1 << 7)
	}

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// SLA: (cb-prefixed) shift left into carry. LSB of n set to 0
func (c *CPU) SLA(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	result := val << 1
	isCarry := bits.GetNBit(byte(val), 7)

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// SRA: (cb-prefixed) shift right into carry. MSB unaffected
func (c *CPU) SRA(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	result := (val >> 1) | (val & (1 << 7))
	isCarry := bits.GetNBit(byte(val), 0)

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// SRL: (cb-prefixed) shift right into carry. MSB set to 0
func (c *CPU) SRL(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	result := val >> 1
	isCarry := bits.GetNBit(byte(val), 0)

	c.setRegOrAddr(&ops[0], result)
	c.setRotateShiftFlags(result, isCarry)
}

// SWAP: (cb-prefixed) swap upper & lower nibbles
func (c *CPU) SWAP(ops []instructions.Operand) {
	val := byte(c.valueOf(&ops[0]))

	loNibs := val & 0x0F
	hiNibs := val & 0xF0

	result := (loNibs << 4) | (hiNibs >> 4)

	c.setRegOrAddr(&ops[0], result)
	c.Registers.SetFlag(FlagZ, result == 0)
	c.Registers.SetFlag(FlagN, false)
	c.Registers.SetFlag(FlagH, false)
	c.Registers.SetFlag(FlagC, false)
}

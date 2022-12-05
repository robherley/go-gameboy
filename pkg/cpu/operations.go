package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

// https://gbdev.io/pandocs/CPU_Instruction_Set.html

type Operation func(*CPU, []Operand)

// NOP: No operation
func NOP(cpu *CPU, ops []Operand) {}

// STOP: halts CPU and display until button pressed, changes speed for GBC
func STOP(cpu *CPU, ops []Operand) {
	// TODO: figure out how this should actually behave
	// https://gbdev.io/pandocs/Reducing_Power_Consumption.html?highlight=stop#using-the-stop-instruction
	panic(errs.NewNotImplementedError())
}

// HALT: power down CPU until an interrupt occurs
func HALT(cpu *CPU, ops []Operand) {
	cpu.Halted = true
}

// INC: increment register
func INC(cpu *CPU, ops []Operand) {
	src := ops[0]
	val := cpu.Get(&src)

	result := val + 1

	cpu.Set(&src, result)

	if src.Is16() && !src.Deref {
		// if the parameter is 16 bit (and not a dereference)
		// then return and no flags get set (see instructions 0x03, 0x13, etc)
		return
	}

	cpu.Registers.SetFlag(FlagZ, (result&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, (result&0xF) == 0)
}

// DEC: decrement register
func DEC(cpu *CPU, ops []Operand) {
	src := ops[0]
	val := cpu.Get(&src)

	result := val - 1

	cpu.Set(&src, result)

	if src.Is16() && !src.Deref {
		// if the parameter is 16 bit (and not a dereference)
		// then return and no flags get set (see instructions 0x0B, 0x1B, etc)
		return
	}

	cpu.Registers.SetFlag(FlagZ, (result&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, (result&0xF) == 0xF)
}

// JP: jump to address (and check condition)
func JP(cpu *CPU, ops []Operand) {
	last := ops[len(ops)-1]
	addr := cpu.Get(&last)

	if condition, ok := ops[0].Symbol.(Condition); ok {
		if !cpu.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return
		}
	}

	cpu.Registers.PC = addr
}

// JR: jump to relative address (and check condition)
func JR(cpu *CPU, ops []Operand) {
	last := ops[len(ops)-1]
	addr := cpu.Get(&last)

	if condition, ok := ops[0].Symbol.(Condition); ok {
		if !cpu.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return
		}
	}

	cpu.Registers.PC += addr
}

// CALL: push address of next instruction onto stack (and check condition)
func CALL(cpu *CPU, ops []Operand) {
	last := ops[len(ops)-1]
	addr := cpu.Get(&last)

	if condition, ok := ops[0].Symbol.(Condition); ok {
		if !cpu.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return
		}
	}

	cpu.StackPush16(cpu.Registers.PC)
	cpu.Registers.PC = addr
}

// RST: push address on to stack, jump to n
func RST(cpu *CPU, ops []Operand) {
	addr := cpu.Get(&ops[0])

	cpu.StackPush16(cpu.Registers.PC)
	cpu.Registers.PC = addr
}

// RET: pop two bytes from stack & jump to that address (and check condition)
func RET(cpu *CPU, ops []Operand) {
	if len(ops) > 0 {
		condition, ok := ops[0].Symbol.(Condition)
		if ok && !cpu.Registers.IsCondition(condition) {
			// condition did not pass, so just return
			return
		}
	}

	addr := cpu.StackPop16()
	cpu.Registers.PC = addr
}

// RETI: pop two bytes from stack & jump to that address then enable interrupts immediately
func RETI(cpu *CPU, ops []Operand) {
	cpu.Interrupt.MasterEnabled = true
	addr := cpu.StackPop16()
	cpu.Registers.PC = addr
}

// DI: disables interrupts
func DI(cpu *CPU, ops []Operand) {
	cpu.Interrupt.DI = MASTER_SET_NEXT
}

// EI: enables interrupts (with delay)
func EI(cpu *CPU, ops []Operand) {
	cpu.Interrupt.EI = MASTER_SET_NEXT
}

// LD: puts values from one operand into another
func LD(cpu *CPU, ops []Operand) {
	numOps := len(ops)

	// special case instruction for 0xF8
	if numOps == 3 {
		r8 := ops[2].Symbol.Resolve(cpu)

		// half carry (nibble)
		cpu.Registers.SetFlag(FlagH, (cpu.Registers.SP&0xF)+(r8&0xF) > 0xF)
		// carry (byte)
		cpu.Registers.SetFlag(FlagC, (cpu.Registers.SP&0xFF)+(r8&0xFF) > 0xFF)
		// reset other flags
		cpu.Registers.SetFlag(FlagZ, false)
		cpu.Registers.SetFlag(FlagN, false)

		cpu.Registers.SetHL(cpu.Registers.SP + r8)

		return
	}

	dst := &ops[0]
	src := &ops[1]
	srcData := cpu.Get(src)

	if _, ok := src.Symbol.(Address); ok && src.Deref {
		// if the source is an address and is deref, get the value
		// this is for instruction 0xFA: LD A,(a16)
		srcData = uint16(cpu.MMU.Read8(srcData))
	}

	cpu.Set(dst, srcData)

	// check if any HL+ or HL-, and adjust
	for i := range ops {
		if ops[i].Symbol != HL {
			continue
		}

		hl := cpu.Registers.Get(HL)

		if ops[i].Inc {
			cpu.Registers.Set(HL, hl+1)
		}

		if ops[i].Dec {
			cpu.Registers.Set(HL, hl-1)
		}
	}
}

// LDH: loads/sets A from 8-bit signed data
func LDH(cpu *CPU, ops []Operand) {
	if ops[0].Symbol == A { // LDH A,(a8)
		addr := cpu.Get(&ops[1])
		val := cpu.MMU.Read8(addr)
		cpu.Registers.Set(A, uint16(val))
	} else { // LDH (a8),A
		addr := cpu.Get(&ops[0])
		val := cpu.Get(&ops[1])
		cpu.MMU.Write8(addr, byte(val))
	}
}

// POP: pops a two byte value off the stack
func POP(cpu *CPU, ops []Operand) {
	reg := ops[0].Symbol.(Register)
	val := cpu.StackPop16()
	cpu.Registers.Set(reg, val)
}

// PUSH: pushes a two byte value on the stack
func PUSH(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])
	cpu.StackPush16(val)
}

// ADD: Add a value to another value
func ADD(cpu *CPU, ops []Operand) {
	valA := cpu.Get(&ops[0])
	valB := cpu.Get(&ops[1])

	reg := ops[0].Symbol.(Register)
	sum := valA + valB

	cpu.Registers.Set(reg, sum)
	cpu.Registers.SetFlag(FlagN, false)

	// special case for 0xE8, adding n to stack pointer
	if ops[0].Symbol == SP {
		cpu.Registers.SetFlag(FlagZ, false)
		cpu.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF) > 0xF)
		cpu.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF) > 0xFF)
	} else if ops[0].Is16() { // 16bit add
		cpu.Registers.SetFlag(FlagH, (valA&0xFFF)+(valB&0xFFF) > 0xFFF)
		cpu.Registers.SetFlag(FlagC, (uint32(valA)&0xFFFF)+(uint32(valB)&0xFFFF) > 0xFFFF)
	} else { // 8bit add
		cpu.Registers.SetFlag(FlagZ, (sum&0xFF) == 0)
		cpu.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF) > 0xF)
		cpu.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF) > 0xFF)
	}
}

// ADC: Add with carry
func ADC(cpu *CPU, ops []Operand) {
	valA := cpu.Get(&ops[0])
	valB := cpu.Get(&ops[1])

	var carry uint16
	if cpu.Registers.GetFlag(FlagC) {
		carry = 1
	}

	sum := valA + valB + carry

	cpu.Registers.Set(A, sum)
	cpu.Registers.SetFlag(FlagZ, (sum&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, (valA&0xF)+(valB&0xF)+carry > 0xF)
	cpu.Registers.SetFlag(FlagC, (valA&0xFF)+(valB&0xFF)+carry > 0xFF)
}

// SUB: Subtract a value from another value
func SUB(cpu *CPU, ops []Operand) {
	valA := uint16(cpu.Registers.A)
	valB := cpu.Get(&ops[0])

	diff := valA - valB

	cpu.Registers.Set(A, diff)
	cpu.Registers.SetFlag(FlagZ, (diff&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	cpu.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))
}

// SBC: Subtract a value (with carry flag) from another value
func SBC(cpu *CPU, ops []Operand) {
	valA := uint16(cpu.Registers.A)
	valB := cpu.Get(&ops[0])

	var carry uint16
	if cpu.Registers.GetFlag(FlagC) {
		carry = 1
	}

	diff := valA - valB + carry

	cpu.Registers.Set(A, diff)
	cpu.Registers.SetFlag(FlagZ, (diff&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF)+carry)
	cpu.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF)+carry)
}

// AND: logical AND with register A
func AND(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	cpu.Registers.A &= bits.Lo(val)

	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, true)
	cpu.Registers.SetFlag(FlagC, false)
}

// OR: logical OR with register A
func OR(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	cpu.Registers.A |= bits.Lo(val)

	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
}

// XOR: logical exclusive OR with register A
func XOR(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	cpu.Registers.A ^= bits.Lo(val)

	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
}

// CP: compare with A (subtraction without setting result)
func CP(cpu *CPU, ops []Operand) {
	valA := uint16(cpu.Registers.A)
	valB := cpu.Get(&ops[0])

	diff := valA - valB

	cpu.Registers.SetFlag(FlagZ, (diff&0xFF) == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, (valA&0xF) < (valB&0xF))
	cpu.Registers.SetFlag(FlagC, (valA&0xFF) < (valB&0xFF))
}

// RLCA: rotate A left. Old bit 7 to carry flag
func RLCA(cpu *CPU, ops []Operand) {
	val := cpu.Registers.A

	isCarry := bits.GetNBit(val, 7)
	result := val << 1
	if isCarry {
		result |= 1
	}

	cpu.Registers.A = result
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RLCA: rotate A right. Old bit 0 to carry flag
func RRCA(cpu *CPU, ops []Operand) {
	val := cpu.Registers.A

	isCarry := bits.GetNBit(val, 0)
	result := val >> 1
	if isCarry {
		result |= (1 << 7)
	}

	cpu.Registers.A = result
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RLA: rotate A left through carry flag
func RLA(cpu *CPU, ops []Operand) {
	val := cpu.Registers.A

	isCarry := bits.GetNBit(val, 7)
	isCarryFlagSet := cpu.Registers.GetFlag(FlagC)
	result := val << 1
	if isCarryFlagSet {
		result |= 1
	}

	cpu.Registers.A = result
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RLA: rotate A right through carry flag
func RRA(cpu *CPU, ops []Operand) {
	val := cpu.Registers.A

	isCarry := bits.GetNBit(val, 0)
	isCarryFlagSet := cpu.Registers.GetFlag(FlagC)
	result := val >> 1
	if isCarryFlagSet {
		result |= (1 << 7)
	}

	cpu.Registers.A = result
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// DAA: decimal adjust register A
func DAA(cpu *CPU, ops []Operand) {
	// this instruction is really weird, used for binary coded decimals (why is this a thing)
	// tldr: 16 - 10 = 6 👍
	// great explanation here: https://ehaskins.com/2018-01-30%20Z80%20DAA/
	// implementation borrowed from: https://github.com/mvdnes/rboy/blob/d14b07ce600cdba80754873a9ca185e1513f07c5/src/cpu.rs#L791-L807

	var adjust byte

	// if half carry
	if cpu.Registers.GetFlag(FlagH) {
		adjust |= 0x6
	}

	// if full carry
	if cpu.Registers.GetFlag(FlagC) {
		adjust |= 0x60
	}

	// if addition, there's some extra checks
	if !cpu.Registers.GetFlag(FlagN) {
		// ten's place
		if cpu.Registers.A&0xF > 0x9 {
			adjust |= 0x6
		}

		// hundredth's place
		if cpu.Registers.A > 0x99 {
			adjust |= 0x60
		}

		cpu.Registers.A += adjust
	} else {
		cpu.Registers.A -= adjust
	}

	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	// flag N not affected
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, adjust >= 0x60)
}

// CPL: complement A register
func CPL(cpu *CPU, ops []Operand) {
	cpu.Registers.A = ^cpu.Registers.A

	// flag Z not affected
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, true)
	// flag C not affected
}

// SCF: set carry flag
func SCF(cpu *CPU, ops []Operand) {
	// flag z not affected
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, true)
}

// CCF: complement carry flag
func CCF(cpu *CPU, ops []Operand) {
	// flag z not affected
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, !cpu.Registers.GetFlag(FlagC))
}

// PREFIX: denotes the switch to prefixed instructions
func PREFIX(cpu *CPU, ops []Operand) {}

// BIT: (cb-prefixed) test bit in a register
func BIT(cpu *CPU, ops []Operand) {
	bit := cpu.Get(&ops[0])
	val := cpu.Get(&ops[1])

	// will return t/f for nth bit in val
	isSet := bits.GetNBit(byte(val), byte(bit))

	cpu.Registers.SetFlag(FlagZ, !isSet)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, true)
	// carry flag not affected
}

// RES: (cb-prefixed) reset bit b in a register
func RES(cpu *CPU, ops []Operand) {
	bit := cpu.Get(&ops[0])
	val := cpu.Get(&ops[1])

	result := bits.ClearNBit(byte(val), byte(bit))

	// set address if deref
	if ops[1].Deref {
		cpu.MMU.Write8(val, byte(result))
	} else { // otherwise set register
		reg := ops[1].Symbol.(Register)
		cpu.Registers.Set(reg, uint16(result))
	}

	// no flags affected
}

// SET: (cb-prefixed) set bit b in a register
func SET(cpu *CPU, ops []Operand) {
	bit := cpu.Get(&ops[0])
	val := cpu.Get(&ops[1])

	result := bits.SetNBit(byte(val), byte(bit))

	cpu.Set(&ops[1], uint16(result))
	// no flags affected
}

// RLC: (cb-prefixed) rotate left, old bit 7 to carry
func RLC(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	isCarry := bits.GetNBit(byte(val), 7)
	result := byte(val) << 1
	if isCarry {
		result |= 1
	}

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RL: (cb-prefixed) rotate left through carry
func RL(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	isCarry := bits.GetNBit(byte(val), 7)
	isCarryFlagSet := cpu.Registers.GetFlag(FlagC)
	result := byte(val) << 1
	if isCarryFlagSet {
		result |= 1
	}

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RRC: (cb-prefixed) rotate right, old bit 0 to carry
func RRC(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	isCarry := bits.GetNBit(byte(val), 0)
	result := byte(val) >> 1
	if isCarry {
		result |= (1 << 7)
	}

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// RR: (cb-prefixed) rotate right through carry
func RR(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	isCarry := bits.GetNBit(byte(val), 0)
	isCarryFlagSet := cpu.Registers.GetFlag(FlagC)
	result := byte(val) >> 1
	if isCarryFlagSet {
		result |= (1 << 7)
	}

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// SLA: (cb-prefixed) shift left into carry. LSB of n set to 0
func SLA(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	result := byte(val) << 1
	isCarry := bits.GetNBit(byte(val), 7)

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// SRA: (cb-prefixed) shift right into carry. MSB unaffected
func SRA(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	result := (byte(val) >> 1) | (byte(val) & (1 << 7))
	isCarry := bits.GetNBit(byte(val), 0)

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// SRL: (cb-prefixed) shift right into carry. MSB set to 0
func SRL(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	result := byte(val) >> 1
	isCarry := bits.GetNBit(byte(val), 0)

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetRotateAndShiftFlags(result, isCarry)
}

// SWAP: (cb-prefixed) swap upper & lower nibbles
func SWAP(cpu *CPU, ops []Operand) {
	val := cpu.Get(&ops[0])

	loNibs := byte(val) & 0x0F
	hiNibs := byte(val) & 0xF0

	result := (loNibs << 4) | (hiNibs >> 4)

	cpu.Set(&ops[0], uint16(result))
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
}

// ILLEGAL_D3: illegal D3 instruction
func ILLEGAL_D3(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xD3))
}

// ILLEGAL_DB: illegal DB instruction
func ILLEGAL_DB(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xDB))
}

// ILLEGAL_DD: illegal DD instruction
func ILLEGAL_DD(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xDD))
}

// ILLEGAL_E3: illegal E3 instruction
func ILLEGAL_E3(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xE3))
}

// ILLEGAL_E4: illegal E4 instruction
func ILLEGAL_E4(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xE4))
}

// ILLEGAL_EB: illegal EB instruction
func ILLEGAL_EB(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xEB))
}

// ILLEGAL_EC: illegal EC instruction
func ILLEGAL_EC(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xEC))
}

// ILLEGAL_ED: illegal ED instruction
func ILLEGAL_ED(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xED))
}

// ILLEGAL_F4: illegal F4 instruction
func ILLEGAL_F4(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xF4))
}

// ILLEGAL_FC: illegal FC instruction
func ILLEGAL_FC(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xFC))
}

// ILLEGAL_FD: illegal FD instruction
func ILLEGAL_FD(cpu *CPU, ops []Operand) {
	panic(errs.NewIllegalInstructionError(0xFD))
}

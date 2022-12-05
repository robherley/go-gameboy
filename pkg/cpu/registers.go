package cpu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html#registers

type Registers struct {
	A byte
	F byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte

	SP uint16
	PC uint16
}

// https://gbdev.io/pandocs/Power_Up_Sequence.html#cpu-registers
func RegistersForDMG(cart *cartridge.Cartridge) *Registers {
	r := &Registers{
		A:  0x01,
		F:  0x00,
		B:  0x00,
		C:  0x13,
		D:  0x00,
		E:  0xD8,
		H:  0x01,
		L:  0x4D,
		PC: 0x0100,
		SP: 0xFFFE,
	}

	r.SetFlag(FlagZ, true)

	// set carry and half carry if header checksum is != 0x00
	if cart.CalculateHeaderCheckSum() != 0 {
		r.SetFlag(FlagH, true)
		r.SetFlag(FlagC, true)
	}

	return r
}

func (registers *Registers) Set(reg Register, val uint16) {
	switch reg {
	case A:
		registers.A = byte(val)
	case B:
		registers.B = byte(val)
	case C:
		registers.C = byte(val)
	case D:
		registers.D = byte(val)
	case E:
		registers.E = byte(val)
	case F:
		registers.F = byte(val)
	case H:
		registers.H = byte(val)
	case L:
		registers.L = byte(val)
	case SP:
		registers.SP = val
	case PC:
		registers.PC = val
	case AF:
		registers.SetAF(val)
	case BC:
		registers.SetBC(val)
	case DE:
		registers.SetDE(val)
	case HL:
		registers.SetHL(val)
	default:
		panic(errs.NewInvalidOperandError(reg))
	}
}

func (registers *Registers) Get(reg Register) uint16 {
	switch reg {
	case A:
		return uint16(registers.A)
	case B:
		return uint16(registers.B)
	case C:
		return uint16(registers.C)
	case D:
		return uint16(registers.D)
	case E:
		return uint16(registers.E)
	case F:
		return uint16(registers.F)
	case H:
		return uint16(registers.H)
	case L:
		return uint16(registers.L)
	case SP:
		return registers.SP
	case PC:
		return registers.PC
	case AF:
		return registers.GetAF()
	case BC:
		return registers.GetBC()
	case DE:
		return registers.GetDE()
	case HL:
		return registers.GetHL()
	default:
		panic(errs.NewInvalidSymbolError(reg))
	}
}

/*
	Registers can be accessed as one 16 bit register OR separate 8 bit

	|16|Hi|Lo|
	|AF|A |* |
	|BC|B |C |
	|DE|D |E |
	|HL|H |L |
*/

func (r *Registers) GetAF() uint16 {
	return bits.To16(r.A, r.F)
}

func (r *Registers) SetAF(val uint16) {
	r.A = bits.Hi(val)
	r.F = bits.Lo(val & 0x00F0)
}

func (r *Registers) GetBC() uint16 {
	return bits.To16(r.B, r.C)
}

func (r *Registers) SetBC(val uint16) {
	r.B = bits.Hi(val)
	r.C = bits.Lo(val)
}

func (r *Registers) GetDE() uint16 {
	return bits.To16(r.D, r.E)
}

func (r *Registers) SetDE(val uint16) {
	r.D = bits.Hi(val)
	r.E = bits.Lo(val)
}

func (r *Registers) GetHL() uint16 {
	return bits.To16(r.H, r.L)
}

func (r *Registers) SetHL(value uint16) {
	r.H = bits.Hi(value)
	r.L = bits.Lo(value)
}

/*
	Flags

	The "F" register holds the CPU flags like so:

	|7|6|5|4|3|2|1|0|
	|Z|N|H|C|0|0|0|0|
*/

// Flag aliases for specific bits in register F
type Flag = byte

const (
	// Zero flag
	FlagZ Flag = 7
	// Subtraction flag
	FlagN Flag = 6
	// Half carry flag
	FlagH Flag = 5
	// Carry flag
	FlagC Flag = 4
)

// If a condition is true or false based on register flags
func (r *Registers) IsCondition(cond Condition) bool {
	switch cond {
	case NZ:
		return !r.GetFlag(FlagZ)
	case Z:
		return r.GetFlag(FlagZ)
	case NC:
		return !r.GetFlag(FlagC)
	case Ca:
		return r.GetFlag(FlagC)
	default:
		panic(errs.NewInvalidSymbolError(cond))
	}
}

func (r *Registers) GetFlag(f Flag) bool {
	return bits.GetNBit(r.F, f)
}

func (r *Registers) SetFlag(f Flag, set bool) {
	if set {
		r.F = bits.SetNBit(r.F, f)
	} else {
		r.F = bits.ClearNBit(r.F, f)
	}
}

// SetRotateAndShiftFlags: helper to set flags for rotate/shift funcs
func (r *Registers) SetRotateAndShiftFlags(setZero, isCarry bool) {
	r.SetFlag(FlagZ, setZero)
	r.SetFlag(FlagN, false)
	r.SetFlag(FlagH, false)
	r.SetFlag(FlagC, isCarry)
}

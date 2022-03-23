package cpu

import (
	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/cartridge"
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

	r.SetFlag(FlagZ)

	// set carry and half carry if header checksum is != 0x00
	if cart.CalculateHeaderCheckSum() != 0 {
		r.SetFlag(FlagH)
		r.SetFlag(FlagC)
	}

	return r
}

/*
	Registers can be accessed as one 16 bit register OR separate 8 bit

	|16|Hi|Lo|
	|AF|A |F |
	|BC|B |C |
	|DE|D |E |
	|HL|H |L |
*/

func (r *Registers) GetAF() uint16 {
	return bits.To16(r.A, r.F)
}

func (r *Registers) SetAF(value uint16) {
	r.A = bits.Hi(value)
	r.F = bits.Lo(value)
}

func (r *Registers) GetBC() uint16 {
	return bits.To16(r.B, r.C)
}

func (r *Registers) SetBC(value uint16) {
	r.B = bits.Hi(value)
	r.C = bits.Lo(value)
}

func (r *Registers) GetDE() uint16 {
	return bits.To16(r.D, r.E)
}

func (r *Registers) SetDE(value uint16) {
	r.D = bits.Hi(value)
	r.E = bits.Lo(value)
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

func (r *Registers) GetFlag(f Flag) bool {
	return bits.GetNBit(r.F, f)
}

func (r *Registers) SetFlag(f Flag) {
	r.F = bits.SetNBit(r.F, f)
}

func (r *Registers) ClearFlag(f Flag) {
	r.F = bits.ClearNBit(r.F, f)
}

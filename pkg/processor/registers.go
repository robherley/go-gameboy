package processor

// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html#registers

type Registers struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte
	H byte
	L byte

	SP uint16
	PC uint16
}

/*
	Registers can be accessed as one 16bit register OR separate 8 bit

	|16|Hi|Lo|
	|AF|A |- |
	|BC|B |C |
	|DE|D |E |
	|HL|H |L |
*/

func (r *Registers) GetAF() uint16 {
	return toU16(r.A, r.F)
}

func (r *Registers) SetAF(value uint16) {
	r.A = hi(value)
	r.F = lo(value)
}

func (r *Registers) GetBC() uint16 {
	return toU16(r.B, r.C)
}

func (r *Registers) SetBC(value uint16) {
	r.B = hi(value)
	r.C = lo(value)
}

func (r *Registers) GetDE() uint16 {
	return toU16(r.D, r.E)
}

func (r *Registers) SetDE(value uint16) {
	r.D = hi(value)
	r.E = lo(value)
}

func (r *Registers) GetHL() uint16 {
	return toU16(r.H, r.L)
}

func (r *Registers) SetHL(value uint16) {
	r.H = hi(value)
	r.L = lo(value)
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
	return getNBit(r.F, f)
}

func (r *Registers) SetFlag(f Flag) {
	r.F = setNBit(r.F, f)
}

func (r *Registers) ClearFlag(f Flag) {
	r.F = clearNBit(r.F, f)
}

package instructions

type Operand struct {
	// Symbol defines what kind of operand, and where to resolve the data
	Symbol any
	// Deref indicates a dereference of a pointer, ie: HL vs (HL)
	Deref bool
	// Inc indicates an increment, used for LDI, ie: HL vs HL+
	Inc bool
	// Dec indicates a decrement, used for LDD, ie: HL vs HL-
	Dec bool
}

// Registers
type Register string

const (
	// Single
	A Register = "A"
	B Register = "B"
	C Register = "C"
	D Register = "D"
	E Register = "E"
	F Register = "F"
	H Register = "H"
	L Register = "L"
	// Combined
	AF Register = "AF"
	BC Register = "BC"
	DE Register = "DE"
	HL Register = "HL"
	// Program Counter
	PC Register = "PC"
	// Stack Pointer
	SP Register = "SP"
)

type Data string

const (
	// Immediate 8-bit data
	D8 Data = "d8"
	// Immediate little-endian 16-bit data
	D16 Data = "d16"
	// 8-bit unsigned data, added to $FF00 in certain instructions
	A8 Data = "a8"
	// Little-endian 16-bit address
	A16 Data = "a16"
	// 8-bit signed data
	R8 Data = "r8"
)

type Condition string

const (
	// Not zero
	NZ Condition = "NZ"
	// Zero
	Z Condition = "Z"
	// Not carry
	NC Condition = "NC"
	// Carry
	Ca Condition = "C"
)

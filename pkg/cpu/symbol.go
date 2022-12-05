package cpu

import errs "github.com/robherley/go-gameboy/pkg/errors"

// Symbol represents a data type that can be used as an operand
type Symbol interface {
	// Resolve will return the value of the symbol from the CPU.
	// Note: it will not dereference the value held at symbol
	Resolve(c *CPU) uint16
}

type Register string

func (r Register) Resolve(cpu *CPU) uint16 {
	return cpu.Registers.Get(r)
}

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

func (d Data) Resolve(cpu *CPU) uint16 {
	switch d {
	case D8:
		return uint16(cpu.Fetch8())
	case D16:
		return cpu.Fetch16()
	case R8:
		val := cpu.Fetch8()
		// R8 is signed, convert it to int8 first
		return uint16(int8(val))
	default:
		panic(errs.NewInvalidSymbolError(d))
	}
}

const (
	// Immediate 8-bit data
	D8 Data = "d8"
	// Immediate little-endian 16-bit data
	D16 Data = "d16"
	// 8-bit signed data
	R8 Data = "r8"
)

type Address string

func (a Address) Resolve(cpu *CPU) uint16 {
	switch a {
	case A8:
		// alternative definition is ($FF00+a8)
		val := cpu.Fetch8()
		return 0xFF00 | uint16(val)
	case A16:
		return cpu.Fetch16()
	default:
		panic(errs.NewInvalidSymbolError(a))
	}
}

const (
	// 8-bit unsigned data, added to $FF00 in certain instructions
	A8 Address = "a8"
	// Little-endian 16-bit address
	A16 Address = "a16"
)

type Byte byte

func (b Byte) Resolve(cpu *CPU) uint16 {
	return uint16(b)
}

type Condition string

func (c Condition) Resolve(cpu *CPU) uint16 {
	if cpu.Registers.IsCondition(c) {
		return 1
	} else {
		return 0
	}
}

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

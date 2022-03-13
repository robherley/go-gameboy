package bits

// These are based on go's math/bits package. Unfortunately, they don't have methods for types < uint
// https://cs.opensource.google/go/go/+/refs/tags/go1.17.8:src/math/bits/bits.go
// TODO: swap these to generics when go 1.18 is stable

// Addition
// The sum will overflow if both top bits are set (x & y) or if one of them
// is (x | y), and a carry from the lower place happened. If such a carry
// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).

func Add8(x, y, carry byte) (byte, byte) {
	sum := x + y + carry
	carryOut := ((x & y) | ((x | y) &^ sum)) >> 7
	return sum, carryOut
}

func Add16(x, y, carry uint16) (uint16, uint16) {
	sum := x + y + carry
	carryOut := ((x & y) | ((x | y) &^ sum)) >> 15
	return sum, carryOut
}

// Subtraction
// The difference will underflow if the top bit of x is not set and the top
// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
// from the lower place happens. If that borrow happens, the result will be
// 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).

func Sub8(x, y, borrow byte) (byte, byte) {
	diff := x - y - borrow
	borrowOut := ((^x & y) | (^(x ^ y) & diff)) >> 7
	return diff, borrowOut
}

func Sub16(x, y, borrow uint16) (uint16, uint16) {
	diff := x - y - borrow
	borrowOut := ((^x & y) | (^(x ^ y) & diff)) >> 15
	return diff, borrowOut
}

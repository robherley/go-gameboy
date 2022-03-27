package bits

import "unsafe"

// These are based on go's math/bits package. Unfortunately, they don't have methods for types < uint
// https://cs.opensource.google/go/go/+/refs/tags/go1.17.8:src/math/bits/bits.go

type Uintish interface {
	// only need these two for gb emulation, should be extenable for any uint type
	~uint8 | ~uint16
}

// Addition
// The sum will overflow if both top bits are set (x & y) or if one of them
// is (x | y), and a carry from the lower place happened. If such a carry
// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
func Add[T Uintish](x, y T, carry byte) (T, byte) {
	sum := x + y + T(carry)

	// calculate size of union uint type
	numBits := unsafe.Sizeof(x) * 8

	carryOut := ((x & y) | ((x | y) &^ sum)) >> (numBits - 1)
	return sum, byte(carryOut)
}

// Subtraction
// The difference will underflow if the top bit of x is not set and the top
// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
// from the lower place happens. If that borrow happens, the result will be
// 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).
func Sub[T Uintish](x, y T, borrow byte) (T, byte) {
	diff := x - y - T(borrow)

	// calculate size of union uint type
	numBits := unsafe.Sizeof(x) * 8

	borrowOut := ((^x & y) | (^(x ^ y) & diff)) >> (numBits - 1)
	return diff, byte(borrowOut)
}

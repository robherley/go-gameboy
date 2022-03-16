package bits_test

import (
	"fmt"
	"testing"

	"github.com/robherley/go-dmg/internal/bits"
)

const (
	max8  = uint8((1 << 8) - 1)
	max16 = uint16((1 << 16) - 1)
)

type input[T bits.Uintish] struct {
	x             T
	y             T
	carryOrBorrow byte
}

type expected[T bits.Uintish] struct {
	result           T
	carryOrBorrowOut byte
}

func assertResult[T bits.Uintish](t *testing.T, result T, carryOutorBorrow byte, exp expected[T]) {
	if result != exp.result {
		t.Fatalf("result mismatch, got: %d want: %d", result, exp.result)
	}
	if carryOutorBorrow != exp.carryOrBorrowOut {
		t.Fatalf("carryOutorBorrow mismatch, got: %d want: %d", carryOutorBorrow, exp.carryOrBorrowOut)
	}
}

func TestAdd(t *testing.T) {
	u8tests := []struct {
		input[uint8]
		expected[uint8]
	}{
		{input[uint8]{0, 0, 0}, expected[uint8]{0, 0}},
		{input[uint8]{1, 1, 0}, expected[uint8]{2, 0}},
		{input[uint8]{1, 1, 1}, expected[uint8]{3, 0}},
		{input[uint8]{max8, 1, 0}, expected[uint8]{0, 1}},
		{input[uint8]{max8, max8, 0}, expected[uint8]{max8 - 1, 1}},
		{input[uint8]{max8, max8, 1}, expected[uint8]{max8, 1}},
	}

	for _, tc := range u8tests {
		t.Run(fmt.Sprintf("(uint8) Add=%+v", tc.input), func(tt *testing.T) {
			sum, carryOut := bits.Add(tc.x, tc.y, tc.carryOrBorrow)
			assertResult(tt, sum, carryOut, tc.expected)
		})
	}

	u16tests := []struct {
		input[uint16]
		expected[uint16]
	}{
		{input[uint16]{0, 0, 0}, expected[uint16]{0, 0}},
		{input[uint16]{1, 1, 0}, expected[uint16]{2, 0}},
		{input[uint16]{1, 1, 1}, expected[uint16]{3, 0}},
		{input[uint16]{max16, 1, 0}, expected[uint16]{0, 1}},
		{input[uint16]{max16, max16, 0}, expected[uint16]{max16 - 1, 1}},
		{input[uint16]{max16, max16, 1}, expected[uint16]{max16, 1}},
	}

	for _, tc := range u16tests {
		t.Run(fmt.Sprintf("(uint16) Add=%+v", tc.input), func(tt *testing.T) {
			sum, carryOut := bits.Add(tc.x, tc.y, tc.carryOrBorrow)
			assertResult(tt, sum, carryOut, tc.expected)
		})
	}
}

func TestSub(t *testing.T) {
	u8tests := []struct {
		input[uint8]
		expected[uint8]
	}{
		{input[uint8]{0, 0, 0}, expected[uint8]{0, 0}},
		{input[uint8]{1, 1, 0}, expected[uint8]{0, 0}},
		{input[uint8]{1, 1, 1}, expected[uint8]{max8, 1}},
		{input[uint8]{max8, 1, 0}, expected[uint8]{max8 - 1, 0}},
		{input[uint8]{max8, max8, 0}, expected[uint8]{0, 0}},
		{input[uint8]{max8, max8, 1}, expected[uint8]{max8, 1}},
	}

	for _, tc := range u8tests {
		t.Run(fmt.Sprintf("(uint8) Sub=%+v", tc.input), func(tt *testing.T) {
			sum, carryOut := bits.Sub(tc.x, tc.y, tc.carryOrBorrow)
			assertResult(tt, sum, carryOut, tc.expected)
		})
	}

	u16tests := []struct {
		input[uint16]
		expected[uint16]
	}{
		{input[uint16]{0, 0, 0}, expected[uint16]{0, 0}},
		{input[uint16]{1, 1, 0}, expected[uint16]{0, 0}},
		{input[uint16]{1, 1, 1}, expected[uint16]{max16, 1}},
		{input[uint16]{max16, 1, 0}, expected[uint16]{max16 - 1, 0}},
		{input[uint16]{max16, max16, 0}, expected[uint16]{0, 0}},
		{input[uint16]{max16, max16, 1}, expected[uint16]{max16, 1}},
	}

	for _, tc := range u16tests {
		t.Run(fmt.Sprintf("(uint16) Sub=%+v", tc.input), func(tt *testing.T) {
			sum, carryOut := bits.Sub(tc.x, tc.y, tc.carryOrBorrow)
			assertResult(tt, sum, carryOut, tc.expected)
		})
	}
}

package bits_test

import (
	"fmt"
	"testing"

	"github.com/robherley/go-gameboy/internal/bits"
)

func TestLo(t *testing.T) {
	cases := []struct {
		input    uint16
		expected byte
	}{
		{0xFFFF, 0xFF},
		{0x00FF, 0xFF},
		{0xFF00, 0x00},
		{0x0000, 0x00},
		{0xDEAD, 0xAD},
		{0xBEEF, 0xEF},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Lo=0x%x", tc.input), func(tt *testing.T) {
			actual := bits.Lo(tc.input)
			if actual != tc.expected {
				t.Fatalf("got: %x, want: %x", actual, tc.expected)
			}
		})
	}
}

func TestHi(t *testing.T) {
	cases := []struct {
		input    uint16
		expected byte
	}{
		{0xFFFF, 0xFF},
		{0x00FF, 0x00},
		{0xFF00, 0xFF},
		{0x0000, 0x00},
		{0xDEAD, 0xDE},
		{0xBEEF, 0xBE},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Hi=0x%x", tc.input), func(tt *testing.T) {
			actual := bits.Hi(tc.input)
			if actual != tc.expected {
				t.Fatalf("got: %x, want: %x", actual, tc.expected)
			}
		})
	}
}

func TestTo16(t *testing.T) {
	cases := []struct {
		hi       byte
		lo       byte
		expected uint16
	}{
		{0xFF, 0xFF, 0xFFFF},
		{0x00, 0xFF, 0x00FF},
		{0xFF, 0x00, 0xFF00},
		{0xDE, 0xAD, 0xDEAD},
		{0xBE, 0xEF, 0xBEEF},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("To16=0x%x,0x%x", tc.hi, tc.lo), func(tt *testing.T) {
			actual := bits.To16(tc.hi, tc.lo)
			if actual != tc.expected {
				t.Fatalf("got: %x, want: %x", actual, tc.expected)
			}
		})
	}
}

func TestGetNBit(t *testing.T) {
	cases := []struct {
		num      byte
		bit      byte
		expected bool
	}{
		{0b00000000, 0, false},
		{0b00000000, 7, false},
		{0b00000001, 0, true},
		{0b10000000, 7, true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("GetNBit=%d,0b%b", tc.num, tc.bit), func(tt *testing.T) {
			actual := bits.GetNBit(tc.num, tc.bit)
			if actual != tc.expected {
				t.Fatalf("got: %v, want: %v", actual, tc.expected)
			}
		})
	}
}

func TestSetNBit(t *testing.T) {
	cases := []struct {
		num      byte
		bit      byte
		expected byte
	}{
		{0b00000000, 0, 0b00000001},
		{0b00000000, 7, 0b10000000},
		{0b11111111, 0, 0b11111111},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("SetNBit=%d,0b%b", tc.num, tc.bit), func(tt *testing.T) {
			actual := bits.SetNBit(tc.num, tc.bit)
			if actual != tc.expected {
				t.Fatalf("got: 0b%08b, want: 0b%08b", actual, tc.expected)
			}
		})
	}
}

func TestClearNBit(t *testing.T) {
	cases := []struct {
		num      byte
		bit      byte
		expected byte
	}{
		{0b11111111, 0, 0b11111110},
		{0b11111111, 7, 0b01111111},
		{0b00000000, 0, 0b00000000},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("ClearNBit=%d,0b%b", tc.num, tc.bit), func(tt *testing.T) {
			actual := bits.ClearNBit(tc.num, tc.bit)
			if actual != tc.expected {
				t.Fatalf("got: 0b%08b, want: 0b%08b", actual, tc.expected)
			}
		})
	}
}

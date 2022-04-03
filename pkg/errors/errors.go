package errors

import (
	"errors"
	"fmt"
)

var (
	InvalidAddressError  = errors.New("invalid address access")
	InvalidMnemonicError = errors.New("invalid mnemonic")
	InvalidOperandError  = errors.New("invalid operand")
	NotImplementedError  = errors.New("not implemented")
)

func NewOperandSymbolError(got, want any) error {
	return fmt.Errorf("%w: invalid symbol: got: %T, want: %T", InvalidOperandError, got, want)
}

func NewReadError(addr uint16) error {
	return fmt.Errorf("%w: unable to write 0x%04x", InvalidAddressError, addr)
}

func NewWriteError(addr uint16) error {
	return fmt.Errorf("%w: unable to write 0x%04x", InvalidAddressError, addr)
}

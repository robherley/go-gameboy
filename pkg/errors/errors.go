package errors

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrorInvalidAddress     = errors.New("invalid address access")
	ErrorInvalidMnemonic    = errors.New("invalid mnemonic")
	ErrorInvalidOperand     = errors.New("invalid operand")
	ErrorInvalidSymbol      = errors.New("invalid symbol")
	ErrorInvalidInstruction = errors.New("invalid instruction")
	ErrorIllegalInstruction = errors.New("illegal instruction")
	ErrorNotImplemented     = errors.New("not implemented")
)

func NewInvalidOperandError(operand any) error {
	return fmt.Errorf("%w: %v (%T)", ErrorInvalidOperand, operand, operand)
}

func NewInvalidSymbolError(symbol any) error {
	return fmt.Errorf("%w: %v (%T)", ErrorInvalidSymbol, symbol, symbol)
}

func NewInvalidGetOperandError(operand any) error {
	return fmt.Errorf("%w: cannot get value of %v (%T)", ErrorInvalidOperand, operand, operand)
}

func NewInvalidSetOperandError(operand any) error {
	return fmt.Errorf("%w: cannot set value of %v (%T)", ErrorInvalidOperand, operand, operand)
}

func NewIllegalInstructionError(opcode byte) error {
	return fmt.Errorf("%w: 0x%X", ErrorIllegalInstruction, opcode)
}

func NewOperandSymbolError(got, want any) error {
	return fmt.Errorf("%w: invalid symbol: got: %T, want: %T", ErrorInvalidOperand, got, want)
}

func NewAccessError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to access %s at 0x%04x", ErrorInvalidAddress, resource, addr)
}

func NewReadError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to read %s at 0x%04x", ErrorInvalidAddress, resource, addr)
}

func NewWriteError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to write %s at 0x%04x", ErrorInvalidAddress, resource, addr)
}

func NewInvalidMnemonicError(mnemonic string) error {
	return fmt.Errorf("%w: %q", ErrorInvalidMnemonic, mnemonic)
}

func NewUnknownOPCodeError(opcode byte) error {
	return fmt.Errorf("%w: unknown opcode 0x%02x", ErrorInvalidInstruction, opcode)
}

func NewNotImplementedError() error {
	caller := "unknown"
	lineNo := 0
	if pc, _, line, ok := runtime.Caller(1); ok {
		if details := runtime.FuncForPC(pc); details != nil {
			caller = details.Name()
			lineNo = line
		}
	}
	return fmt.Errorf("%w: called from %s:%d", ErrorNotImplemented, caller, lineNo)
}

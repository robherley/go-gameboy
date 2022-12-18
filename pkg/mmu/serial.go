package mmu

import errs "github.com/robherley/go-gameboy/pkg/errors"

const (
	SB_SERIAL_TRANSFER = 0xFF01
	SC_SERIAL_CONTROL  = 0xFF02
)

type serial struct {
	transfer byte
	control  byte
}

func newSerial() *serial {
	return &serial{0x0, 0x0}
}

func (i *serial) Read(address uint16) byte {
	switch address {
	case SB_SERIAL_TRANSFER:
		return i.transfer
	case SC_SERIAL_CONTROL:
		return i.control
	default:
		panic(errs.NewReadError(address, "serial"))
	}
}

func (i *serial) Write(address uint16, data byte) {
	switch address {
	case SB_SERIAL_TRANSFER:
		i.transfer = data
	case SC_SERIAL_CONTROL:
		i.control = data
	default:
		panic(errs.NewWriteError(address, "serial"))
	}
}

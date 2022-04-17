package mmu

import (
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

const (
	SB_SERIAL_TRANSFER = 0xFF01
	SC_SERIAL_CONTROL  = 0xFF02
)

type io struct {
	serialData [2]byte
}

func newIO() *io {
	return &io{
		serialData: [2]byte{0x0, 0x0},
	}
}

func (i *io) Read(address uint16) byte {
	switch address {
	case SB_SERIAL_TRANSFER:
		return i.serialData[0]
	case SC_SERIAL_CONTROL:
		return i.serialData[1]
	default:
		panic(errs.NewReadError(address, "io"))
	}
}

func (i *io) Write(address uint16, data byte) {
	switch address {
	case SB_SERIAL_TRANSFER:
		i.serialData[0] = data
	case SC_SERIAL_CONTROL:
		i.serialData[1] = data
	default:
		panic(errs.NewWriteError(address, "io"))
	}
}

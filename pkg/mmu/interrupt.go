package mmu

import errs "github.com/robherley/go-gameboy/pkg/errors"

const (
	IF_INTERRUPT_FLAG   = 0xFF0F
	IE_INTERRUPT_ENABLE = 0xFFFF
)

type interrupt struct {
	// (IF - $FFOF) flag identifies if a specific interrupt bit becomes set
	flag byte
	// (IE - $FFFF) enable specifies if a specific interrupt bit is enabled
	enable byte
}

func newInterrupt() *interrupt {
	return &interrupt{0x0, 0x0}
}

func (i *interrupt) Read(address uint16) byte {
	switch address {
	case IF_INTERRUPT_FLAG:
		return i.flag
	case IE_INTERRUPT_ENABLE:
		return i.enable
	default:
		panic(errs.NewReadError(address, "interrupt"))
	}
}

func (i *interrupt) Write(address uint16, data byte) {
	switch address {
	case IF_INTERRUPT_FLAG:
		i.flag = data
	case IE_INTERRUPT_ENABLE:
		i.enable = data
	default:
		panic(errs.NewWriteError(address, "interrupt"))
	}
}

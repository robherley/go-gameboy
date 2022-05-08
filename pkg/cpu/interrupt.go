package cpu

import (
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/mmu"
)

/*
	Bit 0: VBlank   Interrupt Enable  (INT $40)
	Bit 1: LCD STAT Interrupt Enable  (INT $48)
	Bit 2: Timer    Interrupt Enable  (INT $50)
	Bit 3: Serial   Interrupt Enable  (INT $58)
	Bit 4: Joypad   Interrupt Enable  (INT $60)
*/

type InterruptType byte

const (
	VBLANK   InterruptType = 1
	LCD_STAT InterruptType = 2
	TIMER    InterruptType = 4
	SERIAL   InterruptType = 8
	JOYPAD   InterruptType = 16
)

type InterruptMasterChange byte

const (
	MASTER_SET_NONE InterruptMasterChange = 0
	MASTER_SET_NOW  InterruptMasterChange = 1
	MASTER_SET_NEXT InterruptMasterChange = 2
)

var (
	interrupts = [...]InterruptType{
		VBLANK, LCD_STAT, TIMER, SERIAL, JOYPAD,
	}

	// https://gbdev.io/pandocs/Interrupts.html#ff0f---if---interrupt-flag-rw
	interruptsToAddress = map[InterruptType]uint16{
		VBLANK:   0x40,
		LCD_STAT: 0x48,
		TIMER:    0x50,
		SERIAL:   0x58,
		JOYPAD:   0x60,
	}
)

// https://gbdev.io/pandocs/Interrupts.html
type Interrupt struct {
	// (IME) MasterEnabled is used to disabled all interrupts on the IE register
	MasterEnabled bool
	// EI: sets master to be enabled (delayed one instruction)
	EI InterruptMasterChange
	// DI: sets master to be disabled (delayed one instruction)
	DI InterruptMasterChange
	// (IE) enable specifies if a specific interrupt bit is enabled
	enable byte
	// (IF) flag identifies if a specific interrupt bit becomes set
	flag byte
}

// Write used by MMU to set IE flag
func (i *Interrupt) Write(address uint16, val byte) {
	if address != mmu.IN_ENABLE_REG {
		panic(errs.NewWriteError(address, "interrupt enable"))
	}

	i.enable = val
}

// Read used by MMU to read IE flag
func (i *Interrupt) Read(address uint16) byte {
	if address != mmu.IN_ENABLE_REG {
		panic(errs.NewWriteError(address, "interrupt enable"))
	}

	return i.enable
}

func (i *Interrupt) IsRequested() bool {
	return i.flag != 0
}

func (i *Interrupt) IsFlagged(it InterruptType) bool {
	return i.flag&byte(it) != 0
}

func (i *Interrupt) IsEnabled(it InterruptType) bool {
	return i.enable&byte(it) != 0
}

package cpu

import "github.com/robherley/go-gameboy/pkg/mmu"

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
}

func InterruptRequested(cpu *CPU) bool {
	return cpu.MMU.Read8(mmu.IF_INTERRUPT_FLAG) != 0
}

func InterruptTriggered(cpu *CPU, it InterruptType) bool {
	return InterruptEnabled(cpu, it) && InterruptFlagged(cpu, it)
}

func InterruptEnabled(cpu *CPU, it InterruptType) bool {
	enable := cpu.MMU.Read8(mmu.IE_INTERRUPT_ENABLE)
	return enable&byte(it) != 0
}

func InterruptFlagged(cpu *CPU, it InterruptType) bool {
	flag := cpu.MMU.Read8(mmu.IF_INTERRUPT_FLAG)
	return flag&byte(it) != 0
}

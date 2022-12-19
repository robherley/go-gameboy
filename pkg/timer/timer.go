package timer

import (
	"github.com/robherley/go-gameboy/internal/bits"
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

// https://gbdev.io/pandocs/Timer_and_Divider_Registers.html
type Timer struct {
	// FF04 - Divider register, interal system 16bit counter
	DIV uint16
	// FF05 - Timer counter
	TIMA byte
	// FF06 - Timer modulo
	TMA byte
	// FFO7 - Timer control
	TAC byte
	// Callback for interrupt
	OnInterrupt func()
}

func New(interruptFunc func()) *Timer {
	return &Timer{
		DIV:         0xAC00, // TODO: find out why this value
		TIMA:        0x00,
		TMA:         0x00,
		TAC:         0x00,
		OnInterrupt: interruptFunc,
	}
}

func (t *Timer) Tick() {
	prevDiv := t.DIV
	t.DIV++

	timerUpdate := false

	switch t.TAC & 0b11 {
	case 0b00:
		// 00: CPU Clock / 1024 (DMG, SGB2, CGB Single Speed Mode:   4096 Hz, SGB1:   ~4194 Hz, CGB Double Speed Mode:   8192 Hz)
		timerUpdate = checkBit(prevDiv, t.DIV, 9)
	case 0b01:
		// 01: CPU Clock / 16   (DMG, SGB2, CGB Single Speed Mode: 262144 Hz, SGB1: ~268400 Hz, CGB Double Speed Mode: 524288 Hz)
		timerUpdate = checkBit(prevDiv, t.DIV, 3)
	case 0b10:
		// 10: CPU Clock / 64   (DMG, SGB2, CGB Single Speed Mode:  65536 Hz, SGB1:  ~67110 Hz, CGB Double Speed Mode: 131072 Hz)
		timerUpdate = checkBit(prevDiv, t.DIV, 5)
	case 0b11:
		// 11: CPU Clock / 256  (DMG, SGB2, CGB Single Speed Mode:  16384 Hz, SGB1:  ~16780 Hz, CGB Double Speed Mode:  32768 Hz)
		timerUpdate = checkBit(prevDiv, t.DIV, 7)
	}

	// Bit  2   - Timer Enable
	if timerUpdate && (t.TAC&(1<<2) == 1) {
		// on overflow, set TIMA to TMA and request interrupt
		if t.TIMA == 0xFF {
			t.TIMA = t.TMA
			t.OnInterrupt()
		} else {
			t.TIMA++
		}
	}
}

// checkBit checks if the bit is set in the previous value and not set in the current value
func checkBit(prev, current uint16, bit byte) bool {
	return (prev&(1<<bit) == 1) && (current&(1<<bit) == 0)
}

func (t *Timer) Read(address uint16) byte {
	switch address {
	case DIV_ADDRESS:
		return bits.Hi(t.DIV)
	case TIMA_ADDRESS:
		return t.TIMA
	case TMA_ADDRESS:
		return t.TMA
	case TAC_ADDRESS:
		return t.TAC
	default:
		panic(errs.NewReadError(address, "timer"))
	}
}

func (t *Timer) Write(address uint16, data byte) {
	switch address {
	case DIV_ADDRESS:
		// https://gbdev.io/pandocs/Timer_and_Divider_Registers.html?search=#ff04--div-divider-register
		t.DIV = 0x0000
	case TIMA_ADDRESS:
		t.TIMA = data
	case TMA_ADDRESS:
		t.TMA = data
	case TAC_ADDRESS:
		t.TAC = data
	default:
		panic(errs.NewWriteError(address, "timer"))
	}
}

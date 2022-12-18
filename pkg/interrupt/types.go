package interrupt

/*
	Bit 0: VBlank   Interrupt Enable  (INT $40)
	Bit 1: LCD STAT Interrupt Enable  (INT $48)
	Bit 2: Timer    Interrupt Enable  (INT $50)
	Bit 3: Serial   Interrupt Enable  (INT $58)
	Bit 4: Joypad   Interrupt Enable  (INT $60)
*/

type Type byte

const (
	VBLANK   Type = 1
	LCD_STAT Type = 2
	TIMER    Type = 4
	SERIAL   Type = 8
	JOYPAD   Type = 16
)

var (
	Types = [...]Type{
		VBLANK, LCD_STAT, TIMER, SERIAL, JOYPAD,
	}

	// https://gbdev.io/pandocs/Interrupts.html#ff0f---if---interrupt-flag-rw
	TypeToAddress = map[Type]uint16{
		VBLANK:   0x40,
		LCD_STAT: 0x48,
		TIMER:    0x50,
		SERIAL:   0x58,
		JOYPAD:   0x60,
	}
)

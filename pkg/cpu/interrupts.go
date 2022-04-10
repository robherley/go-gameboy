package cpu

// https://gbdev.io/pandocs/Interrupts.html

type Interrupt byte

const (
	VBLANK   Interrupt = 1
	LCD_STAT Interrupt = 2
	TIMER    Interrupt = 4
	SERIAL   Interrupt = 8
	JOYPAD   Interrupt = 16
)

var (
	interrupts = [...]Interrupt{
		VBLANK, LCD_STAT, TIMER, SERIAL, JOYPAD,
	}

	// https://gbdev.io/pandocs/Interrupts.html#ff0f---if---interrupt-flag-rw
	interruptsToAddress = map[Interrupt]uint16{
		VBLANK:   0x40,
		LCD_STAT: 0x48,
		TIMER:    0x50,
		SERIAL:   0x58,
		JOYPAD:   0x60,
	}
)

func (c *CPU) HandleInterrupts() *Interrupt {
	for i := range interrupts {
		it := interrupts[i]
		addr := interruptsToAddress[it]
		itByte := byte(it)

		// only handle if *both* interrupt enable and interrupt flag are set
		if c.IF&itByte != 0 && c.IE&itByte != 0 {
			// 1. push program counter to stack
			c.StackPush16(c.Registers.PC)
			// 2. set program counter to mapped interrupt address
			c.Registers.PC = addr
			// 3. clear interrupt flag
			c.IF &= ^itByte
			// 4. unhalt cpu
			c.IsHalted = false
			// 5. disable all interrupts
			c.IME = false

			// only handle one interrupt at a time, priority is based on bit order
			return &it
		}
	}

	return nil
}

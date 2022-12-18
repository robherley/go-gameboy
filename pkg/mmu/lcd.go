package mmu

type lcd struct{}

func newLCD() *lcd {
	return &lcd{}
}

func (l *lcd) Read(address uint16) byte {
	// TODO(robherley): remove, debug for gb doctor
	if address == 0xFF44 {
		return 0x90
	}

	// TODO(robherley): implement
	return 0x0
}

func (l *lcd) Write(address uint16, data byte) {
	// TODO(robherley): implement
}

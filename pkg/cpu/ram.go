package cpu

import "fmt"

const (
	WRAM_SIZE = 0x2000
	HRAM_SIZE = 0x80
)

type RAM struct {
	wram [WRAM_SIZE]byte
	hram [HRAM_SIZE]byte
}

func (r *RAM) translateAddress(address uint16) uint16 {
	// subtracting where it starts in mmap context (see mmu.go)
	internalAddress := address - 0xC000
	if internalAddress >= WRAM_SIZE {
		panic(fmt.Errorf("invalid wram access at %04x", address))
	}

	return internalAddress
}

func (r *RAM) Read8(address uint16) byte {
	return r.wram[r.translateAddress(address)]
}

func (r *RAM) Write8(address uint16, value byte) {
	r.wram[r.translateAddress(address)] = value
}

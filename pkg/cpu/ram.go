package cpu

import (
	errs "github.com/robherley/go-dmg/pkg/errors"
)

const (
	WRAM_SIZE = 0x2000
	HRAM_SIZE = 0x80

	// offsets of where the memory region starts (see mmu.go)
	WRAM_OFFSET = 0xC000
	HRAM_OFFSET = 0xFF80
)

type RAM struct {
	wram [WRAM_SIZE]byte
	hram [HRAM_SIZE]byte
}

func (r *RAM) translateAddress(address, offset uint16) uint16 {
	internalAddress := address - offset
	if internalAddress >= WRAM_SIZE {
		panic(errs.NewAccessError(address, "wram"))
	}

	return internalAddress
}

func (r *RAM) WRAMRead(address uint16) byte {
	return r.wram[r.translateAddress(address, WRAM_OFFSET)]
}

func (r *RAM) WRAMWrite(address uint16, value byte) {
	r.wram[r.translateAddress(address, WRAM_OFFSET)] = value
}

func (r *RAM) HRAMRead(address uint16) byte {
	return r.hram[r.translateAddress(address, HRAM_OFFSET)]
}

func (r *RAM) HRAMWrite(address uint16, value byte) {
	r.hram[r.translateAddress(address, HRAM_OFFSET)] = value
}

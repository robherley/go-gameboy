package mmu

import (
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

const (
	// fixed size of each memory region
	WRAM_SIZE = 0x2000
	HRAM_SIZE = 0x80

	// offsets of where the memory region starts
	WRAM_OFFSET = 0xC000
	HRAM_OFFSET = 0xFF80
)

type ram struct {
	memory []byte
	size   uint16
	offset uint16
}

func newWRAM() *ram {
	return &ram{
		memory: make([]byte, WRAM_SIZE),
		size:   WRAM_SIZE,
		offset: WRAM_OFFSET,
	}
}

func newHRAM() *ram {
	return &ram{
		memory: make([]byte, HRAM_SIZE),
		size:   HRAM_SIZE,
		offset: HRAM_OFFSET,
	}
}

func (r *ram) translateAddress(address uint16) uint16 {
	internalAddress := address - r.offset
	if internalAddress >= r.size {
		panic(errs.NewAccessError(address, "ram"))
	}

	return internalAddress
}

func (r *ram) Read(address uint16) byte {
	return r.memory[r.translateAddress(address)]
}

func (r *ram) Write(address uint16, data byte) {
	r.memory[r.translateAddress(address)] = data
}

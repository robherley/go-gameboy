package mmu

import (
	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"

	errs "github.com/robherley/go-gameboy/pkg/errors"
)

type MMU struct {
	cartridge   *cartridge.Cartridge
	interruptRW readWriter
	hram        *ram
	wram        *ram
	io          *io
}

func New(cart *cartridge.Cartridge, interruptRW readWriter) *MMU {
	return &MMU{
		cartridge:   cart,
		interruptRW: interruptRW,
		hram:        newHRAM(),
		wram:        newWRAM(),
		io:          newIO(),
	}
}

func (mmu *MMU) Read8(address uint16) byte {
	rw := mmu.readWriterFor(address)
	if rw == nil {
		panic(errs.NewReadError(address, "mmu"))
	}

	return rw.Read(address)
}

func (mmu *MMU) Read16(address uint16) uint16 {
	lo := mmu.Read8(address)
	hi := mmu.Read8(address + 1)

	return bits.To16(hi, lo)
}

func (mmu *MMU) Write8(address uint16, data byte) {
	rw := mmu.readWriterFor(address)
	if rw == nil {
		panic(errs.NewReadError(address, "mmu"))
	}

	rw.Write(address, data)
}

func (mmu *MMU) Write16(address uint16, value uint16) {
	mmu.Write8(address, bits.Lo(value))
	mmu.Write8(address+1, bits.Hi(value))
}

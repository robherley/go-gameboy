package mmu

import (
	"fmt"

	"github.com/robherley/go-gameboy/internal/bits"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/interrupt"
)

type MMU struct {
	cartridge *cartridge.Cartridge
	hram      *ram
	wram      *ram
	serial    *serial
	interrupt *interrupt.Interrupt
	lcd       *lcd
}

func New(cart *cartridge.Cartridge, it *interrupt.Interrupt) *MMU {
	return &MMU{
		cartridge: cart,
		interrupt: it,
		hram:      newHRAM(),
		wram:      newWRAM(),
		serial:    newSerial(),
		lcd:       newLCD(),
	}
}

func (mmu *MMU) Read8(address uint16) byte {
	rw := mmu.readerWriterFor(address)
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
	rw := mmu.readerWriterFor(address)
	if rw == nil {
		panic(errs.NewWriteError(address, "mmu"))
	}

	rw.Write(address, data)
}

func (mmu *MMU) Write16(address uint16, value uint16) {
	mmu.Write8(address, bits.Lo(value))
	mmu.Write8(address+1, bits.Hi(value))
}

func (mmu *MMU) DebugMem() {
	// fmt.Printf("%X\n", mmu.Read8(0xd800))
}

func (mmu *MMU) DebugSerial() {
	// if first and last bits are set, read in debug data
	if mmu.serial.Read(SC_SERIAL_CONTROL) == 0x81 {
		ch := mmu.serial.Read(SB_SERIAL_TRANSFER)
		fmt.Printf("%c", ch)
		mmu.serial.Write(SC_SERIAL_CONTROL, 0)
	}
}

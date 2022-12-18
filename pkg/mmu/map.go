package mmu

import (
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

// Range is a inclusive start/end range defined in memory
type Range struct {
	Start, End uint16
}

// Contains checks if the given address is included in the memory range
func (r *Range) Contains(addr uint16) bool {
	return r.Start <= addr && addr <= r.End
}

// https://gbdev.io/pandocs/Memory_Map.html
var (
	// 0x0000 - 0x3FFF : ROM Bank 0
	// 0x4000 - 0x7FFF : ROM Bank 1 - Switchable
	ROMRange = Range{0x0000, 0x7FFF}
	// 0x8000 - 0x97FF : CHR RAM
	// 0x9800 - 0x9BFF : BG Map 1
	// 0x9C00 - 0x9FFF : BG Map 2
	CharMapRange = Range{0x8000, 0x9FFF}
	// 0xA000 - 0xBFFF : Cartridge RAM
	CartRAMRange = Range{0xA000, 0xBFFF}
	// 0xC000 - 0xCFFF : RAM Bank 0
	// 0xD000 - 0xDFFF : RAM Bank 1-7 - switchable - Color only
	WRAMRange = Range{0xC000, 0xDFFF}
	// 0xE000 - 0xFDFF : Reserved - Echo RAM
	RESERVED_EchoRamRange = Range{0xE000, 0xFDFF}
	// 0xFE00 - 0xFE9F : Object Attribute Memory
	OAMRange = Range{0xFE00, 0xFE9F}
	// 0xFEA0 - 0xFEFF : Reserved - Unusable
	RESERVED_UnusableRange = Range{0xFEA0, 0xFEFF}
	// 0xFF00 : Joypad input
	JobpadInputRange = Range{0xFF00, 0xFF00}
	// 0xFF01 - 0xFF02 : Serial transfer
	SerialTransferRange = Range{0xFF01, 0xFF02}
	// 0xFF04 - 0xFF07 : Timer and divider
	TimerDividerRange = Range{0xFF04, 0xFF07}
	// 0xFF0F : Interrupt flag register
	InterruptFlagRange = Range{0xFF0F, 0xFF0F}
	// 0xFF10 - 0xFF26 : Audio
	// 0xFF30 - 0xFF3F : Wave pattern
	AudioRange = Range{0xFF10, 0xFF3F}
	// 0xFF40 - 0xFF4B : LCD
	LCDRange = Range{0xFF40, 0xFF4B}
	// 0xFF4D : CGB Speed Switch
	ColorSpeedSwitchRange = Range{0xFF4D, 0xFF4D}
	// 0xFF4F : VRAM Bank Select
	VRAMBankSelectRange = Range{0xFF4F, 0xFF4F}
	// 0xFF50 : Disable boot rom
	DisableBootRomRange = Range{0xFF50, 0xFF50}
	// 0xFF51 - 0xFF55 : VRAM DMA
	VRAMDMARange = Range{0xFF51, 0xFF55}
	// 0xFF68 - 0xFF69 : BG / OBJ Palettes
	BgObjPaletteRange = Range{0xFF51, 0xFF55}
	// 0xFF70 : WRAM Bank Select
	WRAMBankSelectRange = Range{0xFF50, 0xFF50}
	// 0xFF80 - 0xFFFE : Zero Page (HRAM)
	HRAMRange = Range{0xFF80, 0xFFFE}
	// 0xFFFF : Interrupt enable register
	InterruptEnableRange = Range{0xFFFF, 0xFFFF}
)

type readerWriter interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}

func (mmu *MMU) readerWriterFor(addr uint16) readerWriter {
	strict := false

	if ROMRange.Contains(addr) {
		return mmu.cartridge
	} else if CharMapRange.Contains(addr) {
		return newNoop(strict)
	} else if CartRAMRange.Contains(addr) {
		return mmu.cartridge
	} else if WRAMRange.Contains(addr) {
		return mmu.wram
	} else if RESERVED_EchoRamRange.Contains(addr) {
		panic(errs.NewAccessError(addr, "reserved echo memory"))
	} else if OAMRange.Contains(addr) {
		return newNoop(strict)
	} else if RESERVED_UnusableRange.Contains(addr) {
		panic(errs.NewAccessError(addr, "reserved unused memory"))
	} else if JobpadInputRange.Contains(addr) {
		return newNoop(strict)
	} else if SerialTransferRange.Contains(addr) {
		return mmu.serial
	} else if TimerDividerRange.Contains(addr) {
		return newNoop(strict)
	} else if InterruptFlagRange.Contains(addr) {
		return mmu.interrupt
	} else if AudioRange.Contains(addr) {
		return newNoop(strict)
	} else if LCDRange.Contains(addr) {
		return newNoop(strict)
	} else if ColorSpeedSwitchRange.Contains(addr) {
		return newNoop(strict)
	} else if VRAMBankSelectRange.Contains(addr) {
		return newNoop(strict)
	} else if DisableBootRomRange.Contains(addr) {
		return newNoop(strict)
	} else if VRAMDMARange.Contains(addr) {
		return newNoop(strict)
	} else if BgObjPaletteRange.Contains(addr) {
		return newNoop(strict)
	} else if WRAMBankSelectRange.Contains(addr) {
		return newNoop(strict)
	} else if HRAMRange.Contains(addr) {
		return mmu.hram
	} else if InterruptEnableRange.Contains(addr) {
		return mmu.interrupt
	} else {
		panic(errs.NewAccessError(addr, "mmu"))
	}
}

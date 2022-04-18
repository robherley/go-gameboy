package mmu

import (
	errs "github.com/robherley/go-gameboy/pkg/errors"
)

// https://gbdev.io/pandocs/Memory_Map.html
const (
	// 0x0000 - 0x3FFF : ROM Bank 0
	// 0x4000 - 0x7FFF : ROM Bank 1 - Switchable
	ROM_END = 0x7FFF
	// 0x8000 - 0x97FF : CHR RAM
	// 0x9800 - 0x9BFF : BG Map 1
	// 0x9C00 - 0x9FFF : BG Map 2
	CHAR_MAP_END = 0x9FFF
	// 0xA000 - 0xBFFF : Cartridge RAM
	CART_RAM_END = 0xBFFF
	// 0xC000 - 0xCFFF : RAM Bank 0
	// 0xD000 - 0xDFFF : RAM Bank 1-7 - switchable - Color only
	WRAM_END = 0xDFFF
	// 0xE000 - 0xFDFF : Reserved - Echo RAM
	RES_ECHO_END = 0xFDFF
	// 0xFE00 - 0xFE9F : Object Attribute Memory
	OAM_END = 0xFE9F
	// 0xFEA0 - 0xFEFF : Reserved - Unusable
	RES_UNUSE_END = 0xFE9F
	// 0xFF00 - 0xFF7F : I/O Registers
	IO_REG_END = 0xFF7F
	// 0xFF80 - 0xFFFE : Zero Page (HRAM)
	HRAM_END = 0xFFFE
	// 0xFFFE - 0xFFFF : Interrupt enable register
	IN_ENABLE_REG = 0xFFFF
)

type readWriter interface {
	Read(address uint16) byte
	Write(address uint16, data byte)
}

func (mmu *MMU) readWriterFor(address uint16) readWriter {
	if address <= ROM_END {
		return mmu.cartridge
	} else if address <= CHAR_MAP_END {
		return nil
	} else if address <= CART_RAM_END {
		return mmu.cartridge
	} else if address <= WRAM_END {
		return mmu.wram
	} else if address <= RES_ECHO_END {
		panic(errs.NewAccessError(address, "reserved echo memory"))
	} else if address <= OAM_END {
		return nil
	} else if address <= RES_UNUSE_END {
		panic(errs.NewAccessError(address, "reserved unused memory"))
	} else if address <= IO_REG_END {
		return mmu.io
	} else if address <= HRAM_END {
		return mmu.hram
	} else if address == IN_ENABLE_REG {
		return mmu.interruptRW
	}

	panic(errs.NewReadError(address, "mmu"))
}

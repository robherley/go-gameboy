package cpu

import (
	"fmt"

	"github.com/robherley/go-gameboy/internal/bits"
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
	// 0xFF80 - 0xFFFF : Interrupt enable register
	IN_ENABLE_REG = 0xFFFF
)

func (c *CPU) Read8(address uint16) byte {
	if address <= ROM_END {
		return c.Cartridge.Read(address)
	} else if address <= CHAR_MAP_END {
		panic(errs.NewReadError(address, "TODO read char map"))
	} else if address <= CART_RAM_END {
		return c.Cartridge.Read(address)
	} else if address <= WRAM_END {
		return c.RAM.WRAMRead(address)
	} else if address <= RES_ECHO_END {
		panic(errs.NewReadError(address, "reserved echo memory"))
	} else if address <= OAM_END {
		panic(errs.NewReadError(address, "TODO OAM"))
	} else if address <= RES_UNUSE_END {
		panic(errs.NewReadError(address, "reserved unused memory"))
	} else if address <= IO_REG_END {
		panic(errs.NewReadError(address, "TODO io read reg"))
	} else if address <= HRAM_END {
		return c.RAM.HRAMRead(address)
	} else if address == IN_ENABLE_REG {
		return c.IE
	}

	panic(errs.NewReadError(address, "mmu"))
}

func (c *CPU) Read16(address uint16) uint16 {
	lo := c.Read8(address)
	hi := c.Read8(address + 1)

	return bits.To16(hi, lo)
}

func (c *CPU) Write8(address uint16, value byte) {
	if address <= ROM_END {
		c.Cartridge.Write(address, value)
		return
	} else if address <= CHAR_MAP_END {
		panic(errs.NewWriteError(address, "TODO char map"))
	} else if address <= CART_RAM_END {
		c.Cartridge.Write(address, value)
		return
	} else if address <= WRAM_END {
		c.RAM.WRAMWrite(address, value)
		return
	} else if address <= RES_ECHO_END {
		panic(errs.NewWriteError(address, "reserved echo memory"))
	} else if address <= OAM_END {
		panic(errs.NewWriteError(address, "TODO OAM"))
	} else if address <= RES_UNUSE_END {
		panic(errs.NewWriteError(address, "reserved unused memory"))
	} else if address <= IO_REG_END {
		fmt.Println(fmt.Errorf("TODO write io reg: 0x%04x", address))
		return
	} else if address <= HRAM_END {
		c.RAM.HRAMWrite(address, value)
		return
	} else if address == IN_ENABLE_REG {
		c.IE = value
		return
	}

	panic(errs.NewWriteError(address, "mmu"))
}

func (c *CPU) Write16(address uint16, value uint16) {
	c.Write8(address, bits.Lo(value))
	c.Write8(address+1, bits.Hi(value))
}

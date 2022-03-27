package cpu

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/cartridge"
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

type MMU struct {
	Cartridge *cartridge.Cartridge
}

func (m *MMU) Read8(address uint16) byte {
	if address <= ROM_END {
		return m.Cartridge.Read(address)
	} else if address <= CHAR_MAP_END {
		panic(fmt.Errorf("TODO read char map: 0x%04x", address))
	} else if address <= CART_RAM_END {
		return m.Cartridge.Read(address)
	} else if address <= WRAM_END {
		panic(fmt.Errorf("TODO read wram: 0x%04x", address))
	} else if address <= RES_ECHO_END {
		panic(fmt.Errorf("invalid read of reserved echo memory: 0x%04x", address))
	} else if address <= OAM_END {
		panic(fmt.Errorf("TODO read OAM: 0x%04x", address))
	} else if address <= RES_UNUSE_END {
		panic(fmt.Errorf("invalid read of reserved unused memory: 0x%04x", address))
	} else if address <= IO_REG_END {
		panic(fmt.Errorf("TODO read io reg: 0x%04x", address))
	} else if address <= HRAM_END {
		panic(fmt.Errorf("TODO read hram: 0x%04x", address))
	} else if address == IN_ENABLE_REG {
		panic(fmt.Errorf("TODO read interrupt reg: 0x%04x", address))
	}

	panic(fmt.Errorf("invalid out of bound memory read: 0x%04x", address))
}

func (m *MMU) Read16(address uint16) uint16 {
	lo := m.Read8(address)
	hi := m.Read8(address + 1)

	return bits.To16(hi, lo)
}

func (m *MMU) Write8(address uint16, value byte) {
	if address <= ROM_END {
		m.Cartridge.Write(address, value)
		return
	} else if address <= CHAR_MAP_END {
		panic(fmt.Errorf("TODO write char map: 0x%04x", address))
	} else if address <= CART_RAM_END {
		m.Cartridge.Write(address, value)
		return
	} else if address <= WRAM_END {
		panic(fmt.Errorf("TODO write wram: 0x%04x", address))
	} else if address <= RES_ECHO_END {
		panic(fmt.Errorf("invalid write of reserved echo memory: 0x%04x", address))
	} else if address <= OAM_END {
		panic(fmt.Errorf("TODO write OAM: 0x%04x", address))
	} else if address <= RES_UNUSE_END {
		panic(fmt.Errorf("invalid write of reserved unused memory: 0x%04x", address))
	} else if address <= IO_REG_END {
		panic(fmt.Errorf("TODO write io reg: 0x%04x", address))
	} else if address <= HRAM_END {
		panic(fmt.Errorf("TODO write hram: 0x%04x", address))
	} else if address == IN_ENABLE_REG {
		panic(fmt.Errorf("TODO write interrupt reg: 0x%04x", address))
	}

	panic(fmt.Errorf("invalid out of bound memory write: 0x%04x", address))
}

func (m *MMU) Write16(address uint16, value uint16) {
	m.Write8(address, bits.Lo(value))
	m.Write8(address+1, bits.Hi(value))
}

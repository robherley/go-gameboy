package cpu

import (
	"github.com/robherley/go-dmg/internal/bits"
	"github.com/robherley/go-dmg/pkg/cartridge"
)

// https://gbdev.io/pandocs/Memory_Map.html

const (
	ROM_END = 0x8000
)

type MMU struct {
	Cartridge *cartridge.Cartridge
}

func (m *MMU) read8(address uint16) byte {
	if address < ROM_END {
		return m.Cartridge.Read(address)
	}

	panic("not implemented")
}

func (m *MMU) read16(address uint16) uint16 {
	hi := m.read8(address)
	lo := m.read8(address + 1)

	return bits.To16(hi, lo)
}

func (m *MMU) write8(address uint16, value byte) {
	if address < ROM_END {
		// some cartridges allow special ops for ROM write
		m.Cartridge.Write(address, value)
	}

	panic("not implemented")
}

func (m *MMU) write16(address uint16, value uint16) {
	m.write8(address, bits.Hi(value))
	m.write8(address+1, bits.Lo(value))
}

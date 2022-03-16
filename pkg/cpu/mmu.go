package cpu

import "github.com/robherley/go-dmg/pkg/cartridge"

// https://gbdev.io/pandocs/Memory_Map.html

const (
	ROM_END = 0x8000
)

type MMU struct {
	Cartridge *cartridge.Cartridge
}

func (m *MMU) read(address uint16) byte {
	if address < ROM_END {
		return m.Cartridge.Read(address)
	}

	panic("not implemented")
}

func (m *MMU) write(address uint16, value byte) {
	if address < ROM_END {
		// some cartridges allow special ops for ROM write
		m.Cartridge.Write(address, value)
	}

	panic("not implemented")
}

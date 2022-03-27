package cpu

import (
	"fmt"

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

func (m *MMU) Read8(address uint16) byte {
	if address < ROM_END {
		return m.Cartridge.Read(address)
	}

	fmt.Printf("UNSUPPORTED read of 0x%X\n", address)
	// panic("not implemented")

	return 0
}

func (m *MMU) Read16(address uint16) uint16 {
	lo := m.Read8(address)
	hi := m.Read8(address + 1)

	return bits.To16(hi, lo)
}

func (m *MMU) Write8(address uint16, value byte) {
	// if address < ROM_END {
	// 	// some cartridges allow special ops for ROM write
	// 	m.Cartridge.Write(address, value)
	// }

	fmt.Printf("UNSUPPORTED write of 0x%X\n", address)
	// panic("not implemented")
}

func (m *MMU) Write16(address uint16, value uint16) {
	m.Write8(address, bits.Lo(value))
	m.Write8(address+1, bits.Hi(value))
}

package dmg

// https://gbdev.io/pandocs/Memory_Map.html

const (
	ROM_END = 0x8000
)

func (e *Emulator) BusRead(address uint16) byte {
	if address < ROM_END {
		return e.Cartridge.Read(address)
	}

	panic("not implemented")
}

func (e *Emulator) BusWrite(address uint16, value byte) {
	if address < ROM_END {
		// some cartridges allow special ops for ROM write
		e.Cartridge.Write(address, value)
	}

	panic("not implemented")
}

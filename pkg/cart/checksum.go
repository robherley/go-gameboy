package cart

// https://gbdev.io/pandocs/The_Cartridge_Header.html#014d---header-checksum
func (c *Cartridge) CalculateHeaderCheckSum() byte {
	sum := byte(0)
	addr := 0x134
	for addr <= 0x14C {
		sum = sum - c.Data[addr] - 1
		addr++
	}

	return sum
}

func (c *Cartridge) IsValidHeaderCheckSum() bool {
	return c.CalculateHeaderCheckSum() == c.HeaderChecksum()
}

// https://gbdev.io/pandocs/The_Cartridge_Header.html#014e-014f---global-checksum
// The GameBoy actually doesn't verify this 🤔
func (c *Cartridge) CalculateGlobalCheckSum() uint16 {
	sum := uint16(0)

	for addr := range c.Data {
		// skip global checksum bytes
		if addr == 0x14E || addr == 0x14F {
			continue
		}
		sum += uint16(c.Data[addr])
	}

	return sum
}

func (c *Cartridge) IsValidGlobalCheckSum() bool {
	return c.CalculateGlobalCheckSum() == c.GlobalChecksum()
}

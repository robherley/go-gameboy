package cartridge

import (
	"fmt"

	"github.com/robherley/go-gameboy/internal/bits"
)

// https://gbdev.io/pandocs/The_Cartridge_Header.html

func (c *Cartridge) EntryPoint() [4]byte {
	return *(*[4]byte)(c.Data[0x100 : 0x103+1])
}

func (c *Cartridge) NintendoLogo() [48]byte {
	return *(*[48]byte)(c.Data[0x104 : 0x133+1])
}

func (c *Cartridge) Title() [10]byte {
	return *(*[10]byte)(c.Data[0x134 : 0x143+1])
}

func (c *Cartridge) TitleString() string {
	return fmt.Sprintf("%s", c.Title())
}

func (c *Cartridge) ManufacturerCode() [4]byte {
	return *(*[4]byte)(c.Data[0x13F : 0x142+1])
}

func (c *Cartridge) CGBFlag() byte {
	return c.Data[0x143]
}

func (c *Cartridge) SupportsColor() bool {
	// https://gbdev.io/pandocs/The_Cartridge_Header.html#0143---cgb-flag
	return c.CGBFlag() == 0x80
}

func (c *Cartridge) ColorOnly() bool {
	// https://gbdev.io/pandocs/The_Cartridge_Header.html#0143---cgb-flag
	return c.CGBFlag() == 0xC0
}

func (c *Cartridge) NewLicenseeCode() [2]byte {
	return *(*[2]byte)(c.Data[0x144 : 0x145+1])
}

func (c *Cartridge) NewLicenseeString() string {
	code := fmt.Sprint(c.NewLicenseeCode())
	return NewLicenseeToPublisher[code]
}

func (c *Cartridge) SGBFlag() byte {
	return c.Data[0x146]
}

func (c *Cartridge) CartridgeType() CartridgeType {
	return CartridgeType(c.Data[0x147])
}

func (c *Cartridge) ROMSize() byte {
	return c.Data[0x148]
}

func (c *Cartridge) RAMSize() byte {
	return c.Data[0x149]
}

func (c *Cartridge) DestinationCode() byte {
	return c.Data[0x14A]
}

func (c *Cartridge) OldLicenseeCode() byte {
	return c.Data[0x14B]
}

func (c *Cartridge) OldLicenseeString() string {
	return OldLicenseeToPublisher[c.Data[0x14B]]
}

func (c *Cartridge) IsNewLicensee() bool {
	// https://gbdev.io/pandocs/The_Cartridge_Header.html#014b---old-licensee-code
	return c.OldLicenseeCode() == 0x33
}

func (c *Cartridge) LicenseeString() string {
	if c.IsNewLicensee() {
		return c.NewLicenseeString()
	}

	return c.OldLicenseeString()
}

func (c *Cartridge) MaskRomVersion() byte {
	return c.Data[0x14C]
}

func (c *Cartridge) HeaderChecksum() byte {
	return c.Data[0x14D]
}

func (c *Cartridge) GlobalChecksum() uint16 {
	return bits.To16(c.Data[0x14E], c.Data[0x14F])
}

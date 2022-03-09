package cart

import (
	"fmt"
	"os"
)

type Cartridge struct {
	Data []byte
	Size int
}

func FromFile(filepath string) (*Cartridge, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read cartridge file: %w", err)
	}

	return FromBytes(data)
}

func FromBytes(data []byte) (*Cartridge, error) {
	return &Cartridge{
		Data: data,
		Size: len(data),
	}, nil
}

func (c *Cartridge) PrettyPrint() {
	fmt.Println("Title:", c.TitleString())

	licensee := c.OldLicenseeCodeString()
	if c.ShouldUseNewLicenseeCode() {
		licensee = c.NewLicenseeCodeString()
	}

	fmt.Println("Licensee:", licensee)
	fmt.Printf("Size: %dK\n", c.Size/1024)
	fmt.Println("Header Checksum Match:", c.IsValidHeaderCheckSum())
	fmt.Println("Global Checksum Match:", c.IsValidGlobalCheckSum())
}

func (c *Cartridge) Read(address uint16) byte {
	return c.Data[address]
}

func (e *Cartridge) Write(address uint16, value byte) {
	panic("not implemented")
}

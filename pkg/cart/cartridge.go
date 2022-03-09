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
}

package cartridge

import (
	"fmt"
	"os"

	errs "github.com/robherley/go-gameboy/pkg/errors"
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

func (c *Cartridge) Read(address uint16) byte {
	return c.Data[address]
}

func (c *Cartridge) Write(address uint16, value byte) {
	panic(errs.NewWriteError(address, "cartridge"))
}

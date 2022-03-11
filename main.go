package main

import (
	"github.com/robherley/go-dmg/pkg/cartridge"
)

func main() {
	cart, err := cartridge.FromFile("roms/Tetris.gb")
	if err != nil {
		panic(err)
	}

	cart.PrettyPrint()
}

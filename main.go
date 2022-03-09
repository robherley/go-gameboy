package main

import (
	"github.com/robherley/go-dmg/pkg/cart"
)

func main() {
	cartridge, err := cart.FromFile("roms/Tetris.gb")
	if err != nil {
		panic(err)
	}

	cartridge.PrettyPrint()
}

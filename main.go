package main

import (
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/emulator"
)

func main() {
	cart, err := cartridge.FromFile("roms/Tetris.gb")
	if err != nil {
		panic(err)
	}

	cart.PrettyPrint()

	emu := emulator.New(cart)
	emu.Boot()
}

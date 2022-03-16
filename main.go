package main

import (
	"github.com/robherley/go-dmg/internal/pretty"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/emulator"
)

func main() {
	pretty.Title()

	cart, err := cartridge.FromFile("roms/cpu_instrs.gb")
	if err != nil {
		panic(err)
	}

	pretty.Cart(cart)

	emu := emulator.New(cart)
	emu.Boot()
}

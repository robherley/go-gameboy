package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/emulator"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <path-to-rom>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	cart, err := cartridge.FromFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// debug.Cart(cart)

	emu := emulator.New(cart)
	emu.Boot()
}

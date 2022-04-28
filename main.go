package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robherley/go-gameboy/internal/pretty"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/emulator"
	"github.com/robherley/go-gameboy/pkg/ui"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <path-to-rom>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	fmt.Println("hello?")

	pretty.Title()

	cart, err := cartridge.FromFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	pretty.Cart(cart)

	go func() {
		emu := emulator.New(cart)
		emu.Boot()
	}()

	// TODO: maybe wait for sdl window to init first? seems slow

	err = ui.NewSDLWindow()
	if err != nil {
		panic(err)
	}
}

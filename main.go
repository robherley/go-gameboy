package main

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/emulator"
)

// TODO: move these pretty* funcs into their own internal pkg

func prettyTitle() {
	fmt.Println()
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("go-dmg", pterm.NewStyle(pterm.FgLightMagenta)),
	).Render()
}

func prettyCart(c *cartridge.Cartridge) {
	headerCheck := pterm.FgGreen.Sprintf("Match (0x%02x) ✅", c.HeaderChecksum())
	if !c.IsValidHeaderCheckSum() {
		headerCheck = pterm.FgRed.Sprintf("Mismatch (0x%02x) ❌", c.HeaderChecksum())
	}

	globalCheck := pterm.FgGreen.Sprintf("Match (0x%04x) ✅", c.GlobalChecksum())
	if !c.IsValidGlobalCheckSum() {
		globalCheck = pterm.FgRed.Sprintf("Mismatch (0x%04x) ❌", c.GlobalChecksum())
	}

	pterm.DefaultTable.WithData(pterm.TableData{
		{"Title", pterm.FgCyan.Sprint(c.TitleString())},
		{"Type", fmt.Sprintf("%s (0x%02x)", c.CartridgeTypeString(), c.CartridgeType())},
		{"Licensee", c.LicenseeString()},
		{"Size", fmt.Sprintf("%dK", c.Size/1024)},
		{"Header Checksum", headerCheck},
		{"Global Checksum", globalCheck},
	}).Render()
	fmt.Println()
}

func main() {
	prettyTitle()

	cart, err := cartridge.FromFile("roms/cpu_instrs.gb")
	if err != nil {
		panic(err)
	}

	prettyCart(cart)

	emu := emulator.New(cart)
	emu.Boot()
}

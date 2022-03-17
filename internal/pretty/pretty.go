package pretty

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/cpu"
)

var Hide = false

func Title() {
	if Hide {
		return
	}

	fmt.Println()
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("go-dmg", pterm.NewStyle(pterm.FgLightMagenta)),
	).Render()
}

func Cart(c *cartridge.Cartridge) {
	if Hide {
		return
	}

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

func Instruction(instr *cpu.Instruction, opcode byte, pc uint16) {
	if Hide {
		return
	}

	opsString := ""
	if instr.Operands != nil {
		opsString = pterm.Cyan(strings.Join(instr.Operands, " "))
	}

	instructionString := pterm.NewStyle(pterm.BgCyan, pterm.FgBlack).Sprintf(" %s (%02x) ", instr.Mnemonic, opcode)

	fmt.Printf("%s %s PC: 0x%x\n", instructionString, opsString, pc)
}

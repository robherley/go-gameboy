package pretty

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/instructions"
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

func Instruction(pc uint16, opcode byte, in *instructions.Instruction) {
	if Hide {
		return
	}

	pterm.NewStyle(pterm.FgBlack, pterm.BgWhite, pterm.Bold).Printf(" %04x ", pc)
	pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold).Printf(" % -3s (%02x) ", in.Mnemonic, opcode)
	pterm.Print(" ")
	if in.Operands == nil {
		pterm.FgMagenta.Print("<nil>")
	} else {
		for i, op := range in.Operands {
			symbol := fmt.Sprintf("%v", op.Symbol)
			if bite, ok := op.Symbol.(byte); ok {
				symbol = fmt.Sprintf("0x%02x", bite)
			}

			if op.Inc {
				symbol += "+"
			}

			if op.Dec {
				symbol += "-"
			}

			if op.Deref {
				symbol = fmt.Sprintf("(%s)", symbol)
			}

			pterm.FgMagenta.Print(symbol)

			if i != len(in.Operands)-1 {
				pterm.FgMagenta.Print(", ")
			}
		}
	}
	pterm.FgGray.Printf(" (len=%d)\n", len(in.Operands))
}

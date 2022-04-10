package pretty

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/cpu"
	"github.com/robherley/go-gameboy/pkg/instructions"
)

var Hide = false

func Title() {
	if Hide {
		return
	}

	fmt.Println()
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("go-gameboy", pterm.NewStyle(pterm.FgLightMagenta)),
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
		{"Type", fmt.Sprintf("%s (0x%02x)", c.CartridgeType().String(), c.CartridgeType())},
		{"Licensee", c.LicenseeString()},
		{"Size", fmt.Sprintf("%dK", c.Size/1024)},
		{"Header Checksum", headerCheck},
		{"Global Checksum", globalCheck},
	}).Render()
	fmt.Println()
}

func Instruction(pc uint16, opcode byte, in *instructions.Instruction, isCBPrefixed bool) {
	if Hide {
		return
	}

	pterm.NewStyle(pterm.FgBlack, pterm.BgWhite, pterm.Bold).Printf(" %04x ", pc)
	pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold).Printf(" % -3s (%02x) ", in.Mnemonic, opcode)
	pterm.Print(" ")
	if in.Operands == nil {
		pterm.FgMagenta.Print("<nil>")
	} else {
		if isCBPrefixed {
			pterm.FgBlue.Print("[CB] ")
		}

		for i, op := range in.Operands {
			symbol := fmt.Sprintf("%v", op.Symbol)
			if num, ok := op.Symbol.(byte); ok {
				symbol = fmt.Sprintf("%d", num)
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
	pterm.FgGray.Printf(" (args=%d)\n", len(in.Operands))
}

func CPU(c *cpu.CPU) {
	if Hide {
		return
	}

	r8 := func(b byte) string {
		return pterm.FgYellow.Sprintf("%02x", b)
	}

	r16 := func(w uint16) string {
		return pterm.FgYellow.Sprintf("%04x", w)
	}

	flagStr := func(f cpu.Flag) string {
		if c.Registers.GetFlag(f) {
			return pterm.Green("1")
		} else {
			return pterm.Red("0")
		}
	}

	pterm.FgDarkGray.Printf("  PC: %s · SP: %s\n", r16(c.Registers.PC), r16(c.Registers.SP))
	pterm.FgDarkGray.Printf("  A: %s · F: %s · B: %s · C: %s · D: %s · E: %s · H: %s · L: %s\n", r8(c.Registers.A), r8(c.Registers.F), r8(c.Registers.B), r8(c.Registers.C), r8(c.Registers.D), r8(c.Registers.E), r8(c.Registers.H), r8(c.Registers.L))
	pterm.FgDarkGray.Printf("  Z: %s · N: %s · H: %s · C: %s\n", flagStr(cpu.FlagZ), flagStr(cpu.FlagN), flagStr(cpu.FlagH), flagStr(cpu.FlagC))
}

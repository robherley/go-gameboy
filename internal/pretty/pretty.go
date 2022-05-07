package pretty

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/cpu"
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

func Instruction(pc, sp uint16, opcode byte, in *cpu.Instruction) {
	if Hide {
		return
	}

	isCBPrefixed := opcode == 0xCB

	pterm.NewStyle(pterm.FgBlack, pterm.BgWhite, pterm.Bold).Printf(" %04x | %04x ", pc, sp)
	pterm.FgLightCyan.Printf(" % 3s ", in.Mnemonic)
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
}

func CPU(c *cpu.CPU) {
	if Hide {
		return
	}

	r8 := func(b byte) string {
		return pterm.FgYellow.Sprintf("%02x", b)
	}

	flagStr := func(f cpu.Flag) string {
		if c.Registers.GetFlag(f) {
			return pterm.Green("1")
		} else {
			return pterm.Red("0")
		}
	}

	pterm.FgDarkGray.Printf(" | AF: %s%s BC: %s%s DE: %s%s HL: %s%s ", r8(c.Registers.A), r8(c.Registers.F), r8(c.Registers.B), r8(c.Registers.C), r8(c.Registers.D), r8(c.Registers.E), r8(c.Registers.H), r8(c.Registers.L))
	pterm.FgDarkGray.Printf("| ZNHC: %s%s%s%s", flagStr(cpu.FlagZ), flagStr(cpu.FlagN), flagStr(cpu.FlagH), flagStr(cpu.FlagC))
}

func Interrupt(it cpu.InterruptType) {
	if Hide {
		return
	}

	pterm.NewStyle(pterm.FgBlack, pterm.BgYellow, pterm.Bold).Printf(" ❗ Interrupt: %s  \n", it)
}

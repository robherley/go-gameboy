package debug

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/cpu"
)

var Hide = false

func Cart(cart *cartridge.Cartridge) {
	if Hide {
		return
	}

	headerCheck := fmt.Sprintf("Match (0x%02x) ✅", cart.HeaderChecksum())
	if !cart.IsValidHeaderCheckSum() {
		headerCheck = fmt.Sprintf("Mismatch (0x%02x) ❌", cart.HeaderChecksum())
	}

	globalCheck := fmt.Sprintf("Match (0x%04x) ✅", cart.GlobalChecksum())
	if !cart.IsValidGlobalCheckSum() {
		globalCheck = fmt.Sprintf("Mismatch (0x%04x) ❌", cart.GlobalChecksum())
	}

	rows := []struct {
		label, data string
	}{
		{"Title", cart.TitleString()},
		{"Type", fmt.Sprintf("%s (0x%02x)", cart.CartridgeType().String(), cart.CartridgeType())},
		{"Licensee", cart.LicenseeString()},
		{"Size", fmt.Sprintf("%dK", cart.Size/1024)},
		{"Header Checksum", headerCheck},
		{"Global Checksum", globalCheck},
	}

	for _, row := range rows {
		fmt.Printf("%-16s: %s\n", row.label, row.data)
	}
}

func CPU(c *cpu.CPU) {
	if Hide {
		return
	}

	flagStr := func(f cpu.Flag) string {
		if c.Registers.GetFlag(f) {
			return "1"
		} else {
			return "0"
		}
	}

	fmt.Printf(" | AF: %02X%02X BC: %02X%02X DE: %02X%02X HL: %02X%02X ", c.Registers.A, c.Registers.F, c.Registers.B, c.Registers.C, c.Registers.D, c.Registers.E, c.Registers.H, c.Registers.L)
	fmt.Printf("| ZNHC: %s%s%s%s", flagStr(cpu.FlagZ), flagStr(cpu.FlagN), flagStr(cpu.FlagH), flagStr(cpu.FlagC))

	fmt.Println()
}

func Interrupt(it cpu.InterruptType) {
	if Hide {
		return
	}

	fmt.Println(" ❗ Interrupt:", it)
}

func Instruction(pc, sp uint16, opcode byte, in *cpu.Instruction) {
	if Hide {
		return
	}

	fmt.Printf("%04X - %04X: [%02X]", sp, pc, opcode)
	instructionStr := fmt.Sprintf("%s ", operationName(in.Operation))

	if in.Operands != nil {
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

			instructionStr += symbol

			if i != len(in.Operands)-1 {
				instructionStr += ", "
			}
		}
	}

	fmt.Printf(" %-14s", instructionStr)
}

func operationName(op cpu.Operation) string {
	// not the most efficient, but good enough for debugging rn
	// TODO: add mnemonic to instruction as str const & see if effect perf
	funcName := runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name()
	split := strings.Split(funcName, ".")
	return split[len(split)-1]
}

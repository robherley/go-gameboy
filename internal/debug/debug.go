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

	r8 := func(b byte) string {
		return fmt.Sprintf("%02x", b)
	}

	flagStr := func(f cpu.Flag) string {
		if c.Registers.GetFlag(f) {
			return "1"
		} else {
			return "0"
		}
	}

	fmt.Printf(" | AF: %s%s BC: %s%s DE: %s%s HL: %s%s ", r8(c.Registers.A), r8(c.Registers.F), r8(c.Registers.B), r8(c.Registers.C), r8(c.Registers.D), r8(c.Registers.E), r8(c.Registers.H), r8(c.Registers.L))
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

	fmt.Printf("%04X - %04X: [%02X]", pc, sp, opcode)
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

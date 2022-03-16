package emulator

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/cpu"
)

type Emulator struct {
	CPU *cpu.CPU
}

func New(cart *cartridge.Cartridge) *Emulator {
	return &Emulator{
		CPU: cpu.New(cart),
	}
}

func (e *Emulator) Boot() {
	e.NextTick()
}

func (e *Emulator) NextTick() {
	currentPC := e.CPU.PC

	opcode := e.CPU.Fetch8()

	instr := e.CPU.InstructionForOPCode(opcode)
	if instr == nil {
		err := fmt.Errorf("unknown instruction: 0x%x", opcode)
		panic(err)
	}

	fmt.Printf("%s PC: 0x%x\n", pterm.NewStyle(pterm.BgCyan, pterm.FgBlack).Sprintf("  %s (0x%02x)  ", instr.Mnemonic, opcode), currentPC)

	e.CPU.PC++
	e.NextTick()
}

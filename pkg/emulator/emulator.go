package emulator

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/pretty"
	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/cpu"
	"github.com/robherley/go-dmg/pkg/instructions"
)

type Emulator struct {
	CPU *cpu.CPU
}

func New(cart *cartridge.Cartridge) *Emulator {
	return &Emulator{
		CPU: cpu.New(cart),
	}
}

func (emu *Emulator) Boot() {
	pretty.CPU(emu.CPU)

	for {
		emu.Step()
	}
}

func (emu *Emulator) Step() {
	currentPC := emu.CPU.Registers.PC

	opcode := emu.CPU.Fetch8()
	in := instructions.FromOPCode(opcode, false)
	if in == nil {
		panic(fmt.Errorf("unknown instruction: 0x%x", opcode))
	}
	pretty.Instruction(currentPC, opcode, in)

	ticks := emu.CPU.Process(in)
	pretty.CPU(emu.CPU)

	emu.doTicks(ticks)
}

func (emu *Emulator) doTicks(ticks byte) {
	fmt.Println("  🕓 TODO:", ticks, "ticks")
}

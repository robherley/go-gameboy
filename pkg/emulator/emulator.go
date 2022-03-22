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

func (e *Emulator) Boot() {
	e.Step()
}

func (e *Emulator) Step() {
	currentPC := e.CPU.PC

	opcode := e.CPU.Fetch8()

	in := instructions.FromOPCode(opcode, false)
	if in == nil {
		panic(fmt.Errorf("unknown instruction: 0x%x", opcode))
	}

	pretty.Instruction(currentPC, opcode, in)

	ticks := e.CPU.Process(in)
	e.doTicks(ticks)
	e.Step()
}

func (e *Emulator) doTicks(ticks byte) {
	fmt.Println("\t🕓", ticks, "ticks")
}

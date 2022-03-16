package emulator

import (
	"fmt"

	"github.com/robherley/go-dmg/internal/pretty"
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

	pretty.Instruction(instr.Mnemonic, opcode, currentPC)

	e.CPU.PC++
	e.NextTick()
}

package emulator

import (
	"fmt"

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
	e.NextTick()
}

func (e *Emulator) NextTick() {
	currentPC := e.CPU.PC

	opcode := e.CPU.Fetch8()

	in := instructions.FromOPCode(opcode, false)
	if in == nil {
		err := fmt.Errorf("unknown instruction: 0x%x", opcode)
		panic(err)
	}

	fmt.Printf("PC 0x%x | Instruction (0x%02x): %+v\n", currentPC, opcode, in)

	panic("not implemented")
}

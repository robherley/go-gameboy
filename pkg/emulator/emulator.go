package emulator

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/cpu"
)

type Emulator struct {
	Cartridge *cartridge.Cartridge
	CPU       *cpu.CPU
}

func New(cart *cartridge.Cartridge) *Emulator {
	return &Emulator{
		Cartridge: cart,
		CPU:       cpu.New(),
	}
}

func (e *Emulator) Boot() {
	fmt.Println("beep boop")
	e.NextTick()
}

func (e *Emulator) NextTick() {
	opcode := e.BusRead(e.CPU.PC)

	instr := e.CPU.InstructionForOPCode(opcode)
	if instr == nil {
		err := fmt.Errorf("unknown instruction: 0x%x", opcode)
		panic(err)
	}

	fmt.Printf("%s | PC: 0x%x\n", instr.Mnemonic, e.CPU.PC)

	e.CPU.PC++
	e.NextTick()
}

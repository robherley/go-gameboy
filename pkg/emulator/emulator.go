package emulator

import (
	"fmt"

	"github.com/robherley/go-gameboy/internal/pretty"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/cpu"
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
		// TODO: clock
	}
}

func (emu *Emulator) Step() {
	currentPC := emu.CPU.Registers.PC
	currentSP := emu.CPU.Registers.SP

	pretty.Hide = true

	if !emu.CPU.Halted {
		opcode, instruction, cb := emu.CPU.NextInstruction()
		mnem := instruction.Mnemonic
		if cb {
			mnem = "CB"
		}
		fmt.Printf("%04X %s - %04X \n", currentPC, mnem, opcode)
		pretty.Instruction(currentPC, currentSP, opcode, instruction)
		pretty.CPU(emu.CPU)
		fmt.Println()
		emu.CPU.MMU.DebugSerial()
		emu.CPU.Process(instruction)
	} else {
		// interrupt was requested
		if emu.CPU.Interrupt.Requested() {
			emu.CPU.Halted = false
		}
	}

	emu.CPU.HandleInterrupts()
}

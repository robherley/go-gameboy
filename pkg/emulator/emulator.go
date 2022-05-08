package emulator

import (
	"github.com/robherley/go-gameboy/internal/debug"
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
	debug.CPU(emu.CPU)

	for {
		emu.Step()
		// TODO: clock
	}
}

func (emu *Emulator) Step() {
	currentPC := emu.CPU.Registers.PC
	currentSP := emu.CPU.Registers.SP

	if !emu.CPU.Halted {
		opcode, instruction := emu.CPU.NextInstruction()
		debug.Instruction(currentPC, currentSP, opcode, instruction)
		debug.CPU(emu.CPU)
		emu.CPU.MMU.DebugSerial()
		instruction.Operation(emu.CPU, instruction.Operands)
	} else {
		// interrupt was requested
		if emu.CPU.Interrupt.IsRequested() {
			emu.CPU.Halted = false
		}
	}

	emu.CPU.HandleInterrupts()
}

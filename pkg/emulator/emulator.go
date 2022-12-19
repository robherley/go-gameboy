package emulator

import (
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
	for {
		emu.Step()
		// TODO: clock
	}
}

func (emu *Emulator) Step() {
	emu.CPU.HandleInterrupts()

	if !emu.CPU.Halted {
		// fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", emu.CPU.Registers.A, emu.CPU.Registers.F, emu.CPU.Registers.B, emu.CPU.Registers.C, emu.CPU.Registers.D, emu.CPU.Registers.E, emu.CPU.Registers.H, emu.CPU.Registers.L, emu.CPU.Registers.SP, emu.CPU.Registers.PC, emu.CPU.MMU.Read8(emu.CPU.Registers.PC), emu.CPU.MMU.Read8(emu.CPU.Registers.PC+1), emu.CPU.MMU.Read8(emu.CPU.Registers.PC+2), emu.CPU.MMU.Read8(emu.CPU.Registers.PC+3))
		_, instruction := emu.CPU.NextInstruction()
		// debug.Instruction(currentPC, currentSP, opcode, instruction)
		// emu.CPU.MMU.DebugMem()

		// debug.CPU(emu.CPU)
		emu.CPU.MMU.DebugSerial()
		instruction.Execute(emu.CPU)
	} else {
		// basically a noop
		emu.CPU.EmulateCycles(1)
		// interrupt was requested
		if emu.CPU.Interrupt.Requested() {
			emu.CPU.Halted = false
		}
	}
}

package emulator

import (
	"github.com/robherley/go-gameboy/internal/pretty"
	"github.com/robherley/go-gameboy/pkg/cartridge"
	"github.com/robherley/go-gameboy/pkg/cpu"
	errs "github.com/robherley/go-gameboy/pkg/errors"
	"github.com/robherley/go-gameboy/pkg/instructions"
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

	if !emu.CPU.IsHalted {
		opcode := emu.CPU.Fetch8()
		isCBPrexied := opcode == 0xCB
		if isCBPrexied {
			// cb-prefixed instructions have opcode on next fetch
			opcode = emu.CPU.Fetch8()
		}

		instruction := instructions.FromOPCode(opcode, isCBPrexied)
		if instruction == nil {
			panic(errs.NewUnknownOPCodeError(opcode))
		}
		pretty.Instruction(currentPC, opcode, instruction, isCBPrexied)

		emu.CPU.Process(instruction)
		pretty.CPU(emu.CPU)
	} else {
		// TODO emulate cycles

		// interrupt was requested
		if emu.CPU.IF != 0 {
			emu.CPU.IsHalted = false
		}
	}

	// if master interrupt enabled, handle any interrupts
	if emu.CPU.IME {
		it := emu.CPU.HandleInterrupts()
		if it != nil {
			pretty.Interrupt(*it)
		}
		emu.CPU.EnablingIME = false
	}

	// EI was called, enable master interrupt for next cycle
	if emu.CPU.EnablingIME {
		emu.CPU.IME = true
	}
}

package cpu

// https://gbdev.io/gb-opcodes/optables/
// http://gameboy.mongenel.com/dmg/opcodes.html

type Instruction struct {
	// Mnemonic as string
	Mnemonic
	// Function to handle instruction, returns number of system clock ticks
	Handle func(c *CPU) byte
}

func (c *CPU) InstructionForOPCode(opcode byte) *Instruction {
	if int(opcode) > len(unprefixed) {
		return nil
	}
	return unprefixed[opcode]
}

var unprefixed = [...]*Instruction{
	0x00: {
		NOP,
		func(c *CPU) byte {
			return 4
		},
	},
}

// var cbPrefixed = [...]*Instruction{}

package cpu

// https://gbdev.io/gb-opcodes/optables/
// https://gbdev.io/pandocs/CPU_Instruction_Set.html
// http://gameboy.mongenel.com/dmg/opcodes.html

type Instruction struct {
	// Mnemonic for an instruction, used for debugging
	Mnemonic
	// Operands as a string, used for debugging
	Operands []string
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
		nil,
		func(c *CPU) byte {
			return 4
		},
	},
	0x01: {
		LD,
		[]string{"BC", "d16"},
		func(c *CPU) byte {
			data := c.Fetch16()
			c.SetBC(data)
			return 12
		},
	},
	0x02: {
		LD,
		[]string{"(BC)", "A"},
		func(c *CPU) byte {
			c.MMU.write8(c.GetBC(), c.A)
			return 8
		},
	},
	0x03: {
		INC,
		[]string{"BC"},
		func(c *CPU) byte {
			c.SetBC(c.GetBC() + 1)
			return 8
		},
	},
	0x04: {
		INC,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x05: {
		DEC,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x06: {
		LD,
		[]string{"B", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x07: {
		RLCA,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x08: {
		LD,
		[]string{"(a16)", "SP"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x09: {
		ADD,
		[]string{"HL", "BC"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0A: {
		LD,
		[]string{"A", "(BC)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0B: {
		DEC,
		[]string{"BC"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0C: {
		INC,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0D: {
		DEC,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0E: {
		LD,
		[]string{"C", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0F: {
		RRCA,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x10: {
		STOP,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x11: {
		LD,
		[]string{"DE", "d16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x12: {
		LD,
		[]string{"(DE)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x13: {
		INC,
		[]string{"DE"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x14: {
		INC,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x15: {
		DEC,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x16: {
		LD,
		[]string{"D", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x17: {
		RLA,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x18: {
		JR,
		[]string{"r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x19: {
		ADD,
		[]string{"HL", "DE"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1A: {
		LD,
		[]string{"A", "(DE)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1B: {
		DEC,
		[]string{"DE"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1C: {
		INC,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1D: {
		DEC,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1E: {
		LD,
		[]string{"E", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1F: {
		RRA,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x20: {
		JR,
		[]string{"NZ", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x21: {
		LD,
		[]string{"HL", "d16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x22: {
		LD,
		[]string{"(HL)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x23: {
		INC,
		[]string{"HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x24: {
		INC,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x25: {
		DEC,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x26: {
		LD,
		[]string{"H", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x27: {
		DAA,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x28: {
		JR,
		[]string{"Z", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x29: {
		ADD,
		[]string{"HL", "HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2A: {
		LD,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2B: {
		DEC,
		[]string{"HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2C: {
		INC,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2D: {
		DEC,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2E: {
		LD,
		[]string{"L", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2F: {
		CPL,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x30: {
		JR,
		[]string{"NC", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x31: {
		LD,
		[]string{"SP", "d16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x32: {
		LD,
		[]string{"(HL)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x33: {
		INC,
		[]string{"SP"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x34: {
		INC,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x35: {
		DEC,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x36: {
		LD,
		[]string{"(HL)", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x37: {
		SCF,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x38: {
		JR,
		[]string{"C", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x39: {
		ADD,
		[]string{"HL", "SP"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3A: {
		LD,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3B: {
		DEC,
		[]string{"SP"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3C: {
		INC,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3D: {
		DEC,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3E: {
		LD,
		[]string{"A", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3F: {
		CCF,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x40: {
		LD,
		[]string{"B", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x41: {
		LD,
		[]string{"B", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x42: {
		LD,
		[]string{"B", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x43: {
		LD,
		[]string{"B", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x44: {
		LD,
		[]string{"B", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x45: {
		LD,
		[]string{"B", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x46: {
		LD,
		[]string{"B", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x47: {
		LD,
		[]string{"B", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x48: {
		LD,
		[]string{"C", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x49: {
		LD,
		[]string{"C", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4A: {
		LD,
		[]string{"C", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4B: {
		LD,
		[]string{"C", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4C: {
		LD,
		[]string{"C", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4D: {
		LD,
		[]string{"C", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4E: {
		LD,
		[]string{"C", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4F: {
		LD,
		[]string{"C", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x50: {
		LD,
		[]string{"D", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x51: {
		LD,
		[]string{"D", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x52: {
		LD,
		[]string{"D", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x53: {
		LD,
		[]string{"D", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x54: {
		LD,
		[]string{"D", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x55: {
		LD,
		[]string{"D", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x56: {
		LD,
		[]string{"D", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x57: {
		LD,
		[]string{"D", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x58: {
		LD,
		[]string{"E", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x59: {
		LD,
		[]string{"E", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5A: {
		LD,
		[]string{"E", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5B: {
		LD,
		[]string{"E", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5C: {
		LD,
		[]string{"E", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5D: {
		LD,
		[]string{"E", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5E: {
		LD,
		[]string{"E", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5F: {
		LD,
		[]string{"E", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x60: {
		LD,
		[]string{"H", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x61: {
		LD,
		[]string{"H", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x62: {
		LD,
		[]string{"H", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x63: {
		LD,
		[]string{"H", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x64: {
		LD,
		[]string{"H", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x65: {
		LD,
		[]string{"H", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x66: {
		LD,
		[]string{"H", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x67: {
		LD,
		[]string{"H", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x68: {
		LD,
		[]string{"L", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x69: {
		LD,
		[]string{"L", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6A: {
		LD,
		[]string{"L", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6B: {
		LD,
		[]string{"L", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6C: {
		LD,
		[]string{"L", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6D: {
		LD,
		[]string{"L", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6E: {
		LD,
		[]string{"L", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6F: {
		LD,
		[]string{"L", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x70: {
		LD,
		[]string{"(HL)", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x71: {
		LD,
		[]string{"(HL)", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x72: {
		LD,
		[]string{"(HL)", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x73: {
		LD,
		[]string{"(HL)", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x74: {
		LD,
		[]string{"(HL)", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x75: {
		LD,
		[]string{"(HL)", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x76: {
		HALT,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x77: {
		LD,
		[]string{"(HL)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x78: {
		LD,
		[]string{"A", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x79: {
		LD,
		[]string{"A", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7A: {
		LD,
		[]string{"A", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7B: {
		LD,
		[]string{"A", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7C: {
		LD,
		[]string{"A", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7D: {
		LD,
		[]string{"A", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7E: {
		LD,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7F: {
		LD,
		[]string{"A", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x80: {
		ADD,
		[]string{"A", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x81: {
		ADD,
		[]string{"A", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x82: {
		ADD,
		[]string{"A", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x83: {
		ADD,
		[]string{"A", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x84: {
		ADD,
		[]string{"A", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x85: {
		ADD,
		[]string{"A", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x86: {
		ADD,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x87: {
		ADD,
		[]string{"A", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x88: {
		ADC,
		[]string{"A", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x89: {
		ADC,
		[]string{"A", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8A: {
		ADC,
		[]string{"A", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8B: {
		ADC,
		[]string{"A", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8C: {
		ADC,
		[]string{"A", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8D: {
		ADC,
		[]string{"A", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8E: {
		ADC,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8F: {
		ADC,
		[]string{"A", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x90: {
		SUB,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x91: {
		SUB,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x92: {
		SUB,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x93: {
		SUB,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x94: {
		SUB,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x95: {
		SUB,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x96: {
		SUB,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x97: {
		SUB,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x98: {
		SBC,
		[]string{"A", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x99: {
		SBC,
		[]string{"A", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9A: {
		SBC,
		[]string{"A", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9B: {
		SBC,
		[]string{"A", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9C: {
		SBC,
		[]string{"A", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9D: {
		SBC,
		[]string{"A", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9E: {
		SBC,
		[]string{"A", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9F: {
		SBC,
		[]string{"A", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA0: {
		AND,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA1: {
		AND,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA2: {
		AND,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA3: {
		AND,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA4: {
		AND,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA5: {
		AND,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA6: {
		AND,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA7: {
		AND,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA8: {
		XOR,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA9: {
		XOR,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAA: {
		XOR,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAB: {
		XOR,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAC: {
		XOR,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAD: {
		XOR,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAE: {
		XOR,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAF: {
		XOR,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB0: {
		OR,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB1: {
		OR,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB2: {
		OR,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB3: {
		OR,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB4: {
		OR,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB5: {
		OR,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB6: {
		OR,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB7: {
		OR,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB8: {
		CP,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB9: {
		CP,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBA: {
		CP,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBB: {
		CP,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBC: {
		CP,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBD: {
		CP,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBE: {
		CP,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBF: {
		CP,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC0: {
		RET,
		[]string{"NZ"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC1: {
		POP,
		[]string{"BC"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC2: {
		JP,
		[]string{"NZ", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC3: {
		JP,
		[]string{"a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC4: {
		CALL,
		[]string{"NZ", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC5: {
		PUSH,
		[]string{"BC"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC6: {
		ADD,
		[]string{"A", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC7: {
		RST,
		[]string{"00H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC8: {
		RET,
		[]string{"Z"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC9: {
		RET,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCA: {
		JP,
		[]string{"Z", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCB: {
		PREFIX,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCC: {
		CALL,
		[]string{"Z", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCD: {
		CALL,
		[]string{"a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCE: {
		ADC,
		[]string{"A", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCF: {
		RST,
		[]string{"08H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD0: {
		RET,
		[]string{"NC"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD1: {
		POP,
		[]string{"DE"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD2: {
		JP,
		[]string{"NC", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD3: {
		ILLEGAL_D3,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD4: {
		CALL,
		[]string{"NC", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD5: {
		PUSH,
		[]string{"DE"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD6: {
		SUB,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD7: {
		RST,
		[]string{"10H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD8: {
		RET,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD9: {
		RETI,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDA: {
		JP,
		[]string{"C", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDB: {
		ILLEGAL_DB,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDC: {
		CALL,
		[]string{"C", "a16"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDD: {
		ILLEGAL_DD,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDE: {
		SBC,
		[]string{"A", "d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDF: {
		RST,
		[]string{"18H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE0: {
		LDH,
		[]string{"(a8)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE1: {
		POP,
		[]string{"HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE2: {
		LD,
		[]string{"(C)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE3: {
		ILLEGAL_E3,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE4: {
		ILLEGAL_E4,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE5: {
		PUSH,
		[]string{"HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE6: {
		AND,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE7: {
		RST,
		[]string{"20H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE8: {
		ADD,
		[]string{"SP", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE9: {
		JP,
		[]string{"HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEA: {
		LD,
		[]string{"(a16)", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEB: {
		ILLEGAL_EB,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEC: {
		ILLEGAL_EC,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xED: {
		ILLEGAL_ED,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEE: {
		XOR,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEF: {
		RST,
		[]string{"28H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF0: {
		LDH,
		[]string{"A", "(a8)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF1: {
		POP,
		[]string{"AF"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF2: {
		LD,
		[]string{"A", "(C)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF3: {
		DI,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF4: {
		ILLEGAL_F4,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF5: {
		PUSH,
		[]string{"AF"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF6: {
		OR,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF7: {
		RST,
		[]string{"30H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF8: {
		LD,
		[]string{"HL", "SP", "r8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF9: {
		LD,
		[]string{"SP", "HL"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFA: {
		LD,
		[]string{"A", "(a16)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFB: {
		EI,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFC: {
		ILLEGAL_FC,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFD: {
		ILLEGAL_FD,
		nil,
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFE: {
		CP,
		[]string{"d8"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFF: {
		RST,
		[]string{"38H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
}

var cbPrefixed = [...]*Instruction{
	0x00: {
		RLC,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x01: {
		RLC,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x02: {
		RLC,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x03: {
		RLC,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x04: {
		RLC,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x05: {
		RLC,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x06: {
		RLC,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x07: {
		RLC,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x08: {
		RRC,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x09: {
		RRC,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0A: {
		RRC,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0B: {
		RRC,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0C: {
		RRC,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0D: {
		RRC,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0E: {
		RRC,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x0F: {
		RRC,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x10: {
		RL,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x11: {
		RL,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x12: {
		RL,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x13: {
		RL,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x14: {
		RL,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x15: {
		RL,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x16: {
		RL,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x17: {
		RL,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x18: {
		RR,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x19: {
		RR,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1A: {
		RR,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1B: {
		RR,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1C: {
		RR,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1D: {
		RR,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1E: {
		RR,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x1F: {
		RR,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x20: {
		SLA,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x21: {
		SLA,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x22: {
		SLA,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x23: {
		SLA,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x24: {
		SLA,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x25: {
		SLA,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x26: {
		SLA,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x27: {
		SLA,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x28: {
		SRA,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x29: {
		SRA,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2A: {
		SRA,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2B: {
		SRA,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2C: {
		SRA,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2D: {
		SRA,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2E: {
		SRA,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x2F: {
		SRA,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x30: {
		SWAP,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x31: {
		SWAP,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x32: {
		SWAP,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x33: {
		SWAP,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x34: {
		SWAP,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x35: {
		SWAP,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x36: {
		SWAP,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x37: {
		SWAP,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x38: {
		SRL,
		[]string{"B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x39: {
		SRL,
		[]string{"C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3A: {
		SRL,
		[]string{"D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3B: {
		SRL,
		[]string{"E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3C: {
		SRL,
		[]string{"H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3D: {
		SRL,
		[]string{"L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3E: {
		SRL,
		[]string{"(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x3F: {
		SRL,
		[]string{"A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x40: {
		BIT,
		[]string{"0", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x41: {
		BIT,
		[]string{"0", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x42: {
		BIT,
		[]string{"0", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x43: {
		BIT,
		[]string{"0", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x44: {
		BIT,
		[]string{"0", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x45: {
		BIT,
		[]string{"0", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x46: {
		BIT,
		[]string{"0", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x47: {
		BIT,
		[]string{"0", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x48: {
		BIT,
		[]string{"1", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x49: {
		BIT,
		[]string{"1", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4A: {
		BIT,
		[]string{"1", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4B: {
		BIT,
		[]string{"1", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4C: {
		BIT,
		[]string{"1", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4D: {
		BIT,
		[]string{"1", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4E: {
		BIT,
		[]string{"1", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x4F: {
		BIT,
		[]string{"1", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x50: {
		BIT,
		[]string{"2", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x51: {
		BIT,
		[]string{"2", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x52: {
		BIT,
		[]string{"2", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x53: {
		BIT,
		[]string{"2", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x54: {
		BIT,
		[]string{"2", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x55: {
		BIT,
		[]string{"2", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x56: {
		BIT,
		[]string{"2", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x57: {
		BIT,
		[]string{"2", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x58: {
		BIT,
		[]string{"3", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x59: {
		BIT,
		[]string{"3", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5A: {
		BIT,
		[]string{"3", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5B: {
		BIT,
		[]string{"3", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5C: {
		BIT,
		[]string{"3", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5D: {
		BIT,
		[]string{"3", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5E: {
		BIT,
		[]string{"3", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x5F: {
		BIT,
		[]string{"3", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x60: {
		BIT,
		[]string{"4", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x61: {
		BIT,
		[]string{"4", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x62: {
		BIT,
		[]string{"4", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x63: {
		BIT,
		[]string{"4", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x64: {
		BIT,
		[]string{"4", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x65: {
		BIT,
		[]string{"4", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x66: {
		BIT,
		[]string{"4", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x67: {
		BIT,
		[]string{"4", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x68: {
		BIT,
		[]string{"5", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x69: {
		BIT,
		[]string{"5", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6A: {
		BIT,
		[]string{"5", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6B: {
		BIT,
		[]string{"5", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6C: {
		BIT,
		[]string{"5", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6D: {
		BIT,
		[]string{"5", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6E: {
		BIT,
		[]string{"5", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x6F: {
		BIT,
		[]string{"5", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x70: {
		BIT,
		[]string{"6", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x71: {
		BIT,
		[]string{"6", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x72: {
		BIT,
		[]string{"6", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x73: {
		BIT,
		[]string{"6", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x74: {
		BIT,
		[]string{"6", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x75: {
		BIT,
		[]string{"6", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x76: {
		BIT,
		[]string{"6", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x77: {
		BIT,
		[]string{"6", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x78: {
		BIT,
		[]string{"7", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x79: {
		BIT,
		[]string{"7", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7A: {
		BIT,
		[]string{"7", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7B: {
		BIT,
		[]string{"7", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7C: {
		BIT,
		[]string{"7", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7D: {
		BIT,
		[]string{"7", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7E: {
		BIT,
		[]string{"7", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x7F: {
		BIT,
		[]string{"7", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x80: {
		RES,
		[]string{"0", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x81: {
		RES,
		[]string{"0", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x82: {
		RES,
		[]string{"0", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x83: {
		RES,
		[]string{"0", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x84: {
		RES,
		[]string{"0", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x85: {
		RES,
		[]string{"0", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x86: {
		RES,
		[]string{"0", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x87: {
		RES,
		[]string{"0", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x88: {
		RES,
		[]string{"1", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x89: {
		RES,
		[]string{"1", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8A: {
		RES,
		[]string{"1", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8B: {
		RES,
		[]string{"1", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8C: {
		RES,
		[]string{"1", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8D: {
		RES,
		[]string{"1", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8E: {
		RES,
		[]string{"1", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x8F: {
		RES,
		[]string{"1", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x90: {
		RES,
		[]string{"2", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x91: {
		RES,
		[]string{"2", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x92: {
		RES,
		[]string{"2", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x93: {
		RES,
		[]string{"2", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x94: {
		RES,
		[]string{"2", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x95: {
		RES,
		[]string{"2", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x96: {
		RES,
		[]string{"2", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x97: {
		RES,
		[]string{"2", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x98: {
		RES,
		[]string{"3", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x99: {
		RES,
		[]string{"3", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9A: {
		RES,
		[]string{"3", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9B: {
		RES,
		[]string{"3", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9C: {
		RES,
		[]string{"3", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9D: {
		RES,
		[]string{"3", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9E: {
		RES,
		[]string{"3", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0x9F: {
		RES,
		[]string{"3", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA0: {
		RES,
		[]string{"4", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA1: {
		RES,
		[]string{"4", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA2: {
		RES,
		[]string{"4", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA3: {
		RES,
		[]string{"4", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA4: {
		RES,
		[]string{"4", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA5: {
		RES,
		[]string{"4", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA6: {
		RES,
		[]string{"4", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA7: {
		RES,
		[]string{"4", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA8: {
		RES,
		[]string{"5", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xA9: {
		RES,
		[]string{"5", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAA: {
		RES,
		[]string{"5", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAB: {
		RES,
		[]string{"5", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAC: {
		RES,
		[]string{"5", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAD: {
		RES,
		[]string{"5", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAE: {
		RES,
		[]string{"5", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xAF: {
		RES,
		[]string{"5", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB0: {
		RES,
		[]string{"6", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB1: {
		RES,
		[]string{"6", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB2: {
		RES,
		[]string{"6", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB3: {
		RES,
		[]string{"6", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB4: {
		RES,
		[]string{"6", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB5: {
		RES,
		[]string{"6", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB6: {
		RES,
		[]string{"6", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB7: {
		RES,
		[]string{"6", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB8: {
		RES,
		[]string{"7", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xB9: {
		RES,
		[]string{"7", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBA: {
		RES,
		[]string{"7", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBB: {
		RES,
		[]string{"7", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBC: {
		RES,
		[]string{"7", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBD: {
		RES,
		[]string{"7", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBE: {
		RES,
		[]string{"7", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xBF: {
		RES,
		[]string{"7", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC0: {
		SET,
		[]string{"0", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC1: {
		SET,
		[]string{"0", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC2: {
		SET,
		[]string{"0", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC3: {
		SET,
		[]string{"0", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC4: {
		SET,
		[]string{"0", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC5: {
		SET,
		[]string{"0", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC6: {
		SET,
		[]string{"0", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC7: {
		SET,
		[]string{"0", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC8: {
		SET,
		[]string{"1", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xC9: {
		SET,
		[]string{"1", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCA: {
		SET,
		[]string{"1", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCB: {
		SET,
		[]string{"1", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCC: {
		SET,
		[]string{"1", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCD: {
		SET,
		[]string{"1", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCE: {
		SET,
		[]string{"1", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xCF: {
		SET,
		[]string{"1", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD0: {
		SET,
		[]string{"2", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD1: {
		SET,
		[]string{"2", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD2: {
		SET,
		[]string{"2", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD3: {
		SET,
		[]string{"2", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD4: {
		SET,
		[]string{"2", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD5: {
		SET,
		[]string{"2", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD6: {
		SET,
		[]string{"2", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD7: {
		SET,
		[]string{"2", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD8: {
		SET,
		[]string{"3", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xD9: {
		SET,
		[]string{"3", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDA: {
		SET,
		[]string{"3", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDB: {
		SET,
		[]string{"3", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDC: {
		SET,
		[]string{"3", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDD: {
		SET,
		[]string{"3", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDE: {
		SET,
		[]string{"3", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xDF: {
		SET,
		[]string{"3", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE0: {
		SET,
		[]string{"4", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE1: {
		SET,
		[]string{"4", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE2: {
		SET,
		[]string{"4", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE3: {
		SET,
		[]string{"4", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE4: {
		SET,
		[]string{"4", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE5: {
		SET,
		[]string{"4", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE6: {
		SET,
		[]string{"4", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE7: {
		SET,
		[]string{"4", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE8: {
		SET,
		[]string{"5", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xE9: {
		SET,
		[]string{"5", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEA: {
		SET,
		[]string{"5", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEB: {
		SET,
		[]string{"5", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEC: {
		SET,
		[]string{"5", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xED: {
		SET,
		[]string{"5", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEE: {
		SET,
		[]string{"5", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xEF: {
		SET,
		[]string{"5", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF0: {
		SET,
		[]string{"6", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF1: {
		SET,
		[]string{"6", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF2: {
		SET,
		[]string{"6", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF3: {
		SET,
		[]string{"6", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF4: {
		SET,
		[]string{"6", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF5: {
		SET,
		[]string{"6", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF6: {
		SET,
		[]string{"6", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF7: {
		SET,
		[]string{"6", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF8: {
		SET,
		[]string{"7", "B"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xF9: {
		SET,
		[]string{"7", "C"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFA: {
		SET,
		[]string{"7", "D"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFB: {
		SET,
		[]string{"7", "E"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFC: {
		SET,
		[]string{"7", "H"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFD: {
		SET,
		[]string{"7", "L"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFE: {
		SET,
		[]string{"7", "(HL)"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
	0xFF: {
		SET,
		[]string{"7", "A"},
		func(c *CPU) byte {
			panic("not implemented")
		},
	},
}

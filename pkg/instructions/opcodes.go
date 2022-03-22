package instructions

// https://gbdev.io/gb-opcodes/optables/
// instructions generated from: https://gbdev.io/gb-opcodes/Opcodes.json
// script: https://gist.github.com/robherley/836369cbd8eb73a286d017626b8376c1

func FromOPCode(code byte, cbprefix bool) *Instruction {
	var mapping map[byte]Instruction
	if cbprefix {
		mapping = cbprefixed
	} else {
		mapping = unprefixed
	}

	if val, ok := mapping[code]; ok {
		return &val
	}

	return nil
}

var unprefixed = map[byte]Instruction{
	0x00: {
		NOP,
		nil,
	},
	0x01: {
		LD,
		Operands{BC, D16},
	},
	0x02: {
		LD,
		Operands{Deref(BC), A},
	},
	0x03: {
		INC,
		Operands{BC},
	},
	0x04: {
		INC,
		Operands{B},
	},
	0x05: {
		DEC,
		Operands{B},
	},
	0x06: {
		LD,
		Operands{B, D8},
	},
	0x07: {
		RLCA,
		nil,
	},
	0x08: {
		LD,
		Operands{Deref(A16), SP},
	},
	0x09: {
		ADD,
		Operands{HL, BC},
	},
	0x0A: {
		LD,
		Operands{A, Deref(BC)},
	},
	0x0B: {
		DEC,
		Operands{BC},
	},
	0x0C: {
		INC,
		Operands{C},
	},
	0x0D: {
		DEC,
		Operands{C},
	},
	0x0E: {
		LD,
		Operands{C, D8},
	},
	0x0F: {
		RRCA,
		nil,
	},
	0x10: {
		STOP,
		Operands{D8},
	},
	0x11: {
		LD,
		Operands{DE, D16},
	},
	0x12: {
		LD,
		Operands{Deref(DE), A},
	},
	0x13: {
		INC,
		Operands{DE},
	},
	0x14: {
		INC,
		Operands{D},
	},
	0x15: {
		DEC,
		Operands{D},
	},
	0x16: {
		LD,
		Operands{D, D8},
	},
	0x17: {
		RLA,
		nil,
	},
	0x18: {
		JR,
		Operands{R8},
	},
	0x19: {
		ADD,
		Operands{HL, DE},
	},
	0x1A: {
		LD,
		Operands{A, Deref(DE)},
	},
	0x1B: {
		DEC,
		Operands{DE},
	},
	0x1C: {
		INC,
		Operands{E},
	},
	0x1D: {
		DEC,
		Operands{E},
	},
	0x1E: {
		LD,
		Operands{E, D8},
	},
	0x1F: {
		RRA,
		nil,
	},
	0x20: {
		JR,
		Operands{NZ, R8},
	},
	0x21: {
		LD,
		Operands{HL, D16},
	},
	0x22: {
		LD,
		Operands{Deref(Inc(HL)), A},
	},
	0x23: {
		INC,
		Operands{HL},
	},
	0x24: {
		INC,
		Operands{H},
	},
	0x25: {
		DEC,
		Operands{H},
	},
	0x26: {
		LD,
		Operands{H, D8},
	},
	0x27: {
		DAA,
		nil,
	},
	0x28: {
		JR,
		Operands{Z, R8},
	},
	0x29: {
		ADD,
		Operands{HL, HL},
	},
	0x2A: {
		LD,
		Operands{A, Deref(Inc(HL))},
	},
	0x2B: {
		DEC,
		Operands{HL},
	},
	0x2C: {
		INC,
		Operands{L},
	},
	0x2D: {
		DEC,
		Operands{L},
	},
	0x2E: {
		LD,
		Operands{L, D8},
	},
	0x2F: {
		CPL,
		nil,
	},
	0x30: {
		JR,
		Operands{NC, R8},
	},
	0x31: {
		LD,
		Operands{SP, D16},
	},
	0x32: {
		LD,
		Operands{Deref(Dec(HL)), A},
	},
	0x33: {
		INC,
		Operands{SP},
	},
	0x34: {
		INC,
		Operands{Deref(HL)},
	},
	0x35: {
		DEC,
		Operands{Deref(HL)},
	},
	0x36: {
		LD,
		Operands{Deref(HL), D8},
	},
	0x37: {
		SCF,
		nil,
	},
	0x38: {
		JR,
		Operands{C, R8},
	},
	0x39: {
		ADD,
		Operands{HL, SP},
	},
	0x3A: {
		LD,
		Operands{A, Deref(Dec(HL))},
	},
	0x3B: {
		DEC,
		Operands{SP},
	},
	0x3C: {
		INC,
		Operands{A},
	},
	0x3D: {
		DEC,
		Operands{A},
	},
	0x3E: {
		LD,
		Operands{A, D8},
	},
	0x3F: {
		CCF,
		nil,
	},
	0x40: {
		LD,
		Operands{B, B},
	},
	0x41: {
		LD,
		Operands{B, C},
	},
	0x42: {
		LD,
		Operands{B, D},
	},
	0x43: {
		LD,
		Operands{B, E},
	},
	0x44: {
		LD,
		Operands{B, H},
	},
	0x45: {
		LD,
		Operands{B, L},
	},
	0x46: {
		LD,
		Operands{B, Deref(HL)},
	},
	0x47: {
		LD,
		Operands{B, A},
	},
	0x48: {
		LD,
		Operands{C, B},
	},
	0x49: {
		LD,
		Operands{C, C},
	},
	0x4A: {
		LD,
		Operands{C, D},
	},
	0x4B: {
		LD,
		Operands{C, E},
	},
	0x4C: {
		LD,
		Operands{C, H},
	},
	0x4D: {
		LD,
		Operands{C, L},
	},
	0x4E: {
		LD,
		Operands{C, Deref(HL)},
	},
	0x4F: {
		LD,
		Operands{C, A},
	},
	0x50: {
		LD,
		Operands{D, B},
	},
	0x51: {
		LD,
		Operands{D, C},
	},
	0x52: {
		LD,
		Operands{D, D},
	},
	0x53: {
		LD,
		Operands{D, E},
	},
	0x54: {
		LD,
		Operands{D, H},
	},
	0x55: {
		LD,
		Operands{D, L},
	},
	0x56: {
		LD,
		Operands{D, Deref(HL)},
	},
	0x57: {
		LD,
		Operands{D, A},
	},
	0x58: {
		LD,
		Operands{E, B},
	},
	0x59: {
		LD,
		Operands{E, C},
	},
	0x5A: {
		LD,
		Operands{E, D},
	},
	0x5B: {
		LD,
		Operands{E, E},
	},
	0x5C: {
		LD,
		Operands{E, H},
	},
	0x5D: {
		LD,
		Operands{E, L},
	},
	0x5E: {
		LD,
		Operands{E, Deref(HL)},
	},
	0x5F: {
		LD,
		Operands{E, A},
	},
	0x60: {
		LD,
		Operands{H, B},
	},
	0x61: {
		LD,
		Operands{H, C},
	},
	0x62: {
		LD,
		Operands{H, D},
	},
	0x63: {
		LD,
		Operands{H, E},
	},
	0x64: {
		LD,
		Operands{H, H},
	},
	0x65: {
		LD,
		Operands{H, L},
	},
	0x66: {
		LD,
		Operands{H, Deref(HL)},
	},
	0x67: {
		LD,
		Operands{H, A},
	},
	0x68: {
		LD,
		Operands{L, B},
	},
	0x69: {
		LD,
		Operands{L, C},
	},
	0x6A: {
		LD,
		Operands{L, D},
	},
	0x6B: {
		LD,
		Operands{L, E},
	},
	0x6C: {
		LD,
		Operands{L, H},
	},
	0x6D: {
		LD,
		Operands{L, L},
	},
	0x6E: {
		LD,
		Operands{L, Deref(HL)},
	},
	0x6F: {
		LD,
		Operands{L, A},
	},
	0x70: {
		LD,
		Operands{Deref(HL), B},
	},
	0x71: {
		LD,
		Operands{Deref(HL), C},
	},
	0x72: {
		LD,
		Operands{Deref(HL), D},
	},
	0x73: {
		LD,
		Operands{Deref(HL), E},
	},
	0x74: {
		LD,
		Operands{Deref(HL), H},
	},
	0x75: {
		LD,
		Operands{Deref(HL), L},
	},
	0x76: {
		HALT,
		nil,
	},
	0x77: {
		LD,
		Operands{Deref(HL), A},
	},
	0x78: {
		LD,
		Operands{A, B},
	},
	0x79: {
		LD,
		Operands{A, C},
	},
	0x7A: {
		LD,
		Operands{A, D},
	},
	0x7B: {
		LD,
		Operands{A, E},
	},
	0x7C: {
		LD,
		Operands{A, H},
	},
	0x7D: {
		LD,
		Operands{A, L},
	},
	0x7E: {
		LD,
		Operands{A, Deref(HL)},
	},
	0x7F: {
		LD,
		Operands{A, A},
	},
	0x80: {
		ADD,
		Operands{A, B},
	},
	0x81: {
		ADD,
		Operands{A, C},
	},
	0x82: {
		ADD,
		Operands{A, D},
	},
	0x83: {
		ADD,
		Operands{A, E},
	},
	0x84: {
		ADD,
		Operands{A, H},
	},
	0x85: {
		ADD,
		Operands{A, L},
	},
	0x86: {
		ADD,
		Operands{A, Deref(HL)},
	},
	0x87: {
		ADD,
		Operands{A, A},
	},
	0x88: {
		ADC,
		Operands{A, B},
	},
	0x89: {
		ADC,
		Operands{A, C},
	},
	0x8A: {
		ADC,
		Operands{A, D},
	},
	0x8B: {
		ADC,
		Operands{A, E},
	},
	0x8C: {
		ADC,
		Operands{A, H},
	},
	0x8D: {
		ADC,
		Operands{A, L},
	},
	0x8E: {
		ADC,
		Operands{A, Deref(HL)},
	},
	0x8F: {
		ADC,
		Operands{A, A},
	},
	0x90: {
		SUB,
		Operands{B},
	},
	0x91: {
		SUB,
		Operands{C},
	},
	0x92: {
		SUB,
		Operands{D},
	},
	0x93: {
		SUB,
		Operands{E},
	},
	0x94: {
		SUB,
		Operands{H},
	},
	0x95: {
		SUB,
		Operands{L},
	},
	0x96: {
		SUB,
		Operands{Deref(HL)},
	},
	0x97: {
		SUB,
		Operands{A},
	},
	0x98: {
		SBC,
		Operands{A, B},
	},
	0x99: {
		SBC,
		Operands{A, C},
	},
	0x9A: {
		SBC,
		Operands{A, D},
	},
	0x9B: {
		SBC,
		Operands{A, E},
	},
	0x9C: {
		SBC,
		Operands{A, H},
	},
	0x9D: {
		SBC,
		Operands{A, L},
	},
	0x9E: {
		SBC,
		Operands{A, Deref(HL)},
	},
	0x9F: {
		SBC,
		Operands{A, A},
	},
	0xA0: {
		AND,
		Operands{B},
	},
	0xA1: {
		AND,
		Operands{C},
	},
	0xA2: {
		AND,
		Operands{D},
	},
	0xA3: {
		AND,
		Operands{E},
	},
	0xA4: {
		AND,
		Operands{H},
	},
	0xA5: {
		AND,
		Operands{L},
	},
	0xA6: {
		AND,
		Operands{Deref(HL)},
	},
	0xA7: {
		AND,
		Operands{A},
	},
	0xA8: {
		XOR,
		Operands{B},
	},
	0xA9: {
		XOR,
		Operands{C},
	},
	0xAA: {
		XOR,
		Operands{D},
	},
	0xAB: {
		XOR,
		Operands{E},
	},
	0xAC: {
		XOR,
		Operands{H},
	},
	0xAD: {
		XOR,
		Operands{L},
	},
	0xAE: {
		XOR,
		Operands{Deref(HL)},
	},
	0xAF: {
		XOR,
		Operands{A},
	},
	0xB0: {
		OR,
		Operands{B},
	},
	0xB1: {
		OR,
		Operands{C},
	},
	0xB2: {
		OR,
		Operands{D},
	},
	0xB3: {
		OR,
		Operands{E},
	},
	0xB4: {
		OR,
		Operands{H},
	},
	0xB5: {
		OR,
		Operands{L},
	},
	0xB6: {
		OR,
		Operands{Deref(HL)},
	},
	0xB7: {
		OR,
		Operands{A},
	},
	0xB8: {
		CP,
		Operands{B},
	},
	0xB9: {
		CP,
		Operands{C},
	},
	0xBA: {
		CP,
		Operands{D},
	},
	0xBB: {
		CP,
		Operands{E},
	},
	0xBC: {
		CP,
		Operands{H},
	},
	0xBD: {
		CP,
		Operands{L},
	},
	0xBE: {
		CP,
		Operands{Deref(HL)},
	},
	0xBF: {
		CP,
		Operands{A},
	},
	0xC0: {
		RET,
		Operands{NZ},
	},
	0xC1: {
		POP,
		Operands{BC},
	},
	0xC2: {
		JP,
		Operands{NZ, A16},
	},
	0xC3: {
		JP,
		Operands{A16},
	},
	0xC4: {
		CALL,
		Operands{NZ, A16},
	},
	0xC5: {
		PUSH,
		Operands{BC},
	},
	0xC6: {
		ADD,
		Operands{A, D8},
	},
	0xC7: {
		RST,
		Operands{Hex(0x00)},
	},
	0xC8: {
		RET,
		Operands{Z},
	},
	0xC9: {
		RET,
		nil,
	},
	0xCA: {
		JP,
		Operands{Z, A16},
	},
	0xCB: {
		PREFIX,
		nil,
	},
	0xCC: {
		CALL,
		Operands{Z, A16},
	},
	0xCD: {
		CALL,
		Operands{A16},
	},
	0xCE: {
		ADC,
		Operands{A, D8},
	},
	0xCF: {
		RST,
		Operands{Hex(0x08)},
	},
	0xD0: {
		RET,
		Operands{NC},
	},
	0xD1: {
		POP,
		Operands{DE},
	},
	0xD2: {
		JP,
		Operands{NC, A16},
	},
	0xD3: {
		ILLEGAL_D3,
		nil,
	},
	0xD4: {
		CALL,
		Operands{NC, A16},
	},
	0xD5: {
		PUSH,
		Operands{DE},
	},
	0xD6: {
		SUB,
		Operands{D8},
	},
	0xD7: {
		RST,
		Operands{Hex(0x10)},
	},
	0xD8: {
		RET,
		Operands{C},
	},
	0xD9: {
		RETI,
		nil,
	},
	0xDA: {
		JP,
		Operands{Ca, A16},
	},
	0xDB: {
		ILLEGAL_DB,
		nil,
	},
	0xDC: {
		CALL,
		Operands{C, A16},
	},
	0xDD: {
		ILLEGAL_DD,
		nil,
	},
	0xDE: {
		SBC,
		Operands{A, D8},
	},
	0xDF: {
		RST,
		Operands{Hex(0x18)},
	},
	0xE0: {
		LDH,
		Operands{Deref(A8), A},
	},
	0xE1: {
		POP,
		Operands{HL},
	},
	0xE2: {
		LD,
		Operands{Deref(C), A},
	},
	0xE3: {
		ILLEGAL_E3,
		nil,
	},
	0xE4: {
		ILLEGAL_E4,
		nil,
	},
	0xE5: {
		PUSH,
		Operands{HL},
	},
	0xE6: {
		AND,
		Operands{D8},
	},
	0xE7: {
		RST,
		Operands{Hex(0x20)},
	},
	0xE8: {
		ADD,
		Operands{SP, R8},
	},
	0xE9: {
		JP,
		Operands{HL},
	},
	0xEA: {
		LD,
		Operands{Deref(A16), A},
	},
	0xEB: {
		ILLEGAL_EB,
		nil,
	},
	0xEC: {
		ILLEGAL_EC,
		nil,
	},
	0xED: {
		ILLEGAL_ED,
		nil,
	},
	0xEE: {
		XOR,
		Operands{D8},
	},
	0xEF: {
		RST,
		Operands{Hex(0x28)},
	},
	0xF0: {
		LDH,
		Operands{A, Deref(A8)},
	},
	0xF1: {
		POP,
		Operands{AF},
	},
	0xF2: {
		LD,
		Operands{A, Deref(C)},
	},
	0xF3: {
		DI,
		nil,
	},
	0xF4: {
		ILLEGAL_F4,
		nil,
	},
	0xF5: {
		PUSH,
		Operands{AF},
	},
	0xF6: {
		OR,
		Operands{D8},
	},
	0xF7: {
		RST,
		Operands{Hex(0x30)},
	},
	0xF8: {
		LD,
		Operands{HL, Inc(SP), R8},
	},
	0xF9: {
		LD,
		Operands{SP, HL},
	},
	0xFA: {
		LD,
		Operands{A, Deref(A16)},
	},
	0xFB: {
		EI,
		nil,
	},
	0xFC: {
		ILLEGAL_FC,
		nil,
	},
	0xFD: {
		ILLEGAL_FD,
		nil,
	},
	0xFE: {
		CP,
		Operands{D8},
	},
	0xFF: {
		RST,
		Operands{Hex(0x38)},
	},
}

var cbprefixed = map[byte]Instruction{
	0x00: {
		RLC,
		Operands{B},
	},
	0x01: {
		RLC,
		Operands{C},
	},
	0x02: {
		RLC,
		Operands{D},
	},
	0x03: {
		RLC,
		Operands{E},
	},
	0x04: {
		RLC,
		Operands{H},
	},
	0x05: {
		RLC,
		Operands{L},
	},
	0x06: {
		RLC,
		Operands{Deref(HL)},
	},
	0x07: {
		RLC,
		Operands{A},
	},
	0x08: {
		RRC,
		Operands{B},
	},
	0x09: {
		RRC,
		Operands{C},
	},
	0x0A: {
		RRC,
		Operands{D},
	},
	0x0B: {
		RRC,
		Operands{E},
	},
	0x0C: {
		RRC,
		Operands{H},
	},
	0x0D: {
		RRC,
		Operands{L},
	},
	0x0E: {
		RRC,
		Operands{Deref(HL)},
	},
	0x0F: {
		RRC,
		Operands{A},
	},
	0x10: {
		RL,
		Operands{B},
	},
	0x11: {
		RL,
		Operands{C},
	},
	0x12: {
		RL,
		Operands{D},
	},
	0x13: {
		RL,
		Operands{E},
	},
	0x14: {
		RL,
		Operands{H},
	},
	0x15: {
		RL,
		Operands{L},
	},
	0x16: {
		RL,
		Operands{Deref(HL)},
	},
	0x17: {
		RL,
		Operands{A},
	},
	0x18: {
		RR,
		Operands{B},
	},
	0x19: {
		RR,
		Operands{C},
	},
	0x1A: {
		RR,
		Operands{D},
	},
	0x1B: {
		RR,
		Operands{E},
	},
	0x1C: {
		RR,
		Operands{H},
	},
	0x1D: {
		RR,
		Operands{L},
	},
	0x1E: {
		RR,
		Operands{Deref(HL)},
	},
	0x1F: {
		RR,
		Operands{A},
	},
	0x20: {
		SLA,
		Operands{B},
	},
	0x21: {
		SLA,
		Operands{C},
	},
	0x22: {
		SLA,
		Operands{D},
	},
	0x23: {
		SLA,
		Operands{E},
	},
	0x24: {
		SLA,
		Operands{H},
	},
	0x25: {
		SLA,
		Operands{L},
	},
	0x26: {
		SLA,
		Operands{Deref(HL)},
	},
	0x27: {
		SLA,
		Operands{A},
	},
	0x28: {
		SRA,
		Operands{B},
	},
	0x29: {
		SRA,
		Operands{C},
	},
	0x2A: {
		SRA,
		Operands{D},
	},
	0x2B: {
		SRA,
		Operands{E},
	},
	0x2C: {
		SRA,
		Operands{H},
	},
	0x2D: {
		SRA,
		Operands{L},
	},
	0x2E: {
		SRA,
		Operands{Deref(HL)},
	},
	0x2F: {
		SRA,
		Operands{A},
	},
	0x30: {
		SWAP,
		Operands{B},
	},
	0x31: {
		SWAP,
		Operands{C},
	},
	0x32: {
		SWAP,
		Operands{D},
	},
	0x33: {
		SWAP,
		Operands{E},
	},
	0x34: {
		SWAP,
		Operands{H},
	},
	0x35: {
		SWAP,
		Operands{L},
	},
	0x36: {
		SWAP,
		Operands{Deref(HL)},
	},
	0x37: {
		SWAP,
		Operands{A},
	},
	0x38: {
		SRL,
		Operands{B},
	},
	0x39: {
		SRL,
		Operands{C},
	},
	0x3A: {
		SRL,
		Operands{D},
	},
	0x3B: {
		SRL,
		Operands{E},
	},
	0x3C: {
		SRL,
		Operands{H},
	},
	0x3D: {
		SRL,
		Operands{L},
	},
	0x3E: {
		SRL,
		Operands{Deref(HL)},
	},
	0x3F: {
		SRL,
		Operands{A},
	},
	0x40: {
		BIT,
		Operands{Bit(0), B},
	},
	0x41: {
		BIT,
		Operands{Bit(0), C},
	},
	0x42: {
		BIT,
		Operands{Bit(0), D},
	},
	0x43: {
		BIT,
		Operands{Bit(0), E},
	},
	0x44: {
		BIT,
		Operands{Bit(0), H},
	},
	0x45: {
		BIT,
		Operands{Bit(0), L},
	},
	0x46: {
		BIT,
		Operands{Bit(0), Deref(HL)},
	},
	0x47: {
		BIT,
		Operands{Bit(0), A},
	},
	0x48: {
		BIT,
		Operands{Bit(1), B},
	},
	0x49: {
		BIT,
		Operands{Bit(1), C},
	},
	0x4A: {
		BIT,
		Operands{Bit(1), D},
	},
	0x4B: {
		BIT,
		Operands{Bit(1), E},
	},
	0x4C: {
		BIT,
		Operands{Bit(1), H},
	},
	0x4D: {
		BIT,
		Operands{Bit(1), L},
	},
	0x4E: {
		BIT,
		Operands{Bit(1), Deref(HL)},
	},
	0x4F: {
		BIT,
		Operands{Bit(1), A},
	},
	0x50: {
		BIT,
		Operands{Bit(2), B},
	},
	0x51: {
		BIT,
		Operands{Bit(2), C},
	},
	0x52: {
		BIT,
		Operands{Bit(2), D},
	},
	0x53: {
		BIT,
		Operands{Bit(2), E},
	},
	0x54: {
		BIT,
		Operands{Bit(2), H},
	},
	0x55: {
		BIT,
		Operands{Bit(2), L},
	},
	0x56: {
		BIT,
		Operands{Bit(2), Deref(HL)},
	},
	0x57: {
		BIT,
		Operands{Bit(2), A},
	},
	0x58: {
		BIT,
		Operands{Bit(3), B},
	},
	0x59: {
		BIT,
		Operands{Bit(3), C},
	},
	0x5A: {
		BIT,
		Operands{Bit(3), D},
	},
	0x5B: {
		BIT,
		Operands{Bit(3), E},
	},
	0x5C: {
		BIT,
		Operands{Bit(3), H},
	},
	0x5D: {
		BIT,
		Operands{Bit(3), L},
	},
	0x5E: {
		BIT,
		Operands{Bit(3), Deref(HL)},
	},
	0x5F: {
		BIT,
		Operands{Bit(3), A},
	},
	0x60: {
		BIT,
		Operands{Bit(4), B},
	},
	0x61: {
		BIT,
		Operands{Bit(4), C},
	},
	0x62: {
		BIT,
		Operands{Bit(4), D},
	},
	0x63: {
		BIT,
		Operands{Bit(4), E},
	},
	0x64: {
		BIT,
		Operands{Bit(4), H},
	},
	0x65: {
		BIT,
		Operands{Bit(4), L},
	},
	0x66: {
		BIT,
		Operands{Bit(4), Deref(HL)},
	},
	0x67: {
		BIT,
		Operands{Bit(4), A},
	},
	0x68: {
		BIT,
		Operands{Bit(5), B},
	},
	0x69: {
		BIT,
		Operands{Bit(5), C},
	},
	0x6A: {
		BIT,
		Operands{Bit(5), D},
	},
	0x6B: {
		BIT,
		Operands{Bit(5), E},
	},
	0x6C: {
		BIT,
		Operands{Bit(5), H},
	},
	0x6D: {
		BIT,
		Operands{Bit(5), L},
	},
	0x6E: {
		BIT,
		Operands{Bit(5), Deref(HL)},
	},
	0x6F: {
		BIT,
		Operands{Bit(5), A},
	},
	0x70: {
		BIT,
		Operands{Bit(6), B},
	},
	0x71: {
		BIT,
		Operands{Bit(6), C},
	},
	0x72: {
		BIT,
		Operands{Bit(6), D},
	},
	0x73: {
		BIT,
		Operands{Bit(6), E},
	},
	0x74: {
		BIT,
		Operands{Bit(6), H},
	},
	0x75: {
		BIT,
		Operands{Bit(6), L},
	},
	0x76: {
		BIT,
		Operands{Bit(6), Deref(HL)},
	},
	0x77: {
		BIT,
		Operands{Bit(6), A},
	},
	0x78: {
		BIT,
		Operands{Bit(7), B},
	},
	0x79: {
		BIT,
		Operands{Bit(7), C},
	},
	0x7A: {
		BIT,
		Operands{Bit(7), D},
	},
	0x7B: {
		BIT,
		Operands{Bit(7), E},
	},
	0x7C: {
		BIT,
		Operands{Bit(7), H},
	},
	0x7D: {
		BIT,
		Operands{Bit(7), L},
	},
	0x7E: {
		BIT,
		Operands{Bit(7), Deref(HL)},
	},
	0x7F: {
		BIT,
		Operands{Bit(7), A},
	},
	0x80: {
		RES,
		Operands{Bit(0), B},
	},
	0x81: {
		RES,
		Operands{Bit(0), C},
	},
	0x82: {
		RES,
		Operands{Bit(0), D},
	},
	0x83: {
		RES,
		Operands{Bit(0), E},
	},
	0x84: {
		RES,
		Operands{Bit(0), H},
	},
	0x85: {
		RES,
		Operands{Bit(0), L},
	},
	0x86: {
		RES,
		Operands{Bit(0), Deref(HL)},
	},
	0x87: {
		RES,
		Operands{Bit(0), A},
	},
	0x88: {
		RES,
		Operands{Bit(1), B},
	},
	0x89: {
		RES,
		Operands{Bit(1), C},
	},
	0x8A: {
		RES,
		Operands{Bit(1), D},
	},
	0x8B: {
		RES,
		Operands{Bit(1), E},
	},
	0x8C: {
		RES,
		Operands{Bit(1), H},
	},
	0x8D: {
		RES,
		Operands{Bit(1), L},
	},
	0x8E: {
		RES,
		Operands{Bit(1), Deref(HL)},
	},
	0x8F: {
		RES,
		Operands{Bit(1), A},
	},
	0x90: {
		RES,
		Operands{Bit(2), B},
	},
	0x91: {
		RES,
		Operands{Bit(2), C},
	},
	0x92: {
		RES,
		Operands{Bit(2), D},
	},
	0x93: {
		RES,
		Operands{Bit(2), E},
	},
	0x94: {
		RES,
		Operands{Bit(2), H},
	},
	0x95: {
		RES,
		Operands{Bit(2), L},
	},
	0x96: {
		RES,
		Operands{Bit(2), Deref(HL)},
	},
	0x97: {
		RES,
		Operands{Bit(2), A},
	},
	0x98: {
		RES,
		Operands{Bit(3), B},
	},
	0x99: {
		RES,
		Operands{Bit(3), C},
	},
	0x9A: {
		RES,
		Operands{Bit(3), D},
	},
	0x9B: {
		RES,
		Operands{Bit(3), E},
	},
	0x9C: {
		RES,
		Operands{Bit(3), H},
	},
	0x9D: {
		RES,
		Operands{Bit(3), L},
	},
	0x9E: {
		RES,
		Operands{Bit(3), Deref(HL)},
	},
	0x9F: {
		RES,
		Operands{Bit(3), A},
	},
	0xA0: {
		RES,
		Operands{Bit(4), B},
	},
	0xA1: {
		RES,
		Operands{Bit(4), C},
	},
	0xA2: {
		RES,
		Operands{Bit(4), D},
	},
	0xA3: {
		RES,
		Operands{Bit(4), E},
	},
	0xA4: {
		RES,
		Operands{Bit(4), H},
	},
	0xA5: {
		RES,
		Operands{Bit(4), L},
	},
	0xA6: {
		RES,
		Operands{Bit(4), Deref(HL)},
	},
	0xA7: {
		RES,
		Operands{Bit(4), A},
	},
	0xA8: {
		RES,
		Operands{Bit(5), B},
	},
	0xA9: {
		RES,
		Operands{Bit(5), C},
	},
	0xAA: {
		RES,
		Operands{Bit(5), D},
	},
	0xAB: {
		RES,
		Operands{Bit(5), E},
	},
	0xAC: {
		RES,
		Operands{Bit(5), H},
	},
	0xAD: {
		RES,
		Operands{Bit(5), L},
	},
	0xAE: {
		RES,
		Operands{Bit(5), Deref(HL)},
	},
	0xAF: {
		RES,
		Operands{Bit(5), A},
	},
	0xB0: {
		RES,
		Operands{Bit(6), B},
	},
	0xB1: {
		RES,
		Operands{Bit(6), C},
	},
	0xB2: {
		RES,
		Operands{Bit(6), D},
	},
	0xB3: {
		RES,
		Operands{Bit(6), E},
	},
	0xB4: {
		RES,
		Operands{Bit(6), H},
	},
	0xB5: {
		RES,
		Operands{Bit(6), L},
	},
	0xB6: {
		RES,
		Operands{Bit(6), Deref(HL)},
	},
	0xB7: {
		RES,
		Operands{Bit(6), A},
	},
	0xB8: {
		RES,
		Operands{Bit(7), B},
	},
	0xB9: {
		RES,
		Operands{Bit(7), C},
	},
	0xBA: {
		RES,
		Operands{Bit(7), D},
	},
	0xBB: {
		RES,
		Operands{Bit(7), E},
	},
	0xBC: {
		RES,
		Operands{Bit(7), H},
	},
	0xBD: {
		RES,
		Operands{Bit(7), L},
	},
	0xBE: {
		RES,
		Operands{Bit(7), Deref(HL)},
	},
	0xBF: {
		RES,
		Operands{Bit(7), A},
	},
	0xC0: {
		SET,
		Operands{Bit(0), B},
	},
	0xC1: {
		SET,
		Operands{Bit(0), C},
	},
	0xC2: {
		SET,
		Operands{Bit(0), D},
	},
	0xC3: {
		SET,
		Operands{Bit(0), E},
	},
	0xC4: {
		SET,
		Operands{Bit(0), H},
	},
	0xC5: {
		SET,
		Operands{Bit(0), L},
	},
	0xC6: {
		SET,
		Operands{Bit(0), Deref(HL)},
	},
	0xC7: {
		SET,
		Operands{Bit(0), A},
	},
	0xC8: {
		SET,
		Operands{Bit(1), B},
	},
	0xC9: {
		SET,
		Operands{Bit(1), C},
	},
	0xCA: {
		SET,
		Operands{Bit(1), D},
	},
	0xCB: {
		SET,
		Operands{Bit(1), E},
	},
	0xCC: {
		SET,
		Operands{Bit(1), H},
	},
	0xCD: {
		SET,
		Operands{Bit(1), L},
	},
	0xCE: {
		SET,
		Operands{Bit(1), Deref(HL)},
	},
	0xCF: {
		SET,
		Operands{Bit(1), A},
	},
	0xD0: {
		SET,
		Operands{Bit(2), B},
	},
	0xD1: {
		SET,
		Operands{Bit(2), C},
	},
	0xD2: {
		SET,
		Operands{Bit(2), D},
	},
	0xD3: {
		SET,
		Operands{Bit(2), E},
	},
	0xD4: {
		SET,
		Operands{Bit(2), H},
	},
	0xD5: {
		SET,
		Operands{Bit(2), L},
	},
	0xD6: {
		SET,
		Operands{Bit(2), Deref(HL)},
	},
	0xD7: {
		SET,
		Operands{Bit(2), A},
	},
	0xD8: {
		SET,
		Operands{Bit(3), B},
	},
	0xD9: {
		SET,
		Operands{Bit(3), C},
	},
	0xDA: {
		SET,
		Operands{Bit(3), D},
	},
	0xDB: {
		SET,
		Operands{Bit(3), E},
	},
	0xDC: {
		SET,
		Operands{Bit(3), H},
	},
	0xDD: {
		SET,
		Operands{Bit(3), L},
	},
	0xDE: {
		SET,
		Operands{Bit(3), Deref(HL)},
	},
	0xDF: {
		SET,
		Operands{Bit(3), A},
	},
	0xE0: {
		SET,
		Operands{Bit(4), B},
	},
	0xE1: {
		SET,
		Operands{Bit(4), C},
	},
	0xE2: {
		SET,
		Operands{Bit(4), D},
	},
	0xE3: {
		SET,
		Operands{Bit(4), E},
	},
	0xE4: {
		SET,
		Operands{Bit(4), H},
	},
	0xE5: {
		SET,
		Operands{Bit(4), L},
	},
	0xE6: {
		SET,
		Operands{Bit(4), Deref(HL)},
	},
	0xE7: {
		SET,
		Operands{Bit(4), A},
	},
	0xE8: {
		SET,
		Operands{Bit(5), B},
	},
	0xE9: {
		SET,
		Operands{Bit(5), C},
	},
	0xEA: {
		SET,
		Operands{Bit(5), D},
	},
	0xEB: {
		SET,
		Operands{Bit(5), E},
	},
	0xEC: {
		SET,
		Operands{Bit(5), H},
	},
	0xED: {
		SET,
		Operands{Bit(5), L},
	},
	0xEE: {
		SET,
		Operands{Bit(5), Deref(HL)},
	},
	0xEF: {
		SET,
		Operands{Bit(5), A},
	},
	0xF0: {
		SET,
		Operands{Bit(6), B},
	},
	0xF1: {
		SET,
		Operands{Bit(6), C},
	},
	0xF2: {
		SET,
		Operands{Bit(6), D},
	},
	0xF3: {
		SET,
		Operands{Bit(6), E},
	},
	0xF4: {
		SET,
		Operands{Bit(6), H},
	},
	0xF5: {
		SET,
		Operands{Bit(6), L},
	},
	0xF6: {
		SET,
		Operands{Bit(6), Deref(HL)},
	},
	0xF7: {
		SET,
		Operands{Bit(6), A},
	},
	0xF8: {
		SET,
		Operands{Bit(7), B},
	},
	0xF9: {
		SET,
		Operands{Bit(7), C},
	},
	0xFA: {
		SET,
		Operands{Bit(7), D},
	},
	0xFB: {
		SET,
		Operands{Bit(7), E},
	},
	0xFC: {
		SET,
		Operands{Bit(7), H},
	},
	0xFD: {
		SET,
		Operands{Bit(7), L},
	},
	0xFE: {
		SET,
		Operands{Bit(7), Deref(HL)},
	},
	0xFF: {
		SET,
		Operands{Bit(7), A},
	},
}

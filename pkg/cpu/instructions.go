package cpu

// https://gbdev.io/pandocs/CPU_Instruction_Set.html
// https://gbdev.io/gb-opcodes/optables/
// instructions generated from: https://gbdev.io/gb-opcodes/Opcodes.json
// script: https://gist.github.com/robherley/836369cbd8eb73a286d017626b8376c1

type Instruction struct {
	Operation
	Operands []Operand
}

func InstructionFromOPCode(code Byte, cbprefix bool) *Instruction {
	var mapping map[Byte]Instruction
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

var unprefixed = map[Byte]Instruction{
	0x00: {
		NOP,
		nil,
	},
	0x01: {
		LD,
		[]Operand{
			{Symbol: BC},
			{Symbol: D16},
		},
	},
	0x02: {
		LD,
		[]Operand{
			{Symbol: BC, Deref: true},
			{Symbol: A},
		},
	},
	0x03: {
		INC,
		[]Operand{
			{Symbol: BC},
		},
	},
	0x04: {
		INC,
		[]Operand{
			{Symbol: B},
		},
	},
	0x05: {
		DEC,
		[]Operand{
			{Symbol: B},
		},
	},
	0x06: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: D8},
		},
	},
	0x07: {
		RLCA,
		nil,
	},
	0x08: {
		LD,
		[]Operand{
			{Symbol: A16, Deref: true},
			{Symbol: SP},
		},
	},
	0x09: {
		ADD,
		[]Operand{
			{Symbol: HL},
			{Symbol: BC},
		},
	},
	0x0A: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: BC, Deref: true},
		},
	},
	0x0B: {
		DEC,
		[]Operand{
			{Symbol: BC},
		},
	},
	0x0C: {
		INC,
		[]Operand{
			{Symbol: C},
		},
	},
	0x0D: {
		DEC,
		[]Operand{
			{Symbol: C},
		},
	},
	0x0E: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: D8},
		},
	},
	0x0F: {
		RRCA,
		nil,
	},
	0x10: {
		STOP,
		[]Operand{
			{Symbol: D8},
		},
	},
	0x11: {
		LD,
		[]Operand{
			{Symbol: DE},
			{Symbol: D16},
		},
	},
	0x12: {
		LD,
		[]Operand{
			{Symbol: DE, Deref: true},
			{Symbol: A},
		},
	},
	0x13: {
		INC,
		[]Operand{
			{Symbol: DE},
		},
	},
	0x14: {
		INC,
		[]Operand{
			{Symbol: D},
		},
	},
	0x15: {
		DEC,
		[]Operand{
			{Symbol: D},
		},
	},
	0x16: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: D8},
		},
	},
	0x17: {
		RLA,
		nil,
	},
	0x18: {
		JR,
		[]Operand{
			{Symbol: R8},
		},
	},
	0x19: {
		ADD,
		[]Operand{
			{Symbol: HL},
			{Symbol: DE},
		},
	},
	0x1A: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: DE, Deref: true},
		},
	},
	0x1B: {
		DEC,
		[]Operand{
			{Symbol: DE},
		},
	},
	0x1C: {
		INC,
		[]Operand{
			{Symbol: E},
		},
	},
	0x1D: {
		DEC,
		[]Operand{
			{Symbol: E},
		},
	},
	0x1E: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: D8},
		},
	},
	0x1F: {
		RRA,
		nil,
	},
	0x20: {
		JR,
		[]Operand{
			{Symbol: NZ},
			{Symbol: R8},
		},
	},
	0x21: {
		LD,
		[]Operand{
			{Symbol: HL},
			{Symbol: D16},
		},
	},
	0x22: {
		LD,
		[]Operand{
			{Symbol: HL, Inc: true, Deref: true},
			{Symbol: A},
		},
	},
	0x23: {
		INC,
		[]Operand{
			{Symbol: HL},
		},
	},
	0x24: {
		INC,
		[]Operand{
			{Symbol: H},
		},
	},
	0x25: {
		DEC,
		[]Operand{
			{Symbol: H},
		},
	},
	0x26: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: D8},
		},
	},
	0x27: {
		DAA,
		nil,
	},
	0x28: {
		JR,
		[]Operand{
			{Symbol: Z},
			{Symbol: R8},
		},
	},
	0x29: {
		ADD,
		[]Operand{
			{Symbol: HL},
			{Symbol: HL},
		},
	},
	0x2A: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Inc: true, Deref: true},
		},
	},
	0x2B: {
		DEC,
		[]Operand{
			{Symbol: HL},
		},
	},
	0x2C: {
		INC,
		[]Operand{
			{Symbol: L},
		},
	},
	0x2D: {
		DEC,
		[]Operand{
			{Symbol: L},
		},
	},
	0x2E: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: D8},
		},
	},
	0x2F: {
		CPL,
		nil,
	},
	0x30: {
		JR,
		[]Operand{
			{Symbol: NC},
			{Symbol: R8},
		},
	},
	0x31: {
		LD,
		[]Operand{
			{Symbol: SP},
			{Symbol: D16},
		},
	},
	0x32: {
		LD,
		[]Operand{
			{Symbol: HL, Dec: true, Deref: true},
			{Symbol: A},
		},
	},
	0x33: {
		INC,
		[]Operand{
			{Symbol: SP},
		},
	},
	0x34: {
		INC,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x35: {
		DEC,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x36: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: D8},
		},
	},
	0x37: {
		SCF,
		nil,
	},
	0x38: {
		JR,
		[]Operand{
			{Symbol: C},
			{Symbol: R8},
		},
	},
	0x39: {
		ADD,
		[]Operand{
			{Symbol: HL},
			{Symbol: SP},
		},
	},
	0x3A: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Dec: true, Deref: true},
		},
	},
	0x3B: {
		DEC,
		[]Operand{
			{Symbol: SP},
		},
	},
	0x3C: {
		INC,
		[]Operand{
			{Symbol: A},
		},
	},
	0x3D: {
		DEC,
		[]Operand{
			{Symbol: A},
		},
	},
	0x3E: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: D8},
		},
	},
	0x3F: {
		CCF,
		nil,
	},
	0x40: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: B},
		},
	},
	0x41: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: C},
		},
	},
	0x42: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: D},
		},
	},
	0x43: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: E},
		},
	},
	0x44: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: H},
		},
	},
	0x45: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: L},
		},
	},
	0x46: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: HL, Deref: true},
		},
	},
	0x47: {
		LD,
		[]Operand{
			{Symbol: B},
			{Symbol: A},
		},
	},
	0x48: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: B},
		},
	},
	0x49: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: C},
		},
	},
	0x4A: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: D},
		},
	},
	0x4B: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: E},
		},
	},
	0x4C: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: H},
		},
	},
	0x4D: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: L},
		},
	},
	0x4E: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: HL, Deref: true},
		},
	},
	0x4F: {
		LD,
		[]Operand{
			{Symbol: C},
			{Symbol: A},
		},
	},
	0x50: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: B},
		},
	},
	0x51: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: C},
		},
	},
	0x52: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: D},
		},
	},
	0x53: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: E},
		},
	},
	0x54: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: H},
		},
	},
	0x55: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: L},
		},
	},
	0x56: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: HL, Deref: true},
		},
	},
	0x57: {
		LD,
		[]Operand{
			{Symbol: D},
			{Symbol: A},
		},
	},
	0x58: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: B},
		},
	},
	0x59: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: C},
		},
	},
	0x5A: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: D},
		},
	},
	0x5B: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: E},
		},
	},
	0x5C: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: H},
		},
	},
	0x5D: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: L},
		},
	},
	0x5E: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: HL, Deref: true},
		},
	},
	0x5F: {
		LD,
		[]Operand{
			{Symbol: E},
			{Symbol: A},
		},
	},
	0x60: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: B},
		},
	},
	0x61: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: C},
		},
	},
	0x62: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: D},
		},
	},
	0x63: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: E},
		},
	},
	0x64: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: H},
		},
	},
	0x65: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: L},
		},
	},
	0x66: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: HL, Deref: true},
		},
	},
	0x67: {
		LD,
		[]Operand{
			{Symbol: H},
			{Symbol: A},
		},
	},
	0x68: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: B},
		},
	},
	0x69: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: C},
		},
	},
	0x6A: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: D},
		},
	},
	0x6B: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: E},
		},
	},
	0x6C: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: H},
		},
	},
	0x6D: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: L},
		},
	},
	0x6E: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: HL, Deref: true},
		},
	},
	0x6F: {
		LD,
		[]Operand{
			{Symbol: L},
			{Symbol: A},
		},
	},
	0x70: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: B},
		},
	},
	0x71: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: C},
		},
	},
	0x72: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: D},
		},
	},
	0x73: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: E},
		},
	},
	0x74: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: H},
		},
	},
	0x75: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: L},
		},
	},
	0x76: {
		HALT,
		nil,
	},
	0x77: {
		LD,
		[]Operand{
			{Symbol: HL, Deref: true},
			{Symbol: A},
		},
	},
	0x78: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: B},
		},
	},
	0x79: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: C},
		},
	},
	0x7A: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: D},
		},
	},
	0x7B: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: E},
		},
	},
	0x7C: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: H},
		},
	},
	0x7D: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: L},
		},
	},
	0x7E: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Deref: true},
		},
	},
	0x7F: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: A},
		},
	},
	0x80: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: B},
		},
	},
	0x81: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: C},
		},
	},
	0x82: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: D},
		},
	},
	0x83: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: E},
		},
	},
	0x84: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: H},
		},
	},
	0x85: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: L},
		},
	},
	0x86: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Deref: true},
		},
	},
	0x87: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: A},
		},
	},
	0x88: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: B},
		},
	},
	0x89: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: C},
		},
	},
	0x8A: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: D},
		},
	},
	0x8B: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: E},
		},
	},
	0x8C: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: H},
		},
	},
	0x8D: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: L},
		},
	},
	0x8E: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Deref: true},
		},
	},
	0x8F: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: A},
		},
	},
	0x90: {
		SUB,
		[]Operand{
			{Symbol: B},
		},
	},
	0x91: {
		SUB,
		[]Operand{
			{Symbol: C},
		},
	},
	0x92: {
		SUB,
		[]Operand{
			{Symbol: D},
		},
	},
	0x93: {
		SUB,
		[]Operand{
			{Symbol: E},
		},
	},
	0x94: {
		SUB,
		[]Operand{
			{Symbol: H},
		},
	},
	0x95: {
		SUB,
		[]Operand{
			{Symbol: L},
		},
	},
	0x96: {
		SUB,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x97: {
		SUB,
		[]Operand{
			{Symbol: A},
		},
	},
	0x98: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: B},
		},
	},
	0x99: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: C},
		},
	},
	0x9A: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: D},
		},
	},
	0x9B: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: E},
		},
	},
	0x9C: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: H},
		},
	},
	0x9D: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: L},
		},
	},
	0x9E: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: HL, Deref: true},
		},
	},
	0x9F: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: A},
		},
	},
	0xA0: {
		AND,
		[]Operand{
			{Symbol: B},
		},
	},
	0xA1: {
		AND,
		[]Operand{
			{Symbol: C},
		},
	},
	0xA2: {
		AND,
		[]Operand{
			{Symbol: D},
		},
	},
	0xA3: {
		AND,
		[]Operand{
			{Symbol: E},
		},
	},
	0xA4: {
		AND,
		[]Operand{
			{Symbol: H},
		},
	},
	0xA5: {
		AND,
		[]Operand{
			{Symbol: L},
		},
	},
	0xA6: {
		AND,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0xA7: {
		AND,
		[]Operand{
			{Symbol: A},
		},
	},
	0xA8: {
		XOR,
		[]Operand{
			{Symbol: B},
		},
	},
	0xA9: {
		XOR,
		[]Operand{
			{Symbol: C},
		},
	},
	0xAA: {
		XOR,
		[]Operand{
			{Symbol: D},
		},
	},
	0xAB: {
		XOR,
		[]Operand{
			{Symbol: E},
		},
	},
	0xAC: {
		XOR,
		[]Operand{
			{Symbol: H},
		},
	},
	0xAD: {
		XOR,
		[]Operand{
			{Symbol: L},
		},
	},
	0xAE: {
		XOR,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0xAF: {
		XOR,
		[]Operand{
			{Symbol: A},
		},
	},
	0xB0: {
		OR,
		[]Operand{
			{Symbol: B},
		},
	},
	0xB1: {
		OR,
		[]Operand{
			{Symbol: C},
		},
	},
	0xB2: {
		OR,
		[]Operand{
			{Symbol: D},
		},
	},
	0xB3: {
		OR,
		[]Operand{
			{Symbol: E},
		},
	},
	0xB4: {
		OR,
		[]Operand{
			{Symbol: H},
		},
	},
	0xB5: {
		OR,
		[]Operand{
			{Symbol: L},
		},
	},
	0xB6: {
		OR,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0xB7: {
		OR,
		[]Operand{
			{Symbol: A},
		},
	},
	0xB8: {
		CP,
		[]Operand{
			{Symbol: B},
		},
	},
	0xB9: {
		CP,
		[]Operand{
			{Symbol: C},
		},
	},
	0xBA: {
		CP,
		[]Operand{
			{Symbol: D},
		},
	},
	0xBB: {
		CP,
		[]Operand{
			{Symbol: E},
		},
	},
	0xBC: {
		CP,
		[]Operand{
			{Symbol: H},
		},
	},
	0xBD: {
		CP,
		[]Operand{
			{Symbol: L},
		},
	},
	0xBE: {
		CP,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0xBF: {
		CP,
		[]Operand{
			{Symbol: A},
		},
	},
	0xC0: {
		RET,
		[]Operand{
			{Symbol: NZ},
		},
	},
	0xC1: {
		POP,
		[]Operand{
			{Symbol: BC},
		},
	},
	0xC2: {
		JP,
		[]Operand{
			{Symbol: NZ},
			{Symbol: A16},
		},
	},
	0xC3: {
		JP,
		[]Operand{
			{Symbol: A16},
		},
	},
	0xC4: {
		CALL,
		[]Operand{
			{Symbol: NZ},
			{Symbol: A16},
		},
	},
	0xC5: {
		PUSH,
		[]Operand{
			{Symbol: BC},
		},
	},
	0xC6: {
		ADD,
		[]Operand{
			{Symbol: A},
			{Symbol: D8},
		},
	},
	0xC7: {
		RST,
		[]Operand{
			{Symbol: Byte(0x00)},
		},
	},
	0xC8: {
		RET,
		[]Operand{
			{Symbol: Z},
		},
	},
	0xC9: {
		RET,
		nil,
	},
	0xCA: {
		JP,
		[]Operand{
			{Symbol: Z},
			{Symbol: A16},
		},
	},
	0xCB: {
		PREFIX,
		nil,
	},
	0xCC: {
		CALL,
		[]Operand{
			{Symbol: Z},
			{Symbol: A16},
		},
	},
	0xCD: {
		CALL,
		[]Operand{
			{Symbol: A16},
		},
	},
	0xCE: {
		ADC,
		[]Operand{
			{Symbol: A},
			{Symbol: D8},
		},
	},
	0xCF: {
		RST,
		[]Operand{
			{Symbol: Byte(0x08)},
		},
	},
	0xD0: {
		RET,
		[]Operand{
			{Symbol: NC},
		},
	},
	0xD1: {
		POP,
		[]Operand{
			{Symbol: DE},
		},
	},
	0xD2: {
		JP,
		[]Operand{
			{Symbol: NC},
			{Symbol: A16},
		},
	},
	0xD3: {
		ILLEGAL_D3,
		nil,
	},
	0xD4: {
		CALL,
		[]Operand{
			{Symbol: NC},
			{Symbol: A16},
		},
	},
	0xD5: {
		PUSH,
		[]Operand{
			{Symbol: DE},
		},
	},
	0xD6: {
		SUB,
		[]Operand{
			{Symbol: D8},
		},
	},
	0xD7: {
		RST,
		[]Operand{
			{Symbol: Byte(0x10)},
		},
	},
	0xD8: {
		RET,
		[]Operand{
			{Symbol: C},
		},
	},
	0xD9: {
		RETI,
		nil,
	},
	0xDA: {
		JP,
		[]Operand{
			{Symbol: Ca},
			{Symbol: A16},
		},
	},
	0xDB: {
		ILLEGAL_DB,
		nil,
	},
	0xDC: {
		CALL,
		[]Operand{
			{Symbol: C},
			{Symbol: A16},
		},
	},
	0xDD: {
		ILLEGAL_DD,
		nil,
	},
	0xDE: {
		SBC,
		[]Operand{
			{Symbol: A},
			{Symbol: D8},
		},
	},
	0xDF: {
		RST,
		[]Operand{
			{Symbol: Byte(0x18)},
		},
	},
	0xE0: {
		LDH,
		[]Operand{
			{Symbol: A8, Deref: true},
			{Symbol: A},
		},
	},
	0xE1: {
		POP,
		[]Operand{
			{Symbol: HL},
		},
	},
	0xE2: {
		LD,
		[]Operand{
			{Symbol: C, Deref: true},
			{Symbol: A},
		},
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
		[]Operand{
			{Symbol: HL},
		},
	},
	0xE6: {
		AND,
		[]Operand{
			{Symbol: D8},
		},
	},
	0xE7: {
		RST,
		[]Operand{
			{Symbol: Byte(0x20)},
		},
	},
	0xE8: {
		ADD,
		[]Operand{
			{Symbol: SP},
			{Symbol: R8},
		},
	},
	0xE9: {
		JP,
		[]Operand{
			{Symbol: HL},
		},
	},
	0xEA: {
		LD,
		[]Operand{
			{Symbol: A16, Deref: true},
			{Symbol: A},
		},
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
		[]Operand{
			{Symbol: D8},
		},
	},
	0xEF: {
		RST,
		[]Operand{
			{Symbol: Byte(0x28)},
		},
	},
	0xF0: {
		LDH,
		[]Operand{
			{Symbol: A},
			{Symbol: A8, Deref: true},
		},
	},
	0xF1: {
		POP,
		[]Operand{
			{Symbol: AF},
		},
	},
	0xF2: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: C, Deref: true},
		},
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
		[]Operand{
			{Symbol: AF},
		},
	},
	0xF6: {
		OR,
		[]Operand{
			{Symbol: D8},
		},
	},
	0xF7: {
		RST,
		[]Operand{
			{Symbol: Byte(0x30)},
		},
	},
	0xF8: {
		LD,
		[]Operand{
			{Symbol: HL},
			{Symbol: SP, Inc: true},
			{Symbol: R8},
		},
	},
	0xF9: {
		LD,
		[]Operand{
			{Symbol: SP},
			{Symbol: HL},
		},
	},
	0xFA: {
		LD,
		[]Operand{
			{Symbol: A},
			{Symbol: A16, Deref: true},
		},
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
		[]Operand{
			{Symbol: D8},
		},
	},
	0xFF: {
		RST,
		[]Operand{
			{Symbol: Byte(0x38)},
		},
	},
}

var cbprefixed = map[Byte]Instruction{
	0x00: {
		RLC,
		[]Operand{
			{Symbol: B},
		},
	},
	0x01: {
		RLC,
		[]Operand{
			{Symbol: C},
		},
	},
	0x02: {
		RLC,
		[]Operand{
			{Symbol: D},
		},
	},
	0x03: {
		RLC,
		[]Operand{
			{Symbol: E},
		},
	},
	0x04: {
		RLC,
		[]Operand{
			{Symbol: H},
		},
	},
	0x05: {
		RLC,
		[]Operand{
			{Symbol: L},
		},
	},
	0x06: {
		RLC,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x07: {
		RLC,
		[]Operand{
			{Symbol: A},
		},
	},
	0x08: {
		RRC,
		[]Operand{
			{Symbol: B},
		},
	},
	0x09: {
		RRC,
		[]Operand{
			{Symbol: C},
		},
	},
	0x0A: {
		RRC,
		[]Operand{
			{Symbol: D},
		},
	},
	0x0B: {
		RRC,
		[]Operand{
			{Symbol: E},
		},
	},
	0x0C: {
		RRC,
		[]Operand{
			{Symbol: H},
		},
	},
	0x0D: {
		RRC,
		[]Operand{
			{Symbol: L},
		},
	},
	0x0E: {
		RRC,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x0F: {
		RRC,
		[]Operand{
			{Symbol: A},
		},
	},
	0x10: {
		RL,
		[]Operand{
			{Symbol: B},
		},
	},
	0x11: {
		RL,
		[]Operand{
			{Symbol: C},
		},
	},
	0x12: {
		RL,
		[]Operand{
			{Symbol: D},
		},
	},
	0x13: {
		RL,
		[]Operand{
			{Symbol: E},
		},
	},
	0x14: {
		RL,
		[]Operand{
			{Symbol: H},
		},
	},
	0x15: {
		RL,
		[]Operand{
			{Symbol: L},
		},
	},
	0x16: {
		RL,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x17: {
		RL,
		[]Operand{
			{Symbol: A},
		},
	},
	0x18: {
		RR,
		[]Operand{
			{Symbol: B},
		},
	},
	0x19: {
		RR,
		[]Operand{
			{Symbol: C},
		},
	},
	0x1A: {
		RR,
		[]Operand{
			{Symbol: D},
		},
	},
	0x1B: {
		RR,
		[]Operand{
			{Symbol: E},
		},
	},
	0x1C: {
		RR,
		[]Operand{
			{Symbol: H},
		},
	},
	0x1D: {
		RR,
		[]Operand{
			{Symbol: L},
		},
	},
	0x1E: {
		RR,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x1F: {
		RR,
		[]Operand{
			{Symbol: A},
		},
	},
	0x20: {
		SLA,
		[]Operand{
			{Symbol: B},
		},
	},
	0x21: {
		SLA,
		[]Operand{
			{Symbol: C},
		},
	},
	0x22: {
		SLA,
		[]Operand{
			{Symbol: D},
		},
	},
	0x23: {
		SLA,
		[]Operand{
			{Symbol: E},
		},
	},
	0x24: {
		SLA,
		[]Operand{
			{Symbol: H},
		},
	},
	0x25: {
		SLA,
		[]Operand{
			{Symbol: L},
		},
	},
	0x26: {
		SLA,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x27: {
		SLA,
		[]Operand{
			{Symbol: A},
		},
	},
	0x28: {
		SRA,
		[]Operand{
			{Symbol: B},
		},
	},
	0x29: {
		SRA,
		[]Operand{
			{Symbol: C},
		},
	},
	0x2A: {
		SRA,
		[]Operand{
			{Symbol: D},
		},
	},
	0x2B: {
		SRA,
		[]Operand{
			{Symbol: E},
		},
	},
	0x2C: {
		SRA,
		[]Operand{
			{Symbol: H},
		},
	},
	0x2D: {
		SRA,
		[]Operand{
			{Symbol: L},
		},
	},
	0x2E: {
		SRA,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x2F: {
		SRA,
		[]Operand{
			{Symbol: A},
		},
	},
	0x30: {
		SWAP,
		[]Operand{
			{Symbol: B},
		},
	},
	0x31: {
		SWAP,
		[]Operand{
			{Symbol: C},
		},
	},
	0x32: {
		SWAP,
		[]Operand{
			{Symbol: D},
		},
	},
	0x33: {
		SWAP,
		[]Operand{
			{Symbol: E},
		},
	},
	0x34: {
		SWAP,
		[]Operand{
			{Symbol: H},
		},
	},
	0x35: {
		SWAP,
		[]Operand{
			{Symbol: L},
		},
	},
	0x36: {
		SWAP,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x37: {
		SWAP,
		[]Operand{
			{Symbol: A},
		},
	},
	0x38: {
		SRL,
		[]Operand{
			{Symbol: B},
		},
	},
	0x39: {
		SRL,
		[]Operand{
			{Symbol: C},
		},
	},
	0x3A: {
		SRL,
		[]Operand{
			{Symbol: D},
		},
	},
	0x3B: {
		SRL,
		[]Operand{
			{Symbol: E},
		},
	},
	0x3C: {
		SRL,
		[]Operand{
			{Symbol: H},
		},
	},
	0x3D: {
		SRL,
		[]Operand{
			{Symbol: L},
		},
	},
	0x3E: {
		SRL,
		[]Operand{
			{Symbol: HL, Deref: true},
		},
	},
	0x3F: {
		SRL,
		[]Operand{
			{Symbol: A},
		},
	},
	0x40: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: B},
		},
	},
	0x41: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: C},
		},
	},
	0x42: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: D},
		},
	},
	0x43: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: E},
		},
	},
	0x44: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: H},
		},
	},
	0x45: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: L},
		},
	},
	0x46: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: HL, Deref: true},
		},
	},
	0x47: {
		BIT,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: A},
		},
	},
	0x48: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: B},
		},
	},
	0x49: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: C},
		},
	},
	0x4A: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: D},
		},
	},
	0x4B: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: E},
		},
	},
	0x4C: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: H},
		},
	},
	0x4D: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: L},
		},
	},
	0x4E: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: HL, Deref: true},
		},
	},
	0x4F: {
		BIT,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: A},
		},
	},
	0x50: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: B},
		},
	},
	0x51: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: C},
		},
	},
	0x52: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: D},
		},
	},
	0x53: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: E},
		},
	},
	0x54: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: H},
		},
	},
	0x55: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: L},
		},
	},
	0x56: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: HL, Deref: true},
		},
	},
	0x57: {
		BIT,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: A},
		},
	},
	0x58: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: B},
		},
	},
	0x59: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: C},
		},
	},
	0x5A: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: D},
		},
	},
	0x5B: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: E},
		},
	},
	0x5C: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: H},
		},
	},
	0x5D: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: L},
		},
	},
	0x5E: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: HL, Deref: true},
		},
	},
	0x5F: {
		BIT,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: A},
		},
	},
	0x60: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: B},
		},
	},
	0x61: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: C},
		},
	},
	0x62: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: D},
		},
	},
	0x63: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: E},
		},
	},
	0x64: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: H},
		},
	},
	0x65: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: L},
		},
	},
	0x66: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: HL, Deref: true},
		},
	},
	0x67: {
		BIT,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: A},
		},
	},
	0x68: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: B},
		},
	},
	0x69: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: C},
		},
	},
	0x6A: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: D},
		},
	},
	0x6B: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: E},
		},
	},
	0x6C: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: H},
		},
	},
	0x6D: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: L},
		},
	},
	0x6E: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: HL, Deref: true},
		},
	},
	0x6F: {
		BIT,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: A},
		},
	},
	0x70: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: B},
		},
	},
	0x71: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: C},
		},
	},
	0x72: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: D},
		},
	},
	0x73: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: E},
		},
	},
	0x74: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: H},
		},
	},
	0x75: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: L},
		},
	},
	0x76: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: HL, Deref: true},
		},
	},
	0x77: {
		BIT,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: A},
		},
	},
	0x78: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: B},
		},
	},
	0x79: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: C},
		},
	},
	0x7A: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: D},
		},
	},
	0x7B: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: E},
		},
	},
	0x7C: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: H},
		},
	},
	0x7D: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: L},
		},
	},
	0x7E: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: HL, Deref: true},
		},
	},
	0x7F: {
		BIT,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: A},
		},
	},
	0x80: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: B},
		},
	},
	0x81: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: C},
		},
	},
	0x82: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: D},
		},
	},
	0x83: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: E},
		},
	},
	0x84: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: H},
		},
	},
	0x85: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: L},
		},
	},
	0x86: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: HL, Deref: true},
		},
	},
	0x87: {
		RES,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: A},
		},
	},
	0x88: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: B},
		},
	},
	0x89: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: C},
		},
	},
	0x8A: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: D},
		},
	},
	0x8B: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: E},
		},
	},
	0x8C: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: H},
		},
	},
	0x8D: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: L},
		},
	},
	0x8E: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: HL, Deref: true},
		},
	},
	0x8F: {
		RES,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: A},
		},
	},
	0x90: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: B},
		},
	},
	0x91: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: C},
		},
	},
	0x92: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: D},
		},
	},
	0x93: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: E},
		},
	},
	0x94: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: H},
		},
	},
	0x95: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: L},
		},
	},
	0x96: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: HL, Deref: true},
		},
	},
	0x97: {
		RES,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: A},
		},
	},
	0x98: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: B},
		},
	},
	0x99: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: C},
		},
	},
	0x9A: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: D},
		},
	},
	0x9B: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: E},
		},
	},
	0x9C: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: H},
		},
	},
	0x9D: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: L},
		},
	},
	0x9E: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: HL, Deref: true},
		},
	},
	0x9F: {
		RES,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: A},
		},
	},
	0xA0: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: B},
		},
	},
	0xA1: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: C},
		},
	},
	0xA2: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: D},
		},
	},
	0xA3: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: E},
		},
	},
	0xA4: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: H},
		},
	},
	0xA5: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: L},
		},
	},
	0xA6: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: HL, Deref: true},
		},
	},
	0xA7: {
		RES,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: A},
		},
	},
	0xA8: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: B},
		},
	},
	0xA9: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: C},
		},
	},
	0xAA: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: D},
		},
	},
	0xAB: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: E},
		},
	},
	0xAC: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: H},
		},
	},
	0xAD: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: L},
		},
	},
	0xAE: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: HL, Deref: true},
		},
	},
	0xAF: {
		RES,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: A},
		},
	},
	0xB0: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: B},
		},
	},
	0xB1: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: C},
		},
	},
	0xB2: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: D},
		},
	},
	0xB3: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: E},
		},
	},
	0xB4: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: H},
		},
	},
	0xB5: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: L},
		},
	},
	0xB6: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: HL, Deref: true},
		},
	},
	0xB7: {
		RES,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: A},
		},
	},
	0xB8: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: B},
		},
	},
	0xB9: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: C},
		},
	},
	0xBA: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: D},
		},
	},
	0xBB: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: E},
		},
	},
	0xBC: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: H},
		},
	},
	0xBD: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: L},
		},
	},
	0xBE: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: HL, Deref: true},
		},
	},
	0xBF: {
		RES,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: A},
		},
	},
	0xC0: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: B},
		},
	},
	0xC1: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: C},
		},
	},
	0xC2: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: D},
		},
	},
	0xC3: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: E},
		},
	},
	0xC4: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: H},
		},
	},
	0xC5: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: L},
		},
	},
	0xC6: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: HL, Deref: true},
		},
	},
	0xC7: {
		SET,
		[]Operand{
			{Symbol: Byte(0)},
			{Symbol: A},
		},
	},
	0xC8: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: B},
		},
	},
	0xC9: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: C},
		},
	},
	0xCA: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: D},
		},
	},
	0xCB: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: E},
		},
	},
	0xCC: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: H},
		},
	},
	0xCD: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: L},
		},
	},
	0xCE: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: HL, Deref: true},
		},
	},
	0xCF: {
		SET,
		[]Operand{
			{Symbol: Byte(1)},
			{Symbol: A},
		},
	},
	0xD0: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: B},
		},
	},
	0xD1: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: C},
		},
	},
	0xD2: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: D},
		},
	},
	0xD3: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: E},
		},
	},
	0xD4: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: H},
		},
	},
	0xD5: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: L},
		},
	},
	0xD6: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: HL, Deref: true},
		},
	},
	0xD7: {
		SET,
		[]Operand{
			{Symbol: Byte(2)},
			{Symbol: A},
		},
	},
	0xD8: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: B},
		},
	},
	0xD9: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: C},
		},
	},
	0xDA: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: D},
		},
	},
	0xDB: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: E},
		},
	},
	0xDC: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: H},
		},
	},
	0xDD: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: L},
		},
	},
	0xDE: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: HL, Deref: true},
		},
	},
	0xDF: {
		SET,
		[]Operand{
			{Symbol: Byte(3)},
			{Symbol: A},
		},
	},
	0xE0: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: B},
		},
	},
	0xE1: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: C},
		},
	},
	0xE2: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: D},
		},
	},
	0xE3: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: E},
		},
	},
	0xE4: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: H},
		},
	},
	0xE5: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: L},
		},
	},
	0xE6: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: HL, Deref: true},
		},
	},
	0xE7: {
		SET,
		[]Operand{
			{Symbol: Byte(4)},
			{Symbol: A},
		},
	},
	0xE8: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: B},
		},
	},
	0xE9: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: C},
		},
	},
	0xEA: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: D},
		},
	},
	0xEB: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: E},
		},
	},
	0xEC: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: H},
		},
	},
	0xED: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: L},
		},
	},
	0xEE: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: HL, Deref: true},
		},
	},
	0xEF: {
		SET,
		[]Operand{
			{Symbol: Byte(5)},
			{Symbol: A},
		},
	},
	0xF0: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: B},
		},
	},
	0xF1: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: C},
		},
	},
	0xF2: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: D},
		},
	},
	0xF3: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: E},
		},
	},
	0xF4: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: H},
		},
	},
	0xF5: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: L},
		},
	},
	0xF6: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: HL, Deref: true},
		},
	},
	0xF7: {
		SET,
		[]Operand{
			{Symbol: Byte(6)},
			{Symbol: A},
		},
	},
	0xF8: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: B},
		},
	},
	0xF9: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: C},
		},
	},
	0xFA: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: D},
		},
	},
	0xFB: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: E},
		},
	},
	0xFC: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: H},
		},
	},
	0xFD: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: L},
		},
	},
	0xFE: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: HL, Deref: true},
		},
	},
	0xFF: {
		SET,
		[]Operand{
			{Symbol: Byte(7)},
			{Symbol: A},
		},
	},
}

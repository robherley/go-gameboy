package instructions

// https://gbdev.io/gb-opcodes/optables/
// https://gbdev.io/gb-opcodes/Opcodes.json
// http://gameboy.mongenel.com/dmg/opcodes.html

type Instruction struct {
	Mnemonic
	// TODO
}

var unprefixed = [...]*Instruction{}

var cbPrefixed = [...]*Instruction{}

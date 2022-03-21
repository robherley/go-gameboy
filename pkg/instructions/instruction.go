package instructions

type Operands = []interface{}

type Instruction struct {
	Mnemonic
	Operands
}

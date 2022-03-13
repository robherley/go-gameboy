package processor

import "github.com/robherley/go-dmg/pkg/instructions"

type Processor struct {
	Registers
	Instruction *instructions.Instruction
}

package dmg

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/cartridge"
	"github.com/robherley/go-dmg/pkg/processor"
)

type Emulator struct {
	Cartridge *cartridge.Cartridge
	Processor *processor.Processor
}

func New() *Emulator {
	return &Emulator{
		Cartridge: nil,
		Processor: nil,
	}
}

func (e *Emulator) Boot() {
	fmt.Println("beep boop")
}

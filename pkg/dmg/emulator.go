package dmg

import (
	"fmt"

	"github.com/robherley/go-dmg/pkg/cart"
)

type Emulator struct {
	Cartridge *cart.Cartridge
	Processor Processor
	Graphics  Graphics
	Clock     Clock
}

func New() *Emulator {
	return &Emulator{
		Cartridge: nil,
		Processor: nil,
		Graphics:  nil,
		Clock:     nil,
	}
}

func (e *Emulator) Boot() {
	fmt.Println("beep boop")
}

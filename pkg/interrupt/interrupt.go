package interrupt

import errs "github.com/robherley/go-gameboy/pkg/errors"

type InterruptMasterChange byte

const (
	MASTER_SET_NONE InterruptMasterChange = 0
	MASTER_SET_NOW  InterruptMasterChange = 1
	MASTER_SET_NEXT InterruptMasterChange = 2
)

// https://gbdev.io/pandocs/Interrupts.html
type Interrupt struct {
	// (IME) MasterEnabled is used to disabled all interrupts on the IE register
	MasterEnabled bool
	// EI: sets master to be enabled (delayed one instruction)
	EI InterruptMasterChange
	// DI: sets master to be disabled (delayed one instruction)
	DI InterruptMasterChange
	// (IF - $FFOF) flag identifies if a specific interrupt bit becomes set
	Flag byte
	// (IE - $FFFF) enable specifies if a specific interrupt bit is enabled
	Enable byte
}

func New() *Interrupt {
	return &Interrupt{
		MasterEnabled: false,
		EI:            MASTER_SET_NONE,
		DI:            MASTER_SET_NONE,
		Flag:          0x0,
		Enable:        0x0,
	}
}

func (i *Interrupt) Read(address uint16) byte {
	switch address {
	case FLAG_ADDRESS:
		return i.Flag
	case ENABLE_ADDRESS:
		return i.Enable
	default:
		panic(errs.NewReadError(address, "interrupt"))
	}
}

func (i *Interrupt) Write(address uint16, data byte) {
	switch address {
	case FLAG_ADDRESS:
		i.Flag = data
	case ENABLE_ADDRESS:
		i.Enable = data
	default:
		panic(errs.NewWriteError(address, "interrupt"))
	}
}

func (i *Interrupt) Requested() bool {
	return i.Flag != 0
}

func (i *Interrupt) Triggered(t Type) bool {
	return i.Enabled(t) && i.Flagged(t)
}

func (i *Interrupt) Enabled(t Type) bool {
	return i.Enable&byte(t) != 0
}

func (i *Interrupt) Flagged(t Type) bool {
	return i.Flag&byte(t) != 0
}

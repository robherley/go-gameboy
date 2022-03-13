package cpu

type CPU struct {
	*Registers
}

func New() *CPU {
	return &CPU{
		&Registers{
			PC: 0x100,
		},
	}
}

package mmu

import "fmt"

type noop struct {
	strict bool
}

func newNoop(strict bool) *noop {
	return &noop{strict}
}

func (n *noop) Read(address uint16) byte {
	if n.strict {
		fmt.Printf("NOOP[%04X] read", address)
		panic("noop: strict mode")
	}
	return 0x0
}

func (n *noop) Write(address uint16, data byte) {
	if n.strict {
		fmt.Printf("NOOP[%04X] write: %02X\n", address, data)
		panic("noop: strict mode")
	}
}

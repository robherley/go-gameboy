package mmu

import "fmt"

type noop struct {
	debug bool
}

func newNoop(debug bool) *noop {
	return &noop{debug}
}

func (n *noop) Read(address uint16) byte {
	dummy := byte(0x0)

	if n.debug {
		fmt.Printf("NOOP[%04X] read: %02X \n", address, dummy)
	}
	return dummy
}

func (n *noop) Write(address uint16, data byte) {
	if n.debug {
		fmt.Printf("NOOP[%04X] write: %02X\n", address, data)
	}
}

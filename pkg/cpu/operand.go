package cpu

type Operand struct {
	// Symbol defines what kind of operand, and where to resolve the data
	Symbol Symbol
	// Deref indicates a dereference of a pointer, ie: HL vs (HL)
	Deref bool
	// Inc indicates an increment, used for LDI, ie: HL vs HL+
	Inc bool
	// Dec indicates a decrement, used for LDD, ie: HL vs HL-
	Dec bool
}

// Size returns the expected size (n=byte) for the operand, does not account for dereference
func (o *Operand) Size() byte {
	switch o.Symbol {
	case AF, BC, DE, HL, PC, SP, D16, A16:
		return 2
	default:
		return 1
	}
}

func (o *Operand) Is8() bool {
	return o.Size() == 1
}

func (o *Operand) Is16() bool {
	return o.Size() == 2
}

func (o *Operand) IsRegister() bool {
	_, ok := o.Symbol.(Register)
	return ok
}

func (o *Operand) IsData() bool {
	_, ok := o.Symbol.(Data)
	return ok
}

func (o *Operand) IsConditon() bool {
	_, ok := o.Symbol.(Condition)
	return ok
}

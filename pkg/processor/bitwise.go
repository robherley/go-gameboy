package processor

// lower byte
func lo(v uint16) byte {
	return byte(v & 0x00FF)
}

// higher byte
func hi(v uint16) byte {
	return byte(v >> 8)
}

// combine hi and low for u16
func toU16(hi, lo byte) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

// gets the Nth bit from a byte
func getNBit(num, n byte) bool {
	return (num & (1 << n)) != 0
}

// sets the Nth bit from a byte
func setNBit(num, n byte) byte {
	return num | (1 << n)
}

// clears the Nth bit from a byte
func clearNBit(num, n byte) byte {
	return num & ^(1 << n)
}

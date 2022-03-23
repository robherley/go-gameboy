package bits

// Gets the lower byte of uint16
func Lo(v uint16) byte {
	return byte(v & 0x00FF)
}

// Returns the higher byte of uint16
func Hi(v uint16) byte {
	return byte(v >> 8)
}

// Combines hi and lo bytes to a uint16
func To16(hi, lo byte) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

// Returns the Nth bit from a byte
func GetNBit(num, bit byte) bool {
	return (num & (1 << bit)) != 0
}

// Sets the Nth bit in a byte and returns the new byte
func SetNBit(num, bit byte) byte {
	return num | (1 << bit)
}

// Clears the Nth bit in a byte and returns the new byte
func ClearNBit(num, bit byte) byte {
	return num & ^(1 << bit)
}

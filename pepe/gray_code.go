package pepe

func BinaryToGray32(n uint32) uint32 {
	return (n >> 1) ^ n
}

func GrayToBinary32(n uint32) uint32 {
	n ^= n >> 16
	n ^= n >> 8
	n ^= n >> 4
	n ^= n >> 2
	n ^= n >> 1
	return n
}

func BinaryToGray(n uint64) uint64 {
	return (n >> 1) ^ n
}

func GrayToBinary(n uint64) uint64 {
	for shift := 1; shift < 32; shift <<= 1 {
		n ^= n >> uint(shift)
	}
	return n
}

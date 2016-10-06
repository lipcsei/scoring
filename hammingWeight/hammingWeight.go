package hammingweight

func Of(n int) int {
	// each bit in n is a one-bit integer that indicates how many bits are set
	// in that bit.

	n = ((n & 0xAAAAAAAA) >> 1) + (n & 0x55555555)
	// Now every two bits are a two bit integer that indicate how many bits were
	// set in those two bits in the original number

	n = ((n & 0xCCCCCCCC) >> 2) + (n & 0x33333333)
	// Now we're at 4 bits

	n = ((n & 0xF0F0F0F0) >> 4) + (n & 0x0F0F0F0F)
	// 8 bits

	n = ((n & 0xFF00FF00) >> 8) + (n & 0x00FF00FF)
	// 16 bits

	n = ((n & 0xFFFF0000) >> 16) + (n & 0x0000FFFF)
	// kaboom - 32 bits

	return int(n)
}

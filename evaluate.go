package nilsimsa

import (
	"fmt"
)

func Evaluate(a, b []byte) (bitDiffCount int, err error) {
	var c []byte
	if c, err = XOR(a, b); err != nil {
		return
	}
	for _, v := range c {
		bitDiffCount += countBits(v)
	}
	return bitDiffCount, nil
}

// Count the number of `1` bits in a byte
func countBits(x byte) int {
	count := 0
	for x > 0 {
		count += int(x & 1) // Add the last bit
		x >>= 1             // Right shift the bits
	}
	return count
}

// XOR is a simple bitwise XOR on two nilsimsa hashes. This is used in the simple distance count
// in Evaluate but could also be used against multiple samples of text from the same author to
// create a metric of the variety of their text by XORing all of them together.
func XOR(a, b []byte) (c []byte, err error) {
	// Check if the lengths of the strings are the same
	if len(a) != len(b) {
		return nil, fmt.Errorf("byte strings are of different lengths %d and %d", len(a), len(b))
	}
	if len(a) != 32 || len(b) != 32 {
		return nil, fmt.Errorf("input nilsimsa hashes must be 32 bytes each, got %d and %d",
			len(a), len(b))
	}
	c = make([]byte, 32)
	for i := 0; i < len(a); i++ {
		// XOR the bytes and count the number of `1` bits
		c[i] = a[i] ^ b[i]
	}
	return
}

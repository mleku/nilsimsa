package nilsimsa

import (
	"fmt"
)

func Evaluate(a, b []byte) (int, error) {
	// Check if the lengths of the strings are the same
	if len(a) != len(b) {
		return 0, fmt.Errorf("byte strings are of different lengths %d and %d", len(a), len(b))
	}
	if len(a) != 32 || len(b) != 32 {
		return 0, fmt.Errorf("input nilsimsa hashes must be 32 bytes each, got %d and %d",
			len(a), len(b))
	}

	// Count the differing bits
	bitDiffCount := 0
	for i := 0; i < len(a); i++ {
		// XOR the bytes and count the number of `1` bits
		xor := a[i] ^ b[i]
		bitDiffCount += countBits(xor)
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

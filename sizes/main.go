package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Maximum values of integers\n\n")
	sizes := []uint{8, 16, 32, 64}
	for _, size := range sizes {
		fmt.Printf("max int%d : %d\n", size, maxInt(size))
		fmt.Printf("max uint%d: %d\n\n", size, maxUInt(size))
	}

}

func maxInt(nbits uint) int64 {
	x := int64(1 << (nbits - 1))
	if x < 0 {
		x = -1 * (x + 1)
	}
	return x
}
func maxUInt(nbits uint) uint64 {
	return uint64((1 << nbits) - 1)
}

package popcount

import "fmt"

// pc[i] is the population count of 1
var pc [256]byte

func init(){
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	fmt.Printf("Value is %d\n", x)
	fmt.Printf("byte(%d>>(0*8)) is %d\n", x, byte(x>>(0*8)))
	y := byte(x >> (0 * 8))
	return int(pc[y] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
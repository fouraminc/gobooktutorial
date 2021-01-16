package main

import "fmt"

const boilingf = 212.0

func main() {

	var f = boilingf
	var c = (f -32) * 5 / 9
	fmt.Printf("boiling point = %g °F or %g °C\n", f, c)

}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(strings.Join(os.Args[0:1], " "))

}

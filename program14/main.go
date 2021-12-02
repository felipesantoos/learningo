package main

import (
	"fmt"
)

type inteiro int

var x inteiro

func main() {
	fmt.Printf("x: %v, %T\n", x, x)

	x = 42

	fmt.Println(x)
}

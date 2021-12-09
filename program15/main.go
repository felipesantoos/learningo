package main

import (
	"fmt"
)

type inteiro int

var x inteiro
var y int

func main() {
	fmt.Printf("x: %v, %T\n", x, x)

	x = 42

	fmt.Println(x)

	y = int(x)
	fmt.Printf("y: %v, %T\n", y, y)
}

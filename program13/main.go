package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("É %v que %v tem %v anos!", z, y, x)

	fmt.Println(s)
}

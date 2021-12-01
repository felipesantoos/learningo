package main

import (
	"fmt"
)

func main() {
	// Interpreted string literals.
	interpreted := "A\nB\tC \"D\" \\E\\"
	fmt.Println(interpreted)

	// Raw string literals.
	raw := `A\nB\tC \"D\" \\E\\`
	fmt.Println(raw)
}

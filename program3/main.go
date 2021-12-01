package main

import (
	"fmt"
)

var pkgLvlScp = "Package Level Scope"

func main() {
	varDaMain := "Variável declarada na main"
	outraFunc(varDaMain)
}

func outraFunc(arg string) {
	// fmt.Println(varDaMain)
	fmt.Println(pkgLvlScp)
	fmt.Println(arg)
}

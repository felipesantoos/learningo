package main

import "fmt"

var inteiro int
var texto string
var booleano bool
var decimal float64
var matriz [3]int
var fatia []interface{}
var mapa map[string]string

type estrutura struct {
	inteiro  int
	texto    string
	booleano bool
}

func main() {
	inteiro = 1
	texto = "Felipe da Silva Santos"
	booleano = true
	decimal = 3.1415926535
	matriz[0] = 0
	matriz[1] = 1
	matriz[2] = 2
	fatia = append(fatia, 0)
	fatia = append(fatia, 1)
	fatia = append(fatia, 2)
	mapa = make(map[string]string)
	mapa["x"] = "y"
	est := estrutura{
		inteiro:  -1,
		texto:    "Santos Silva da Felipe",
		booleano: false,
	}

	fmt.Println(est)
}

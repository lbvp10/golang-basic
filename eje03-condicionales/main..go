package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/**
if condicinal{
}
- Las llaves deben estar en la misma linea
- Se puede realizar una asignacion y una condicion en la misma linea
*/

func main() {
	condicionalIf()
	condicionalSwitch()
}

func condicionalSwitch() {
	numero, _ := rand.Int(rand.Reader, big.NewInt(4))
	switch numero.Int64() {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(0)
	default:

	}
}

func condicionalIf() {
	//Condicional sencillo
	var isMax bool = true
	if isMax {
		fmt.Println("Is max  true")
	} else {
		fmt.Println("Is max  false")
	}
	//Se puede realizar una asignacion y una condicion en la misma linea
	if isMax = false; isMax {
		fmt.Println("Is max  true")
	} else {
		fmt.Println("Is max  false")
	}
}

package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		//i += 1
	}

	//Ciclo infinito con break

	for {
		fmt.Println("Digite el numero")
		var numero int
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}

	}

	//go to
	y := 0
GOTOPRUEBA:
	fmt.Println("ESte es el inicio del goto")
	for p := 0; p < 5; p++ {
		if y != 2 {
			y++
			goto GOTOPRUEBA
		}

	}

}

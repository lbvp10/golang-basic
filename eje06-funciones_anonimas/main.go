package main

import (
	"fmt"
	"strings"
)

func main() {
	opcion := strings.ToLower(leerOpcion())
	operacion := obtenerOperacion(opcion)
	if operacion == nil {
		println("Seleccione una opcion correcta...")
	} else {
		println("El resultado de la operacion es  : ", operacion(10, 5))
	}

}

func obtenerOperacion(opcion string) func(int, int) int {
	var operacion func(int, int) int
	switch {
	case opcion == "a":
		operacion = func(num1 int, num2 int) int { return num1 + num2 }
	case opcion == "b":
		operacion = func(num1 int, num2 int) int { return num1 - num2 }
	case opcion == "c":
		operacion = func(num1 int, num2 int) int { return num1 * num2 }
	case opcion == "d":
		operacion = func(num1 int, num2 int) int { return num1 / num2 }
	default:
		operacion = nil
	}
	return operacion
}

func leerOpcion() string {
	var opcion string
	fmt.Println("Digite una opcion \n a: Sumar  \n b: Restar \n c: Multiplicar \n d: Dividir")
	fmt.Scanln(&opcion)
	return opcion
}

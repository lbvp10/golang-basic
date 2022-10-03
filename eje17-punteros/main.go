package main

import "fmt"

func main() {
	a, b, c := 1, true, "Sierva María"
	pasoDeVariablesPorValor(a, b, c)
	fmt.Printf("En go la variables pasan por valor a:%d b:%v c:%s", a, b, c)
	fmt.Println("------------------------------------------------------------")
	pasoDeVariablesPorReferenciaPunteros(&a, &b, &c)
	fmt.Printf("En go la variables pasan por por referencia con punteros a:%d b:%v c:%s", a, b, c)
}

func pasoDeVariablesPorReferenciaPunteros(a *int, b *bool, c *string) {
	*a = 2
	*b = false
	*c = "de Todos los Ángeles"
	return
}

func pasoDeVariablesPorValor(a int, b bool, c string) {
	a = 2
	b = false
	c = "de Todos los Ángeles"
	return
}

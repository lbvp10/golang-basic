package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, e := strconv.Atoi("12e")
	if e != nil {
		fmt.Printf("El error es de tipo: %T resultado: %V", e, result)
		return
	}
	fmt.Printf("el numero es:%d", result)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	alumnos := [5]string{"Carlos", "Luis", "Juan", "Ana", "Maria"}
	//sort.Strings(alumnos)
	for i := 0; i < len(alumnos); i++ {
		println(alumnos[i])
	}
	fmt.Println("-------------------------")
	slice()
}

func slice() {
	alumnos := []string{"Carlos", "Luis", "Juan", "Ana", "Maria"}
	alumnos = append(alumnos, "Jose")
	sort.Strings(alumnos)
	for i := 0; i < len(alumnos); i++ {
		println(alumnos[i])
	}
}

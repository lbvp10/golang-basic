package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	//w := &sync.WaitGroup{}
	go deletrearPalabra("Perro camina despacio")
	go imprimirNumeros(12)
	time.Sleep(5000 * time.Millisecond)
}
func deletrearPalabra(palabra string) {
	for _, l := range strings.Split(palabra, "") {
		log.Println(l)
		time.Sleep(400 * time.Millisecond)
	}
}
func imprimirNumeros(numeros int) {
	fmt.Println("Numeros")
	for i := 0; i < numeros; i++ {
		log.Println(i)
		time.Sleep(400 * time.Millisecond)
	}
}
func imprimirNumeros2(numeros int) {
	fmt.Println("Prueba suggestion Github")
	for i := 0; i < numeros; i++ {
		log.Println(i)
		time.Sleep(400 * time.Millisecond)
	}
}
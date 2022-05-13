package main

import (
	"crypto/rand"
	"math"
	"math/big"
)

func main() {
	println(cuadradoMNumero(5))
	a, b := generarDosNumeroAleatorios()
	println("", a, b)

	raizCuadradaNumeros(10, 69, 8, 7, 6)
}

func generarDosNumeroAleatorios() (int64, int64) {
	a, _ := rand.Int(rand.Reader, big.NewInt(1000))
	b, _ := rand.Int(rand.Reader, big.NewInt(1000))

	return a.Int64(), b.Int64()
}

func cuadradoMNumero(numero int) (result int) {
	result = int(math.Pow(float64(numero), 2))
	return
}

func raizCuadradaNumeros(numeros ...int) {
	for _, numero := range numeros {
		println(math.Sqrt(float64(numero)))
	}
}

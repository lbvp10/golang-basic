package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/**
Tipos de Datos:
	bool = Boolean quedan representados por la palabra reservada bool
	int8 =  -128 - 127
	int16 = -32768 - 32767
	int32 = -2147483648 - 2147483647
	int64 = -9223372036854775808 - 9223372036854775807
Si no especificar un tipo concreto, Go utilizará int como tipo. En este caso, el tipo int representará un entero
con signo de 32 bits ó 64 bits dependiendo de la plataforma en la que se esté ejecutando la aplicación #prueba()
	uint8 =  0 - 255
	uint16 = 0 - 65535
	uint32 = 0 - 424967295
	uint64 = 0 - 18446744073709551615

	float32 =  32 bit single precision
	float64 =  64 bit doble precision

	Alias:
		byte=unit8
		rune=int32
	string
*/

//Variable local
var numero1 = 3

//Variable global
var NumeroGlobal = 10

func main() {
	//Tamaño por defecto de un int 32 o 64 dependiento del procesador
	var year int = -1
	fmt.Printf("%v => %T => %d bytes\n", year, year, unsafe.Sizeof(year))
	//cracion de variables y asignacion de valores
	numero1 := 1
	numero2, numero3 := 2, 3
	fmt.Println("numero1 , numero2, numero3 ", numero1, numero2, numero3)
	//Valores por defecto
	var numero4 int
	var nombre string
	var saldo float32
	var esMayordeEdad bool
	fmt.Println("numero4,nombre,saldo,esMayordeEdad", numero4, nombre, saldo, esMayordeEdad)

	// Byte
	fmt.Println()
	fmt.Println("Byte")
	var byteLetter byte = 'A'
	fmt.Printf("%v => %T => %d bytes\n", byteLetter, byteLetter, unsafe.Sizeof(byteLetter))

	// Rune
	fmt.Println(" \nRune")
	var runeLetter rune = 'µ'
	fmt.Printf("%v => %U => %T => %d bytes\n", runeLetter, runeLetter, runeLetter, unsafe.Sizeof(runeLetter))

	// Strings
	fmt.Println()
	fmt.Println("Strings")
	var name = "Jorge"
	var surname = ` ${name} Serrano`
	fmt.Printf("%v => %T\n", name, name)
	fmt.Printf("%v => %T\n", surname, surname)

	//bool
	validUser, validPass := true, false
	hasAccess := validUser && validPass
	fmt.Println("\nTiene acceso: ", hasAccess)

	//Casteo tipo de datos
	saldo = 20200.6
	var saldoEntero = int(saldo)
	fmt.Println("Casteo de de float a entero", saldoEntero)
	//Casteo con libreria strConv -> numero a string
	conv := strconv.Itoa(year)
	fmt.Printf("%v => %T => %d conv\n", conv, conv, unsafe.Sizeof(conv))

	var nombreVacion string = string(0)
	fmt.Printf("String vacio %v  ", nombreVacion)
}

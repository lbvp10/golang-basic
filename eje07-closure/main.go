package main

func main() {

	tabla := Tabla(2)
	for i := 0; i <= 10; i++ {
		tabla()
	}
}
func Tabla(numero int) func() int {
	contador := 0
	return func() int {

		resultado := numero * contador
		println(resultado)
		contador++
		return resultado
	}
}

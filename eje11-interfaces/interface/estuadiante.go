package _interface

import "fmt"

type Estuadiante interface {
	Persona
	Estudiar()
}

type EstudianteUniversitario struct {
	nombreUniversidad string
}

func (e *EstudianteUniversitario) Caminar() {
	fmt.Println("Este universitario puede caminar")
}
func (e *EstudianteUniversitario) Comer() {
	fmt.Println("Este universitario puede comer rapido")
}
func (e *EstudianteUniversitario) Dormir() {
	fmt.Println("Este universitario puede dormir poco")
}
func (e *EstudianteUniversitario) Estudiar() {
	e.nombreUniversidad = "Javeriana"
	fmt.Println("Este universitario puede estudiar ")
}

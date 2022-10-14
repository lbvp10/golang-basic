package main

import (
	_interface "modulos/eje11-interfaces/interface"
)

func main() {
	universitario := new(_interface.EstudianteUniversitario)
	//fmt.Println(*universitario)

	var persona _interface.Persona
	persona = universitario
	persona.Comer()

	universitario2 := persona.(_interface.Estuadiante)
	universitario2.Estudiar()

	p := universitario2.(_interface.Persona)
	p.Dormir()

}

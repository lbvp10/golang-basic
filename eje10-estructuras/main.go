package main

import (
	"./persona"
	"fmt"
	"time"
)

type Estudiante struct {
	p       persona.Persona
	colegio string
}

func main() {
	p := persona.Persona{
		FechaNacimiento: time.Now(),
		Nombre:          "Luis",
		Apellidos:       "XX",
	}
	estudiante := Estudiante{
		p:       p,
		colegio: "ITA",
	}
	fmt.Println(persona.NombreCompleto(&estudiante.p))
	fmt.Println(estudiante)
}

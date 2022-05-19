package persona

import (
	"strings"
	"time"
)

type Persona struct {
	id              int
	Nombre          string
	Apellidos       string
	telefono        string
	FechaNacimiento time.Time
}

func NombreCompleto(p *Persona) string {
	var sb strings.Builder
	sb.WriteString(p.Nombre)
	sb.WriteString(" ")
	sb.WriteString(p.Apellidos)
	return sb.String()
}

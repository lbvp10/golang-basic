package modelo

import "time"

type VacacionesRequest struct {
	FechaInicial time.Time `json:"fecha_inicial"`
	FechaFinal   time.Time `json:"fecha_final"`
	NumeroDias   int8      `json:"numero_dias"`
}

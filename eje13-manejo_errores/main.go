package main

import (
	"log"
	"os"
)

func main() {

	fichero, error := os.OpenFile("", os.O_RDONLY, os.ModeExclusive)
	if error != nil {
		panic("Un fichero base del sistema no existe ")
	}
	fichero.WriteString("Prueba de file.....")

	defer func() {
		rec := recover()
		fichero.Close()
		log.Fatalf("un error fatal...... \n %v", rec)
	}()

}

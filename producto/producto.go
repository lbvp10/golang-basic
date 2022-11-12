package producto

import (
	"gorm.io/gorm"
	"log"
	"time"
)
import "orm/db"

type Producto struct {
	gorm.Model
	Fecha       time.Time `json:"fecha"`
	Descripcion string    `json:"descripcion"`
}

func GetAll() *[]Producto {
	var feriados []Producto
	ddb, _ := db.ConnDB()
	result := ddb.Find(&feriados)
	log.Printf("Se encontraron %d registros", result.RowsAffected)
	if result.Error != nil {
		log.Printf("error consultando los feriasos %v\n", result.Error)
	}
	return &feriados
}

func GetById(id uint) *Producto {
	producto := Producto{Model: gorm.Model{ID: id}}
	ddb, _ := db.ConnDB()
	result := ddb.First(&producto)
	if result.Error != nil {
		log.Printf("Error consultando la db %v", result.Error)
	}
	return &producto
}

func Save(producto *Producto) {
	ddb, _ := db.ConnDB()
	result := ddb.Create(&producto)
	if result.Error != nil {
		log.Printf("Error consultando la db %v", result.Error)
	}
}

func Update(producto *Producto) {
	ddb, _ := db.ConnDB()
	result := ddb.Save(&producto)
	if result.Error != nil {
		log.Printf("Error consultando la db %v", result.Error)
	}
}

func Delete(id uint) {
	ddb, _ := db.ConnDB()
	ddb.Delete(&Producto{}, id)
}

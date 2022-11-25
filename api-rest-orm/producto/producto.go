package producto

import (
	"gorm.io/gorm"
	"orm/db"
	"time"
)

type Producto struct {
	gorm.Model
	Fecha       time.Time `json:"fecha"`
	Descripcion string    `json:"descripcion"`
}

func GetAll() (*[]Producto, error) {
	var feriados []Producto
	ddb, _ := db.ConnDB()
	result := ddb.Find(&feriados)

	return &feriados, result.Error
}

func GetById(id uint) (*Producto, error) {
	producto := Producto{Model: gorm.Model{ID: id}}
	ddb, _ := db.ConnDB()
	result := ddb.First(&producto)

	return &producto, result.Error
}

func Save(producto *Producto) error {
	ddb, _ := db.ConnDB()
	result := ddb.Create(&producto)
	return result.Error
}

func Update(producto *Producto) error {
	ddb, _ := db.ConnDB()
	result := ddb.Save(&producto)
	return result.Error
}

func Delete(id uint) (error, int64) {
	ddb, _ := db.ConnDB()
	result := ddb.Delete(&Producto{}, id)
	return result.Error, result.RowsAffected
}

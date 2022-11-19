package producto

import (
	"gorm.io/gorm"
	"orm/api"
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

func Save(producto *Producto) *gorm.DB {
	ddb, _ := db.ConnDB()
	result := ddb.Create(&producto)
	return result
}

func Update(producto *Producto) *gorm.DB {
	ddb, _ := db.ConnDB()
	result := ddb.Save(&producto)
	return result
}

func Delete(id uint) error {
	ddb, _ := db.ConnDB()
	result := ddb.Delete(&Producto{}, id)
	if result.Error == nil && result.RowsAffected < 1 {
		return api.ErrEntityNotFound
	}
	return result.Error
}

package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"net/http"
	error2 "orm/api"
	"orm/logger"
	p "orm/producto"
)

func ConfigServer(app *fiber.App) {
	// Middleware
	router := app.Group("/api", requestid.New(), error2.LogApi)
	addRutas(router)
}

func addRutas(api fiber.Router) {

	api.Get("/", getHandler)
	api.Get("/:id<int>", getByIdHandler)
	api.Post("/", postHandler)
	api.Put("/", putHandler)
	api.Delete("/:id<int>", deleteHandler)
}

func getByIdHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	producto, err := p.GetById(uint(id))
	if err != nil {
		logger.Error(fmt.Sprintf("Error getById de la db %v", err))
		return ctx.Status(http.StatusInternalServerError).JSON(error2.ErrorApi{Message: "Error getById de la db", Code: "1005"})
	}
	return ctx.JSON(producto)
}

func getHandler(ctx *fiber.Ctx) error {
	productos, err := p.GetAll()
	if err != nil {
		logger.Error(fmt.Sprintf("api consultando los productos %v\n", err))
		return ctx.Status(http.StatusInternalServerError).JSON(error2.ErrorApi{Message: "Error get en la db", Code: "1004"})
	}
	return ctx.JSON(productos)

}
func postHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	result := p.Save(producto)
	if result.Error != nil {
		logger.Error(fmt.Sprintf("Error guardando la db %v", result.Error))
		return ctx.Status(http.StatusInternalServerError).JSON(error2.ErrorApi{Message: "Error guardando en la db", Code: "1003"})
	}
	return ctx.JSON(producto)
}
func putHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	result := p.Update(producto)
	if result.Error != nil {
		logger.Error(fmt.Sprintf("Error actualizando la db %v", result.Error))
		return ctx.Status(http.StatusInternalServerError).JSON(error2.ErrorApi{Message: "Error actualizando en la db", Code: "1002"})
	}
	return ctx.JSON(producto)
}
func deleteHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	error := p.Delete(uint(id))

	if error != nil {
		logger.Error(fmt.Sprintf("Error delete %v", error))
		msgError, codeError := "Error delete", http.StatusInternalServerError
		if error == error2.ErrEntityNotFound {
			msgError, codeError = error2.ErrEntityNotFound.Error(), http.StatusNotFound
		}
		return ctx.Status(codeError).JSON(error2.ErrorApi{Message: msgError, Code: "1001"})
	}

	return ctx.SendStatus(200)

}

func loadProductByRequest(c *fiber.Ctx) (*p.Producto, error) {
	producto := p.Producto{}
	var err error
	if err = c.BodyParser(&producto); err != nil {
		logger.Error(fmt.Sprintf("An api occured: %v", err))
		return nil, c.Status(http.StatusBadRequest).JSON(error2.ErrorApi{Message: "Error body", Code: "1007"})
	}

	return &producto, err
}

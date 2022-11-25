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
		errorApi := error2.NewInternalServerApiError(err, "Error getById de la db")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	} else if producto == nil {
		errorApi := error2.NewNotFoundApiError("")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}

func getHandler(ctx *fiber.Ctx) error {
	productos, err := p.GetAll()
	if err != nil {
		errorApi := error2.NewInternalServerApiError(err, "api consultando los productos")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(productos)

}
func postHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	err := p.Save(producto)
	if err != nil {
		errorApi := error2.NewInternalServerApiError(err, "Error guardando en la db")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}
func putHandler(ctx *fiber.Ctx) error {
	producto, _ := loadProductByRequest(ctx)
	err := p.Update(producto)
	if err.Error != nil {
		errorApi := error2.NewInternalServerApiError(err, "Error actualizando la db")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	}
	return ctx.JSON(producto)
}

func deleteHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err, rowsDeleted := p.Delete(uint(id))

	if err != nil {
		errorApi := error2.NewInternalServerApiError(err, "Error Delete")
		return ctx.Status(errorApi.Status).JSON(errorApi)
	} else if rowsDeleted < 1 {
		errorApi := error2.NewNotFoundApiError("Delete not found")
		return ctx.Status(errorApi.Status).JSON(errorApi)

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

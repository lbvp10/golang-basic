package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	p "orm/producto"
)

const LOG_FORMAT = "${pid} [${ip}]:${port} ${locals:requestid} ${latency} ${status} - ${method} ${path} ${reqHeaders} ${body} ${resBody} ${yellow}\n"

func ConfigServer(app *fiber.App) {
	// Middleware
	router := app.Group("/api", requestid.New(), Log)
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
	producto := p.GetById(uint(id))

	return ctx.JSON(producto)
}

func getHandler(c *fiber.Ctx) error {
	productos := p.GetAll()
	return c.JSON(productos)

}
func postHandler(c *fiber.Ctx) error {
	producto := loadProductByRequest(c)
	p.Save(producto)
	return c.JSON(producto)
}
func putHandler(c *fiber.Ctx) error {
	producto := loadProductByRequest(c)
	p.Update(producto)
	return c.JSON(producto)
}
func deleteHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	p.Delete(uint(id))
	return ctx.JSON(200)
}

func loadProductByRequest(c *fiber.Ctx) *p.Producto {
	producto := p.Producto{}
	if err := c.BodyParser(&producto); err != nil {
		log.Printf("An error occured: %v", err)
	}
	return &producto
}

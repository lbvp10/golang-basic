package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	p "orm/producto"
)

func AddRutas(app *fiber.App) {
	app.Get("/", getHandler)
	app.Get("/:id<int>", getByIdHandler)
	app.Post("/", postHandler)
	app.Put("/", putHandler)
	app.Delete("/:id<int>", deleteHandler)
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
	p.Save(loadProductByRequest(c))
	return c.Redirect("/")
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
		log.Fatalf("An error occured: %v", err)
	}
	return &producto
}

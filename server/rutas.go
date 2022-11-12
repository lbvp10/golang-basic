package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"vacaciones/modelo"
)

func AddRutas(app *fiber.App) {
	app.Post("/", postHandler)

}

func postHandler(c *fiber.Ctx) error {
	request := loadVacaconesRequest(c)
	fmt.Println(request)
	return c.Redirect("/")
}

func loadVacaconesRequest(c *fiber.Ctx) *modelo.VacacionesRequest {
	request := modelo.VacacionesRequest{}
	if err := c.BodyParser(&request); err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	return &request
}

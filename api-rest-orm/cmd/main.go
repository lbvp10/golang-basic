package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orm/db"
	"orm/producto"
	"orm/server"
	"os"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	app := fiber.New()

	server.ConfigServer(app)

	db.ConnDB()
	db.ConfigShema(&producto.Producto{})

	app.Listen(fmt.Sprintf(":%v", port))
}

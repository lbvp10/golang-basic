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
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server.AddRutas(app)

	db.ConnDB()
	db.ConfigShema(&producto.Producto{})

	app.Listen(fmt.Sprintf(":%v", port))
}

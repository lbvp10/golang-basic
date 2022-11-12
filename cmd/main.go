package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"vacaciones/server"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server.AddRutas(app)

	//db.ConnDB()

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

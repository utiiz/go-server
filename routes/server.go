package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/utiiz/go-server/config"
)

func Setup(c config.Config) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/users", GetUsers)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.Server.Port)))
}

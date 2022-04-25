package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/config"
)

func Setup(c config.Config) {
	app := fiber.New()

	app.Get("/users", GetUsers)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.Server.Port)))
}

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
	app.Get("/users/:id", GetUser)
	app.Post("/users", CreateUser)
	app.Patch("/users", UpdateUser)
	app.Delete("/users/:id", DeleteUser)

	app.Get("/issues", GetIssues)
	app.Get("/issues/:id", GetIssue)
	app.Post("/issues", CreateIssue)
	app.Patch("/issues", UpdateIssue)
	app.Delete("/issues/:id", DeleteIssue)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.Server.Port)))
}

package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/config"
	"github.com/utiiz/go-server/model"
)

func GetUsers(c *fiber.Ctx) error {
	users := []model.User{
		model.User{
			Username: "utiiz",
		},
		model.User{
			Username: "mavys",
		},
	}

	return json.NewEncoder(c).Encode(users)
}

func Setup(c config.Config) {
	app := fiber.New()

	app.Get("/users", GetUsers)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.Server.Port)))
}

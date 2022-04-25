package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/models"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{
		models.User{
			Username: "utiiz",
		},
		models.User{
			Username: "mavys",
		},
	}

	return json.NewEncoder(c).Encode(users)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/models"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := models.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all users - " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

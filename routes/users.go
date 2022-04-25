package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/models"
	"github.com/utiiz/go-server/utils"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := models.GetUsers()
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all users - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing param id - " + err.Error(),
		})
	}

	user, err := models.GetUser(id)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing user - " + err.Error(),
		})
	}

	err = models.CreateUser(user)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing user - " + err.Error(),
		})
	}

	err = models.UpdateUser(user)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing param id - " + err.Error(),
		})
	}

	err = models.DeleteUser(id)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted",
	})
}

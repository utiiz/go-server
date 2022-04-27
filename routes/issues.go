package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/utiiz/go-server/models"
	"github.com/utiiz/go-server/utils"
)

func GetIssues(c *fiber.Ctx) error {
	issues, err := models.GetIssues()
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all users - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(issues)
}

func GetIssue(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing param id - " + err.Error(),
		})
	}

	issue, err := models.GetIssue(id)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(issue)
}

func CreateIssue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var issue models.Issue
	err := c.BodyParser(&issue)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing user - " + err.Error(),
		})
	}

	err = models.CreateIssue(issue)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(issue)
}

func UpdateIssue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var issue models.Issue
	err := c.BodyParser(&issue)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing user - " + err.Error(),
		})
	}

	err = models.UpdateIssue(issue)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating user - " + err.Error(),
		})
	}

	utils.Log(c, fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(issue)
}

func DeleteIssue(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		utils.Log(c, fiber.StatusInternalServerError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing param id - " + err.Error(),
		})
	}

	err = models.DeleteIssue(id)
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

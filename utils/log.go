package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func Log(c *fiber.Ctx, status int) {
	whilte := color.New(color.FgHiBlack)
	var bg *color.Color
	if status == fiber.StatusOK {
		bg = whilte.Add(color.BgGreen)
	} else {
		bg = whilte.Add(color.BgRed)
	}
	bg.Printf(" %s ", c.Method())
	fmt.Printf("%4d - %s\n", status, c.OriginalURL())
}

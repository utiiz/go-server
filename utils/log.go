package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func Log(c *fiber.Ctx, status int) {
	whilte := color.New(color.FgHiBlack)
	boldWhite := whilte.Add(color.BgGreen)
	boldWhite.Printf(" %s ", c.Method())
	fmt.Printf("%4d - %s\n", status, c.OriginalURL())
}

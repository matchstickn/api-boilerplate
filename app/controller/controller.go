package controller

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/matchstickn/REPONAME/app/service"
)

func status(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

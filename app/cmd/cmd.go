package cmd

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/REPONAME/app/router"
)

func SetUpFiber() *fiber.App {
	IPValidation := false
	if os.Getenv("ENABLE_IP_VALIDATION") == "true" {
		IPValidation = true
	}

	app := fiber.New(fiber.Config{
		AppName:            os.Getenv("APP_NAME"),
		EnableIPValidation: IPValidation,
	})
	return app
}

func SetUpRouter(app *fiber.App) {
	router.SetUp(app)
}

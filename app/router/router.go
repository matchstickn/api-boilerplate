package router

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/swagger"

	_ "github.com/matchstickn/REPONAME/app/docs"
)

// To regenerate the swagger docs run:
// swag init -g app/controller/controller.go -o ./app/docs

// @title APINAMEORWHATEVER
// @version 1.0
// @description Whatever you want
// @host localhost:PORT
// @BasePath /api/v1
func SetUp(app *fiber.App) {
	SetUpFiberMiddleware(app)
	// if os.Getenv("RATE_LIMITING") == "true" {
	// 	UseRateLimiter(app)
	// }
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}

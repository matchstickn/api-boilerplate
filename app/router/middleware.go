package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetUpFiberMiddleware(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))
	app.Use(cache.New(cache.Config{
		Expiration: time.Second * 30,
	}))
}

func UseRateLimiter(app *fiber.App) {
	api := app.Group("api/v1")
	api.Use(limiter.New(limiter.Config{
		Max:        1,
		Expiration: time.Second,
	}))
}

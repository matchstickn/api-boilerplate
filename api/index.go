package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	// "github.com/gofiber/swagger"

	// _ "github.com/matchstickn/REPONAME/app/docs"
	"github.com/matchstickn/REPONAME/app/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// for fiber.Ctx to work
	r.RequestURI = r.URL.String()

	handlerMain().ServeHTTP(w, r)
}

func handlerMain() http.HandlerFunc {
	app := fiber.New()
	router.SetUpFiberMiddleware(app)

	return adaptor.FiberApp(app)
}

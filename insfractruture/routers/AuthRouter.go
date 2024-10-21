package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/handler"
)

var (
	handlerAuth = handler.NewAuthHandler()
)

func NewAuthRouter(app *fiber.App) {
	api := app.Group("/api/auth")
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerAuth.GetUserFindByEmail(c)
	})
}

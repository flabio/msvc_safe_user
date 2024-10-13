package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/handler"
	"github.com/safe_msvc_user/middleware"
)

var (
	handlerUser = handler.NewUserHandler()
)

func NewUserRouter(app *fiber.App) {

	api := app.Group("/api/user")
	api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerUser.GetUserFindAll(c)
	})
	api.Get("/students", func(c *fiber.Ctx) error {
		return handlerUser.GetStudentsFindAll(c)
	})
	api.Get("/instructor", func(c *fiber.Ctx) error {
		return handlerUser.GetInstructorFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerUser.CreateUser(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerUser.UpdateUser(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerUser.DeleteUser(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerUser.GetUserFindById(c)
	})
}

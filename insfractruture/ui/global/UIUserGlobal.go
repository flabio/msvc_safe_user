package global

import "github.com/gofiber/fiber/v2"

type UIUserGlobal interface {
	GetUserFindAll(c *fiber.Ctx) error
	GetStudentsFindAll(c *fiber.Ctx) error
	GetInstructorFindAll(c *fiber.Ctx) error
	GetUserFindById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

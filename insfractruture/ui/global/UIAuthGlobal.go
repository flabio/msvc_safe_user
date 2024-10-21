package global

import "github.com/gofiber/fiber/v2"

type UIAuthGlobal interface {
	GetUserFindByEmail(c *fiber.Ctx) error
}

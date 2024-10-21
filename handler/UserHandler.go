package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/usecase/service"
)

type userHandler struct {
	user global.UIUserGlobal
}

func NewUserHandler() global.UIUserGlobal {
	return &userHandler{user: service.NewUserService()}
}

func (h *userHandler) GetUserFindAll(c *fiber.Ctx) error {
	return h.user.GetUserFindAll(c)
}
func (h *userHandler) GetStudentsFindAll(c *fiber.Ctx) error {
	return h.user.GetStudentsFindAll(c)
}
func (h *userHandler) GetInstructorFindAll(c *fiber.Ctx) error {
	return h.user.GetInstructorFindAll(c)
}
func (h *userHandler) GetUserFindById(c *fiber.Ctx) error {
	return h.user.GetUserFindById(c)
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	return h.user.CreateUser(c)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	return h.user.UpdateUser(c)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	return h.user.DeleteUser(c)
}

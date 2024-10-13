package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/usecase/service"
)

type authHandler struct {
	auth global.UIAuthGlobal
}

func NewAuthHandler() global.UIAuthGlobal {
	return &authHandler{auth: service.NewAuthService()}
}

func (h *authHandler) GetUserFindByEmail(c *fiber.Ctx) error {
	return h.auth.GetUserFindByEmail(c)
}

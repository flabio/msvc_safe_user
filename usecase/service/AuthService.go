package service

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/core"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
)

type authService struct {
	uiAuth uicore.UIAuthCore
}

func NewAuthService() global.UIAuthGlobal {
	return &authService{uiAuth: core.NewAuthRepository()}
}

func (s *authService) GetUserFindByEmail(c *fiber.Ctx) error {
	body := c.Body()

	var dataMap map[string]interface{}
	json.Unmarshal([]byte(body), &dataMap)

	result, err := s.uiAuth.GetUserFindByEmail(dataMap["username"].(string))
	if !utils.ComparePassword(result.Password, []byte(dataMap["password"].(string))) {
		return c.Status(fiber.StatusNotFound).JSON(utils.USER_NOT_FOUND)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ERROR_QUERY)
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.USER_NOT_FOUND)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

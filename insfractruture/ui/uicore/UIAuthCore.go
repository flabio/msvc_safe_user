package uicore

import (
	"github.com/safe_msvc_user/insfractruture/entities"
)

type UIAuthCore interface {
	GetUserFindByEmail(email string) (entities.User, error)
}

package core

import (
	"sync"

	"github.com/safe_msvc_user/insfractruture/database"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
)

func NewAuthRepository() uicore.UIAuthCore {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (db *OpenConnection) GetUserFindByEmail(email string) (entities.User, error) {
	var user entities.User
	db.mux.Lock()
	query := db.connection.Where(utils.DB_EQUAL_EMAIL_ID, email).Find(&user)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return user, query.Error
}

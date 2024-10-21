package core

import (
	"sync"

	"github.com/safe_msvc_user/insfractruture/database"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
)

func NewUserRepository() uicore.UIUserCore {
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

func (c *OpenConnection) GetUserFindAll(begin int) ([]entities.User, int64, error) {
	var userEntities []entities.User
	var countUser int64
	c.mux.Lock()
	result := c.connection.Offset(begin).Limit(5).Order(utils.DB_ORDER_DESC).Find(&userEntities)
	c.connection.Table("users").Count(&countUser)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return userEntities, countUser, result.Error
}
func (c *OpenConnection) GetStudentsFindAll(begin int) ([]entities.User, int64, error) {
	var userEntities []entities.User
	var countUser int64
	c.mux.Lock()
	result := c.connection.Offset(begin).Limit(5).Order(utils.DB_ORDER_DESC).Where("rol_id=?",3).Find(&userEntities)
	c.connection.Table("users").Where("rol_id=?",3).Count(&countUser)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return userEntities, countUser, result.Error
}
func (c *OpenConnection) GetInstructorFindAll(begin int) ([]entities.User, int64, error) {
	var userEntities []entities.User
	var countUser int64
	c.mux.Lock()
	result := c.connection.Offset(begin).Limit(5).Order(utils.DB_ORDER_DESC).Where("rol_id=?",2).Find(&userEntities)
	c.connection.Table("users").Where("rol_id=?",2).Count(&countUser)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return userEntities, countUser, result.Error
}
func (c *OpenConnection) GetUserFindById(id uint) (entities.User, error) {
	var user entities.User
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&user)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, result.Error
}

func (db *OpenConnection) GetUserFindByEmailAndId(id uint, email string) (bool, error) {

	db.mux.Lock()
	query := db.connection.Where(utils.DB_EQUAL_EMAIL_ID, email)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities.User{})
	defer db.mux.Unlock()
	defer database.CloseConnection()
	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}

func (c *OpenConnection) CreateUser(user entities.User) (entities.User, error) {
	c.mux.Lock()
	err := c.connection.Create(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, err
}
func (c *OpenConnection) UpdateUser(id uint, user entities.User) (entities.User, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, err
}

func (c *OpenConnection) DeleteUser(id uint) (bool, error) {
	c.mux.Lock()
	var user entities.User
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}

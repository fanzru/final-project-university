package repo

import (
	errs "backend/app/accounts/domain/errors"
	"backend/app/accounts/domain/models"
	"backend/infrastructure/config"
	"backend/infrastructure/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Impl interface {
	GetUserByEmail(ctx echo.Context, email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}
type AccountsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a AccountsRepo) AccountsRepo {
	return a
}

func (i *AccountsRepo) CheckUserWithEmail(ctx echo.Context, email string) (bool, error) {
	user := models.User{}
	result := i.MySQL.DB.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (i *AccountsRepo) GetUserByEmail(ctx echo.Context, email string) (models.User, error) {
	user := models.User{}
	result := i.MySQL.DB.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return models.User{}, errs.ErrInstanceNotFound
		}
		return models.User{}, result.Error
	}
	return user, nil
}

func (i *AccountsRepo) CreateUser(user models.User) (models.User, error) {
	result := i.MySQL.DB.Table("users").Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

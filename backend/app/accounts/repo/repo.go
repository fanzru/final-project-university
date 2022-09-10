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

func (i *AccountsRepo) GetUserByEmail(ctx echo.Context, email string) (models.User, error) {
	var user models.User
	result := i.MySQL.DB.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errs.ErrInstanceNotFound
		}
		return user, result.Error
	}
	if result.RowsAffected < 1 {
		return user, errs.ErrInstanceNotFound
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

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
	CreateUser(ctx echo.Context, user models.User) (models.User, error)
	GetAllPaperByUserId(ctx echo.Context) (*models.Profile, error)
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

func (i *AccountsRepo) CreateUser(ctx echo.Context, user models.User) (models.User, error) {
	result := i.MySQL.DB.Table("users").Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (i *AccountsRepo) GetUserProfile(ctx echo.Context) (*models.User, error) {
	user := &models.User{}
	result := i.MySQL.DB.Table("users").Where("id = ?", ctx.Get("user_id")).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (i *AccountsRepo) GetAllPaperByUserId(ctx echo.Context) (*models.Profile, error) {
	profile := models.Profile{}
	err := i.MySQL.DB.Table("users").First(&profile, "id = ?", ctx.Get("user_id")).Error
	if err != nil {
		return nil, err
	}
	papersUser := []models.PapersUser{}
	err = i.MySQL.DB.Table("papers_users").Find(&papersUser, "user_id = ?", ctx.Get("user_id")).Error
	if err != nil {
		return nil, err
	}

	// for _, paper := range papersUser {
	// 	sentencesLabel := []models.SentencesLabel{}
	// 	err = i.MySQL.DB.Table("sentences_labels").Find(&sentencesLabel, "paper_id = ?", paper.Id).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	papersUser[0].SentencesLabel = &sentencesLabel
	// }

	profile.PapersUsers = &papersUser
	return &profile, nil
}

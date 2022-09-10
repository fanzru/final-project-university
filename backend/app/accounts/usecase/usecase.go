package usecase

import (
	errs "backend/app/accounts/domain/errors"
	"backend/app/accounts/domain/models"
	"backend/app/accounts/domain/request"
	"backend/app/accounts/repo"
	"backend/infrastructure/config"
	"errors"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Impl interface {
	UserRegister(ctx echo.Context, param request.UserRegisterReq) error
}

type AccountsApp struct {
	AccountsRepo repo.AccountsRepo
	Cfg          config.Config
}

func New(accounts AccountsApp) AccountsApp {
	return accounts
}

func (i *AccountsApp) UserRegister(ctx echo.Context, param request.UserRegisterReq) error {
	user, err := i.AccountsRepo.GetUserByEmail(ctx, param.Email)
	if err == nil {
		return errs.ErrEmailUsed
	}
	if !errors.Is(err, errs.ErrInstanceNotFound) {
		return err
	}

	cryptPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), i.Cfg.IntBycrptPassword)
	if err != nil {
		return err
	}
	log.Println("--------------------------------")
	_, err = i.AccountsRepo.CreateUser(models.User{
		ID:        0,
		Name:      param.Name,
		Email:     param.Email,
		Password:  string(cryptPass),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

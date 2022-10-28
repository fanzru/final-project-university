package usecase

import (
	errs "backend/app/accounts/domain/errors"
	"backend/app/accounts/domain/models"
	"backend/app/accounts/domain/request"
	"backend/app/accounts/domain/response"
	"backend/app/accounts/repo"
	"backend/infrastructure/config"
	"backend/pkg/jwt"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Impl interface {
	UserRegister(ctx echo.Context, param request.UserRegisterReq) error
	UserLogin(ctx echo.Context, param request.UserLoginReq) (*response.UserLoginRes, error)
}

type AccountsApp struct {
	AccountsRepo repo.AccountsRepo
	Cfg          config.Config
}

func New(accounts AccountsApp) AccountsApp {
	return accounts
}

func (i AccountsApp) UserRegister(ctx echo.Context, param request.UserRegisterReq) error {
	_, err := i.AccountsRepo.GetUserByEmail(ctx, param.Email)
	if err == nil {
		return errs.ErrEmailUsed
	}
	if !errors.Is(err, errs.ErrInstanceNotFound) {
		return err
	}

	cryptPass, err := bcrypt.GenerateFromPassword([]byte(param.Password), i.Cfg.IntBycrptPassword)
	if err != nil {
		return err
	}

	_, err = i.AccountsRepo.CreateUser(ctx, models.User{
		Id:        0,
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

func (i AccountsApp) UserLogin(ctx echo.Context, param request.UserLoginReq) (*response.UserLoginRes, error) {
	user, err := i.AccountsRepo.GetUserByEmail(ctx, param.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return nil, err
	}

	token, err := jwt.EncodeToken(user.Id, user.Email, i.Cfg.JWTTokenSecret)
	if err != nil {
		return nil, err
	}
	return &response.UserLoginRes{
		AccessToken: token,
	}, nil
}

func (i AccountsApp) UserProfile(ctx echo.Context) (*response.UserProfileRes, error) {
	profile, err := i.AccountsRepo.GetAllPaperByUserId(ctx)
	if err != nil {
		return nil, err
	}
	userProfileRes := &response.UserProfileRes{
		ID:          profile.Id,
		Name:        profile.Name,
		Email:       profile.Email,
		PhotoUrl:    profile.PhotoUrl,
		CreatedAt:   profile.CreatedAt,
		DeletedAt:   profile.DeletedAt,
		PapersUsers: profile.PapersUsers,
	}
	return userProfileRes, nil
}

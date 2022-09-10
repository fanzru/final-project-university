package http

import (
	"backend/app/accounts/domain/request"
	accountsapp "backend/app/accounts/usecase"
	"backend/infrastructure/config"
	"backend/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	AccountsApp accountsapp.AccountsApp
	Cfg         config.Config
}

func (h AccountHandler) Register(ctx echo.Context) error {
	userRegisterReq := &request.UserRegisterReq{}

	err := ctx.Bind(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.AccountsApp.UserRegister(ctx, request.UserRegisterReq{
		Name:     userRegisterReq.Name,
		Email:    userRegisterReq.Email,
		Password: userRegisterReq.Password,
	})
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessCreated(ctx, nil)
}

func (h *AccountHandler) Login(ctx echo.Context) error {
	userLoginReq := &request.UserLoginReq{}

	err := ctx.Bind(userLoginReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userLoginReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	token, err := h.AccountsApp.UserLogin(ctx, request.UserLoginReq{
		Email:    userLoginReq.Email,
		Password: userLoginReq.Password,
	})

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, token)
}

func (a AccountHandler) Profile(ctx echo.Context) error {
	res, err := a.AccountsApp.UserProfile(ctx)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, res)
}

package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseSuccessCreated(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusCreated,
		SuccessCreatedMessage{
			Code:    http.StatusCreated,
			Message: SuccessCreatedName,
			Data:    data,
		})
}

func ResponseSuccessOK(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK,
		SuccessCreatedMessage{
			Code:    http.StatusOK,
			Message: SuccessOK,
			Data:    data,
		})
}

func ResponseErrorUnauthorized(ctx echo.Context) error {
	// todo: add logger before response
	return ctx.JSON(http.StatusUnauthorized, ErrorMessage{
		Code: http.StatusUnauthorized,
		Error: &ErrorFormat{
			ErrorName:        UnauthorizedErrorName,
			ErrorDescription: UnauthorizedErrorDescription,
		},
	})
}

func ResponseErrorBadRequest(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, ErrorMessage{
		Code: http.StatusBadRequest,
		Error: &ErrorFormat{
			ErrorName:        BadRequestErrorName,
			ErrorDescription: err.Error(),
		},
	})
}

func ResponseInternalServerError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, ErrorMessage{
		Code: http.StatusInternalServerError,
		Error: &ErrorFormat{
			ErrorName:        InternalServerName,
			ErrorDescription: err.Error(),
		},
	})
}

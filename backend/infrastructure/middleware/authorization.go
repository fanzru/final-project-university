package middleware

import (
	"backend/infrastructure/config"
	"backend/infrastructure/database"
	"backend/pkg/jwt"
	"backend/pkg/response"
	"strings"

	"github.com/labstack/echo/v4"
)

// Service Authorizer
type MiddlewareAuth struct {
	DB  database.Connection
	Cfg config.Config
}

func NewServiceAuthorizer(db database.Connection, cfg config.Config) MiddlewareAuth {
	return MiddlewareAuth{
		DB:  db,
		Cfg: cfg,
	}
}

func (m MiddlewareAuth) BearerTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := m.getUserIdFromJWT(ctx)
		if userId == 0 {
			return response.ResponseErrorUnauthorized(ctx)
		}
		return next(ctx)
	}
}

func (m MiddlewareAuth) getUserIdFromJWT(ctx echo.Context) int {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0
	}

	headerSplit := strings.Split(authHeader, " ")
	if len(headerSplit) < 1 {
		return 0
	}

	// Get Token
	token := headerSplit[1]

	// Get JWT Claims
	claims, err := jwt.DecodeToken(token, m.Cfg.JWTTokenSecret)
	if err != nil {
		return 0
	}

	// Bind UserID to context
	ctx.Set("user_id", claims.UserId)
	ctx.Set("user_email", claims.Email)
	return 1
}

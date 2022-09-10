package services

import (
	"backend/infrastructure/config"
	"backend/infrastructure/database"
	"backend/infrastructure/middleware"
)

func RegisterMiddleware(db database.Connection, cfg config.Config) middleware.MiddlewareAuth {
	middlewareAuth := middleware.NewServiceAuthorizer(db, cfg)
	return middlewareAuth
}

package routes

import (
	"auth_admin/handler"
	"auth_admin/infrastucture/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, adminHandler *handler.AdminHandler) {
	e.POST("/api/admin/login", adminHandler.Login)
	e.POST("/api/admin/logout", adminHandler.Logout, middleware.InitJWTMiddleware(), middleware.CheckBlacklist)
}

package main

import (
	"auth_admin/config"
	"auth_admin/handler"
	"auth_admin/repository"
	"auth_admin/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.ConnectDB()
	defer db.Close()

	// seeders.Seed(db)

	adminRepo := repository.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)
	handler.NewAdminHandler(e, adminUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}

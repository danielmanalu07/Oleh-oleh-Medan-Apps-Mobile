package main

import (
	"auth_admin/config"
	"auth_admin/domain/admin"
	"auth_admin/handler"
	"auth_admin/repository"
	"auth_admin/routes"
	"auth_admin/service"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&admin.Admin{})

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("admin12345"), bcrypt.DefaultCost)
	db.Create(&admin.Admin{
		Username: "admin",
		Password: string(hashPassword),
	})

	adminRepo := repository.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	e := echo.New()

	routes.InitRoutes(e, adminHandler)

	e.Start(":8080")
}

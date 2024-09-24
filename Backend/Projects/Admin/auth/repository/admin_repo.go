package repository

import (
	"auth_admin/domain/admin"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) admin.AdminRepository {
	return &AdminRepo{DB: db}
}

func (a *AdminRepo) FindByUsername(username string) (admin.Admin, error) {
	var admin admin.Admin
	err := a.DB.Where("username = ?", username).First(&admin)
	return admin, err.Error
}

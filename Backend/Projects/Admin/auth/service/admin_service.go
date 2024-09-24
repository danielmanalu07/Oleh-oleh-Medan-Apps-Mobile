package service

import (
	"auth_admin/domain/admin"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	AdminRepo admin.AdminRepository
}

func NewAdminService(repo admin.AdminRepository) *AdminService {
	return &AdminService{AdminRepo: repo}
}

func (s *AdminService) Login(username, password string) (admin.Admin, error) {
	adm, err := s.AdminRepo.FindByUsername(username)
	if err != nil {
		return adm, errors.New("admin not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(adm.Password), []byte(password))
	if err != nil {
		return adm, errors.New("Password Incorrect")
	}

	return adm, nil
}

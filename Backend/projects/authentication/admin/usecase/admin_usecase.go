package usecase

import (
	"auth_admin/domain"
	"auth_admin/utils"
	"context"
	"errors"
)

type adminUsecase struct {
	repo domain.AdminRepository
}

func (a *adminUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	admin, err := a.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", errors.New("Invalid Credentials")
	}

	token, err := utils.GenerateToken(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func NewAdminUsecase(repo domain.AdminRepository) domain.AdminUsecase {
	return &adminUsecase{repo}
}

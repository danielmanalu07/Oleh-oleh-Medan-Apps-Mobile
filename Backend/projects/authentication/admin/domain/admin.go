package domain

import "context"

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminRepository interface {
	GetByUsername(ctx context.Context, username string) (Admin, error)
}

type AdminUsecase interface {
	Login(ctx context.Context, username, password string) (string, error)
}

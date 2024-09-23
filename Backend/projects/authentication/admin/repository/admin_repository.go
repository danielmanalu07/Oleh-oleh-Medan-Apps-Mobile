package repository

import (
	"auth_admin/domain"
	"context"
	"database/sql"
)

type adminRepository struct {
	db *sql.DB
}

func (a *adminRepository) GetByUsername(ctx context.Context, username string) (domain.Admin, error) {
	query := "SELECT id, username, password FROM admins WHERE username = $1"
	row := a.db.QueryRowContext(ctx, query, username)
	var admin domain.Admin
	err := row.Scan(&admin.ID, &admin.Username, &admin.Password)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func NewAdminRepository(db *sql.DB) domain.AdminRepository {
	return &adminRepository{db}
}

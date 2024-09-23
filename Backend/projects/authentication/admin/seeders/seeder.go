package seeders

import (
	"auth_admin/utils"
	"database/sql"
	"log"
)

func Seed(db *sql.DB) {
	admins := []struct {
		Id       int
		Username string
		Password string
	}{
		{1, "admin", "admin12345"},
	}

	for _, admin := range admins {
		hashpw, err := utils.HashPassword(admin.Password)
		if err != nil {
			log.Fatalf("Error hashing password for user %s: %v", admin.Username, err)
		}

		_, err = db.Exec("INSERT INTO admins (id, username, password) VALUES ($1, $2, $3)", admin.Id, admin.Username, hashpw)
		if err != nil {
			log.Fatalf("Error seeding user: %v", err)
		}

	}

	log.Println("Seeding Completed.")
}

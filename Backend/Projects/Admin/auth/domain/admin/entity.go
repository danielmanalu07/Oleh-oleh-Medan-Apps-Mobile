package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"unqiue type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

type AdminRepository interface {
	FindByUsername(username string) (Admin, error)
}

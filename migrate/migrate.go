package migrate

import (
	"prog/db"
	"prog/features/auth/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&data.User{},
		&data.Authentication{},
	)
}

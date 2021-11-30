package migrate

import (
	"prog/db"
	articles "prog/features/articles/data"
	auth "prog/features/auth/data"
	likes "prog/features/likes/data"
	users "prog/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&users.User{},
		&auth.Authentication{},
		&articles.Article{},
		&articles.Tag{},
		&likes.ArticleLikes{},
	)
}

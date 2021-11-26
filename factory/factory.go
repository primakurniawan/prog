package factory

import (
	"prog/db"

	// auth domain
	authBusiness "prog/features/auth/business"
	authData "prog/features/auth/data"
	authPresentation "prog/features/auth/presentation"

	// user domain
	userBusiness "prog/features/users/business"
	userData "prog/features/users/data"
	userPresentation "prog/features/users/presentation"

	// article domain
	articleBusiness "prog/features/articles/business"
	articleData "prog/features/articles/data"
	articlePresentation "prog/features/articles/presentation"
)

type Presenter struct {
	AuthHandler authPresentation.AuthHandler

	UserHandler userPresentation.UserHandler

	ArticleHandler articlePresentation.ArticleHandler
}

func Init() Presenter {
	// auth layer
	authData := authData.NewMysqlAuthRepository(db.DB)
	authBusiness := authBusiness.NewAuthBusiness(authData)
	authPresentation := authPresentation.NewAuthHandler(authBusiness)

	// users layer
	userData := userData.NewMysqlUserRepository(db.DB)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	articleData := articleData.NewMysqlArticleRepository(db.DB)
	articleBusiness := articleBusiness.NewArticleBusiness(articleData)
	articlePresentation := articlePresentation.NewArticleHandler(articleBusiness)

	return Presenter{
		AuthHandler: *authPresentation,
		UserHandler: *userPresentation,

		ArticleHandler: *articlePresentation,
	}
}

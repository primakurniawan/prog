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

	// likes domain
	articleLikesBusiness "prog/features/likes/business"
	articleLikesData "prog/features/likes/data"
	articleLikesPresentation "prog/features/likes/presentation"
)

type Presenter struct {
	AuthHandler         authPresentation.AuthHandler
	UserHandler         userPresentation.UserHandler
	ArticleHandler      articlePresentation.ArticleHandler
	ArticleLikesHandler articleLikesPresentation.ArticleLikesHandler
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

	// articles layer
	articleData := articleData.NewMysqlArticleRepository(db.DB)
	articleBusiness := articleBusiness.NewArticleBusiness(articleData)
	articlePresentation := articlePresentation.NewArticleHandler(articleBusiness)

	// article likes layer
	articleLikesData := articleLikesData.NewMysqlArticleLikesRepository(db.DB)
	articleLikesBusiness := articleLikesBusiness.NewArticleLikesBusiness(articleLikesData)
	articleLikesPresentation := articleLikesPresentation.NewArticleLikesHandler(articleLikesBusiness)

	return Presenter{
		AuthHandler:         *authPresentation,
		UserHandler:         *userPresentation,
		ArticleHandler:      *articlePresentation,
		ArticleLikesHandler: *articleLikesPresentation,
	}
}

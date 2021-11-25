package factory

import (
	"prog/db"
	"prog/features/articles/business"
	"prog/features/articles/data"
	"prog/features/articles/presentation"
)

type ArticlePresenter struct {
	ArticleHandler presentation.ArticleHandler
}

func InitArticle() ArticlePresenter {
	articleData := data.NewMysqlArticleRepository(db.DB)
	articleBusiness := business.NewArticleBusiness(articleData)
	articlePresentation := presentation.NewArticleHandler(articleBusiness)

	return ArticlePresenter{
		ArticleHandler: *articlePresentation,
	}
}

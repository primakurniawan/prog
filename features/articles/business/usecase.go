package business

import (
	"prog/features/articles"
)

type articleUsecase struct {
	ArticleData articles.Data
}

func NewArticleBusiness(articleData articles.Data) articles.Business {
	return &articleUsecase{ArticleData: articleData}
}

func (uu *articleUsecase) CreateArticle(data articles.Core, userId int) error {
	tags, err := uu.ArticleData.CreateTags(data.Tags)

	if err != nil {
		return err
	}
	data.Tags = tags
	data.UserId = userId

	err = uu.ArticleData.CreateArticle(data, userId, tags)
	if err != nil {
		return err
	}
	return nil
}

func (uu *articleUsecase) GetAllArticles() ([]articles.Core, error) {
	articles, err := uu.ArticleData.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (us *articleUsecase) GetArticleById(articleId int) (articles.Core, error) {
	articleData, err := us.ArticleData.GetArticleById(articleId)

	if err != nil {
		return articles.Core{}, err
	}

	return articleData, nil
}

func (uu *articleUsecase) UpdateArticleById(articleId int, data articles.Core, userId int) error {
	err := uu.ArticleData.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return err
	}
	err = uu.ArticleData.UpdateArticleById(articleId, data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *articleUsecase) DeleteArticleById(articleId, userId int) error {
	err := uu.ArticleData.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return err
	}

	err = uu.ArticleData.DeleteArticleById(articleId)
	if err != nil {
		return err
	}

	return nil
}

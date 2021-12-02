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

func (uu *articleUsecase) CreateTags(tags []articles.TagCore) ([]articles.TagCore, error) {
	tags, err := uu.ArticleData.CreateTags(tags)

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (uu *articleUsecase) CreateArticle(data articles.ArticleCore) error {

	err := uu.ArticleData.CreateArticle(data)
	if err != nil {
		return err
	}
	return nil
}

func (uu *articleUsecase) GetAllArticles() ([]articles.ArticleCore, error) {
	articles, err := uu.ArticleData.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (us *articleUsecase) GetArticleById(articleId int) (articles.ArticleCore, error) {
	articleData, err := us.ArticleData.GetArticleById(articleId)

	if err != nil {
		return articles.ArticleCore{}, err
	}

	return articleData, nil
}

func (uu *articleUsecase) UpdateArticleById(articleId int, data articles.ArticleCore) error {

	err := uu.ArticleData.UpdateArticleById(articleId, data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *articleUsecase) DeleteArticleById(articleId int) error {

	err := uu.ArticleData.DeleteArticleById(articleId)
	if err != nil {
		return err
	}

	return nil
}

func (uu *articleUsecase) GetAllUserArticles(userId int) ([]articles.ArticleCore, error) {
	articles, err := uu.ArticleData.GetAllUserArticles(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (uu *articleUsecase) VerifyArticleOwner(articleId, userId int) error {
	err := uu.ArticleData.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return err
	}

	return nil
}

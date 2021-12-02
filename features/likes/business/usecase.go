package business

import (
	"prog/features/articles"
	"prog/features/likes"
	"prog/features/users"
)

type articleLikesUsecase struct {
	LikeData likes.Data
}

func NewArticleLikesBusiness(articleLikesData likes.Data) likes.Business {
	return &articleLikesUsecase{LikeData: articleLikesData}
}

func (alu *articleLikesUsecase) LikeArticle(articleId, userId int) error {
	err := alu.LikeData.LikeArticle(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) UnlikeArticle(articleId, userId int) error {
	err := alu.LikeData.UnlikeArticle(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) GetLikedArticles(userId int) ([]articles.ArticleCore, error) {
	articles, err := alu.LikeData.GetLikedArticles(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (alu *articleLikesUsecase) GetLikingUsers(articleId int) ([]users.Core, error) {
	users, err := alu.LikeData.GetLikingUsers(articleId)
	if err != nil {
		return nil, err
	}

	return users, nil
}

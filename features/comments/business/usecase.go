package business

import (
	"prog/features/comments"
)

type articleLikesUsecase struct {
	CommentData comments.Data
}

func NewCommentsBusiness(articleLikesData comments.Data) comments.Business {
	return &articleLikesUsecase{CommentData: articleLikesData}
}

func (alu *articleLikesUsecase) AddComment(content string, articleId, userId int) error {
	err := alu.CommentData.AddComment(content, articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) UpdateComment(content string, commentId, userId int) error {
	err := alu.CommentData.VerifyCommentOwner(commentId, userId)
	if err != nil {
		return err
	}
	err = alu.CommentData.UpdateComment(commentId, content)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) DeleteComment(commentId, userId int) error {
	err := alu.CommentData.VerifyCommentOwner(commentId, userId)
	if err != nil {
		return err
	}
	err = alu.CommentData.DeleteComment(commentId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) GetArticleComments(articleId int) ([]comments.Core, error) {
	data, err := alu.CommentData.GetArticleComments(articleId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

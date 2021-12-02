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

func (alu *articleLikesUsecase) AddComment(data comments.Core) error {
	err := alu.CommentData.AddComment(data)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) UpdateComment(commentId int, data comments.Core) error {

	err := alu.CommentData.UpdateComment(commentId, data)
	if err != nil {
		return err
	}
	return nil

}

func (alu *articleLikesUsecase) DeleteComment(commentId int) error {

	err := alu.CommentData.DeleteComment(commentId)
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

func (alu *articleLikesUsecase) VerifyCommentOwner(commentId, userId int) error {

	err := alu.CommentData.VerifyCommentOwner(commentId, userId)
	if err != nil {
		return err
	}
	return nil

}

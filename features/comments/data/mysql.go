package data

import (
	"errors"
	"prog/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentsRepository struct {
	Conn *gorm.DB
}

func NewMysqlCommentsRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentsRepository{
		Conn: conn,
	}
}

func (cr *mysqlCommentsRepository) AddComment(content string, articleId, userId int) error {
	comment := toCommentsRecord(comments.Core{
		Content:   content,
		ArticleID: articleId,
		UserID:    userId,
	})

	err := cr.Conn.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *mysqlCommentsRepository) UpdateComment(commentId int, content string) error {
	comment := Comment{}

	err := cr.Conn.First(&comment, commentId).Error
	if err != nil {
		return err
	}
	comment.Content = content
	err = cr.Conn.Save(&comment).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *mysqlCommentsRepository) DeleteComment(commentId int) error {

	err := cr.Conn.Delete(&Comment{}, commentId).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *mysqlCommentsRepository) GetArticleComments(articleId int) ([]comments.Core, error) {

	commentsArticle := []Comment{}
	err := cr.Conn.Where("article_id = ?", articleId).Find(&commentsArticle).Error
	if err != nil {
		return []comments.Core{}, err
	}
	return toCommentCoreList(commentsArticle), nil

}

func (cr *mysqlCommentsRepository) VerifyCommentOwner(articleId, userId int) error {

	err := cr.Conn.Where("article_id = ? AND user_id", articleId, userId).Find(&Comment{}).Error
	if err != nil {
		return errors.New("it's not your comment")
	}
	return nil

}

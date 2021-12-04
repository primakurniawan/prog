package business

import (
	"errors"
	"os"
	"prog/features/articles"
	"prog/features/comments"
	"prog/features/comments/mocks"
	"prog/features/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	commentsRepo     mocks.Data
	commentsBusiness comments.Business
	commentData      comments.Core
	commentsData     []comments.Core
)

func TestMain(m *testing.M) {
	commentsBusiness = NewCommentsBusiness(&commentsRepo)
	commentData = comments.Core{
		ID:      1,
		Content: "Wah keren banget",
		UserID:  1,
		User: users.Core{
			ID:       1,
			Email:    "prima@gmail.com",
			Password: "admin",
			Image:    "prima_image.jpg",
		},
		ArticleID: 1,
		Article: articles.ArticleCore{
			ID:        1,
			Title:     "si malin kundang",
			Image:     "simalin.jpg",
			Content:   "simalin konten",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    2,
			User: users.Core{
				ID:       2,
				Email:    "chris@gmail.com",
				Password: "admin",
				Image:    "chris_image.jpg",
			},
			Tags: []articles.TagCore{
				{
					ID:    1,
					Title: "fabel",
				},
				{
					ID:    2,
					Title: "dongeng",
				},
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	commentsData = []comments.Core{
		{ID: 1,
			Content: "Wah keren banget",
			UserID:  1,
			User: users.Core{
				ID:       1,
				Email:    "prima@gmail.com",
				Password: "admin",
				Image:    "prima_image.jpg",
			},
			ArticleID: 1,
			Article: articles.ArticleCore{
				ID:        1,
				Title:     "si malin kundang",
				Image:     "simalin.jpg",
				Content:   "simalin konten",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				UserID:    2,
				User: users.Core{
					ID:       2,
					Email:    "chris@gmail.com",
					Password: "admin",
					Image:    "chris_image.jpg",
				},
				Tags: []articles.TagCore{
					{
						ID:    1,
						Title: "fabel",
					},
					{
						ID:    2,
						Title: "dongeng",
					},
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{ID: 2,
			Content: "Wah kok keren banget yah",
			UserID:  1,
			User: users.Core{
				ID:       1,
				Email:    "jo@gmail.com",
				Password: "admin",
				Image:    "jo_image.jpg",
			},
			ArticleID: 1,
			Article: articles.ArticleCore{
				ID:        1,
				Title:     "si malin kundang",
				Image:     "simalin.jpg",
				Content:   "simalin konten",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				UserID:    2,
				User: users.Core{
					ID:       2,
					Email:    "chris@gmail.com",
					Password: "admin",
					Image:    "chris_image.jpg",
				},
				Tags: []articles.TagCore{
					{
						ID:    1,
						Title: "fabel",
					},
					{
						ID:    2,
						Title: "dongeng",
					},
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	os.Exit(m.Run())
}

func TestAddComment(t *testing.T) {

	t.Run("comment article success", func(t *testing.T) {
		commentsRepo.On("AddComment", mock.AnythingOfType("comments.Core")).Return(nil).Once()
		err := commentsBusiness.AddComment(commentData)
		assert.Nil(t, err)
	})

	t.Run("comment article fail", func(t *testing.T) {
		commentsRepo.On("AddComment", mock.AnythingOfType("comments.Core")).Return(errors.New("comment article fail")).Once()
		err := commentsBusiness.AddComment(commentData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "comment article fail")
	})

}
func TestGetArticleComments(t *testing.T) {

	t.Run("get article comments success", func(t *testing.T) {
		commentsRepo.On("GetArticleComments", mock.AnythingOfType("int")).Return(commentsData, nil).Once()
		resp, err := commentsBusiness.GetArticleComments(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("get article comments fail", func(t *testing.T) {
		commentsRepo.On("GetArticleComments", mock.AnythingOfType("int")).Return(nil, errors.New("get article comments fail")).Once()
		resp, err := commentsBusiness.GetArticleComments(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "get article comments fail")
	})

}

func TestDeleteComment(t *testing.T) {

	t.Run("delete comment success", func(t *testing.T) {
		commentsRepo.On("DeleteComment", mock.AnythingOfType("int")).Return(nil).Once()
		err := commentsBusiness.DeleteComment(1)
		assert.Nil(t, err)
	})

	t.Run("delete comment fail", func(t *testing.T) {
		commentsRepo.On("DeleteComment", mock.AnythingOfType("int")).Return(errors.New("delete comment fail")).Once()
		err := commentsBusiness.DeleteComment(0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "delete comment fail")
	})

}
func TestUpdateComment(t *testing.T) {

	t.Run("update comment success", func(t *testing.T) {
		commentsRepo.On("UpdateComment", mock.AnythingOfType("int"), mock.AnythingOfType("comments.Core")).Return(nil).Once()
		err := commentsBusiness.UpdateComment(1, commentData)
		assert.Nil(t, err)
	})

	t.Run("update comment fail", func(t *testing.T) {
		commentsRepo.On("UpdateComment", mock.AnythingOfType("int"), mock.AnythingOfType("comments.Core")).Return(errors.New("update comment fail")).Once()
		err := commentsBusiness.UpdateComment(0, commentData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "update comment fail")
	})

}
func TestVerifyCommentOwner(t *testing.T) {

	t.Run("verify comment owner success", func(t *testing.T) {
		commentsRepo.On("VerifyCommentOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		err := commentsBusiness.VerifyCommentOwner(1, 1)
		assert.Nil(t, err)
	})

	t.Run("verify comment owner fail", func(t *testing.T) {
		commentsRepo.On("VerifyCommentOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("verify comment owner fail")).Once()
		err := commentsBusiness.VerifyCommentOwner(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "verify comment owner fail")
	})

}

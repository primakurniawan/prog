package business

import (
	"errors"
	"os"
	"prog/features/articles"
	"prog/features/likes"
	"prog/features/likes/mocks"
	"prog/features/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	likesRepo     mocks.Data
	likesBusiness likes.Business
	usersData     []users.Core
	articlesData  []articles.ArticleCore
)

func TestMain(m *testing.M) {
	likesBusiness = NewArticleLikesBusiness(&likesRepo)
	usersData = []users.Core{
		{
			ID:       1,
			Email:    "prima@gmail.com",
			Password: "admin",
			Image:    "prima_image.jpg",
		},
		{
			ID:       2,
			Email:    "dwi@gmail.com",
			Password: "admin",
			Image:    "dwi_image.jpg",
		},
	}

	articlesData = []articles.ArticleCore{
		{
			ID:        3,
			Title:     "si malin kundang",
			Image:     "simalin.jpg",
			Content:   "simalin konten",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    3,
			User: users.Core{
				ID:       3,
				Email:    "prima@gmail.com",
				Password: "admin",
				Image:    "prima_image.jpg",
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
		{
			ID:        2,
			Title:     "si mardan kundang",
			Image:     "simardan.jpg",
			Content:   "simardan konten",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    2,
			User: users.Core{
				ID:       2,
				Email:    "dwi@gmail.com",
				Password: "admin",
				Image:    "dwi_image.jpg",
			},
			Tags: []articles.TagCore{
				{
					ID:    3,
					Title: "cerpen",
				},
				{
					ID:    4,
					Title: "sinopsis",
				},
			},
		},
	}

	os.Exit(m.Run())
}

func TestLikeArticle(t *testing.T) {

	t.Run("like article success", func(t *testing.T) {
		likesRepo.On("LikeArticle", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		err := likesBusiness.LikeArticle(1, 1)
		assert.Nil(t, err)
	})

	t.Run("like article fail", func(t *testing.T) {
		likesRepo.On("LikeArticle", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("like article fail")).Once()
		err := likesBusiness.LikeArticle(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "like article fail")
	})

}
func TestUnlikeArticle(t *testing.T) {

	t.Run("like article success", func(t *testing.T) {
		likesRepo.On("UnlikeArticle", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		err := likesBusiness.UnlikeArticle(1, 1)
		assert.Nil(t, err)
	})

	t.Run("like article fail", func(t *testing.T) {
		likesRepo.On("UnlikeArticle", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("unlike article fail")).Once()
		err := likesBusiness.UnlikeArticle(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "unlike article fail")
	})

}

func TestGetLikedArticles(t *testing.T) {

	t.Run("Get liked articles success", func(t *testing.T) {
		likesRepo.On("GetLikedArticles", mock.AnythingOfType("int")).Return(articlesData, nil).Once()
		resp, err := likesBusiness.GetLikedArticles(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get liked articles fail", func(t *testing.T) {
		likesRepo.On("GetLikedArticles", mock.AnythingOfType("int")).Return(nil, errors.New("Get liked articles fail")).Once()
		resp, err := likesBusiness.GetLikedArticles(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Get liked articles fail")
	})

}
func TestGetLikingUsers(t *testing.T) {

	t.Run("Get liked articles success", func(t *testing.T) {
		likesRepo.On("GetLikingUsers", mock.AnythingOfType("int")).Return(usersData, nil).Once()
		resp, err := likesBusiness.GetLikingUsers(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get liking users fail", func(t *testing.T) {
		likesRepo.On("GetLikingUsers", mock.AnythingOfType("int")).Return(nil, errors.New("Get liking users fail")).Once()
		resp, err := likesBusiness.GetLikingUsers(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Get liking users fail")
	})

}

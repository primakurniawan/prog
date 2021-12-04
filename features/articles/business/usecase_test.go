package business

import (
	"errors"
	"os"
	"prog/features/articles"
	"prog/features/articles/mocks"
	"prog/features/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	articlesRepo     mocks.Data
	articlesBusiness articles.Business
	articleData      articles.ArticleCore
	articlesData     []articles.ArticleCore
)

func TestMain(m *testing.M) {
	articlesBusiness = NewArticleBusiness(&articlesRepo)

	articleData = articles.ArticleCore{
		ID:        1,
		Title:     "si bodat kundang",
		Image:     "sibodat.jpg",
		Content:   "sibodat konten",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
		User: users.Core{
			ID:       1,
			Email:    "maria@gmail.com",
			Password: "admin",
			Image:    "maria_image.jpg",
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

func TestCreateTags(t *testing.T) {

	t.Run("create tags success", func(t *testing.T) {
		articlesRepo.On("CreateTags", mock.AnythingOfType("[]articles.TagCore")).Return(articleData.Tags, nil).Once()
		tags, err := articlesBusiness.CreateTags(articleData.Tags)
		assert.Nil(t, err)
		assert.Equal(t, len(tags), 2)
	})

	t.Run("create tags fail", func(t *testing.T) {
		articlesRepo.On("CreateTags", mock.AnythingOfType("[]articles.TagCore")).Return(nil, errors.New("create tags fail")).Once()
		tags, err := articlesBusiness.CreateTags(articleData.Tags)
		assert.NotNil(t, err)
		assert.Nil(t, tags)
		assert.Equal(t, err.Error(), "create tags fail")
	})

}

func TestCreateArticle(t *testing.T) {

	t.Run("create article success", func(t *testing.T) {
		articlesRepo.On("CreateArticle", mock.AnythingOfType("articles.ArticleCore")).Return(nil).Once()
		err := articlesBusiness.CreateArticle(articleData)
		assert.Nil(t, err)
	})

	t.Run("Create article fail", func(t *testing.T) {
		articlesRepo.On("CreateArticle", mock.AnythingOfType("articles.ArticleCore")).Return(errors.New("Create article fail")).Once()
		err := articlesBusiness.CreateArticle(articleData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Create article fail")
	})

}

func TestGetAllArticles(t *testing.T) {

	t.Run("Get all articles success", func(t *testing.T) {
		articlesRepo.On("GetAllArticles").Return(articlesData, nil).Once()
		resp, err := articlesBusiness.GetAllArticles()
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get all articles fail", func(t *testing.T) {
		articlesRepo.On("GetAllArticles").Return(nil, errors.New("Get all articles fail")).Once()
		resp, err := articlesBusiness.GetAllArticles()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Get all articles fail")
		assert.Nil(t, resp)
	})

}

func TestUpdateArticleById(t *testing.T) {

	t.Run("Update article by id success", func(t *testing.T) {
		articlesRepo.On("UpdateArticleById", mock.AnythingOfType("int"), mock.AnythingOfType("articles.ArticleCore")).Return(nil).Once()
		err := articlesBusiness.UpdateArticleById(1, articleData)
		assert.Nil(t, err)
	})

	t.Run("Update article by id fail", func(t *testing.T) {
		articlesRepo.On("UpdateArticleById", mock.AnythingOfType("int"), mock.AnythingOfType("articles.ArticleCore")).Return(errors.New("Update article by id fail")).Once()
		err := articlesBusiness.UpdateArticleById(0, articleData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Update article by id fail")
	})

}
func TestDeleteArticleById(t *testing.T) {

	t.Run("Delete article by id success", func(t *testing.T) {
		articlesRepo.On("DeleteArticleById", mock.AnythingOfType("int")).Return(nil).Once()
		err := articlesBusiness.DeleteArticleById(1)
		assert.Nil(t, err)
	})

	t.Run("Delete article by id fail", func(t *testing.T) {
		articlesRepo.On("DeleteArticleById", mock.AnythingOfType("int")).Return(errors.New("Delete article by id fail")).Once()
		err := articlesBusiness.DeleteArticleById(0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Delete article by id fail")
	})

}
func TestGetArticleById(t *testing.T) {

	t.Run("Get article by id success", func(t *testing.T) {
		articlesRepo.On("GetArticleById", mock.AnythingOfType("int")).Return(articleData, nil).Once()
		resp, err := articlesBusiness.GetArticleById(1)
		assert.Nil(t, err)
		assert.Equal(t, resp.ID, 1)
	})

	t.Run("Get article by id fail", func(t *testing.T) {
		articlesRepo.On("GetArticleById", mock.AnythingOfType("int")).Return(articles.ArticleCore{}, errors.New("Get article by id")).Once()
		resp, err := articlesBusiness.GetArticleById(0)
		assert.NotNil(t, err)
		assert.Empty(t, resp)
		assert.Equal(t, err.Error(), "Get article by id")
	})

}
func TestVerifyArticleOwner(t *testing.T) {

	t.Run("Verify article owner success", func(t *testing.T) {
		articlesRepo.On("VerifyArticleOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		err := articlesBusiness.VerifyArticleOwner(1, 1)
		assert.Nil(t, err)
	})

	t.Run("Verify article owner fail", func(t *testing.T) {
		articlesRepo.On("VerifyArticleOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("Verify article owner")).Once()
		err := articlesBusiness.VerifyArticleOwner(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Verify article owner")
	})

}
func TestGetAllUserArticles(t *testing.T) {

	t.Run("Get all user articles success", func(t *testing.T) {
		articlesRepo.On("GetAllUserArticles", mock.AnythingOfType("int")).Return(articlesData, nil).Once()
		resp, err := articlesBusiness.GetAllUserArticles(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get all user articles fail", func(t *testing.T) {
		articlesRepo.On("GetAllUserArticles", mock.AnythingOfType("int")).Return(nil, errors.New("Get all user articles fail")).Once()
		resp, err := articlesBusiness.GetAllUserArticles(0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Get all user articles fail")
		assert.Nil(t, resp)
	})

}

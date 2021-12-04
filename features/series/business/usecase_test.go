package business

import (
	"errors"
	"os"
	"prog/features/articles"
	"prog/features/series"
	"prog/features/series/mocks"
	"prog/features/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	seriesRepo         mocks.Data
	seriesBusiness     series.Business
	seriesData         series.SeriesCore
	seriesDatas        []series.SeriesCore
	articlesSeriesData series.ArticlesSeriesCore
	articlesData       []articles.ArticleCore
)

func TestMain(m *testing.M) {
	seriesBusiness = NewSeriesBusiness(&seriesRepo)

	seriesData = series.SeriesCore{
		ID:          1,
		Title:       "this is series title",
		Description: "this is series description",
		UserID:      1,
		User: users.Core{
			ID:       1,
			Email:    "prima@gmail.com",
			Password: "admin",
			Image:    "prima_image.jpg",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	seriesDatas = []series.SeriesCore{
		{ID: 1,
			Title:       "this is series title",
			Description: "this is series description",
			UserID:      1,
			User: users.Core{
				ID:       1,
				Email:    "prima@gmail.com",
				Password: "admin",
				Image:    "prima_image.jpg",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{ID: 2,
			Title:       "this is second series title",
			Description: "this is second series description",
			UserID:      2,
			User: users.Core{
				ID:       2,
				Email:    "dwi@gmail.com",
				Password: "admin",
				Image:    "dwi_image.jpg",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	articlesSeriesData = series.ArticlesSeriesCore{
		ArticleId: 1,
		Article: articles.ArticleCore{
			ID:        3,
			Title:     "si malin kundang",
			Image:     "simalin.jpg",
			Content:   "simalin content",
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
					Title: "Babel",
				},
				{
					ID:    2,
					Title: "dongeng",
				},
			},
		},
		SeriesId: 1,
		Series: series.SeriesCore{
			ID:          1,
			Title:       "this is series title",
			Description: "this is series description",
			UserID:      1,
			User: users.Core{
				ID:       1,
				Email:    "prima@gmail.com",
				Password: "admin",
				Image:    "prima_image.jpg",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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

func TestCreateSeries(t *testing.T) {

	t.Run("create series success", func(t *testing.T) {
		seriesRepo.On("CreateSeries", mock.AnythingOfType("series.SeriesCore")).Return(nil).Once()
		err := seriesBusiness.CreateSeries(seriesData)
		assert.Nil(t, err)
	})

	t.Run("create series fail", func(t *testing.T) {
		seriesRepo.On("CreateSeries", mock.AnythingOfType("series.SeriesCore")).Return(errors.New("create series fail")).Once()
		err := seriesBusiness.CreateSeries(seriesData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "create series fail")
	})

}

func TestDeleteSeries(t *testing.T) {

	t.Run("Delete series success", func(t *testing.T) {
		seriesRepo.On("DeleteSeries", mock.AnythingOfType("int")).Return(nil).Once()
		err := seriesBusiness.DeleteSeries(1)
		assert.Nil(t, err)
	})

	t.Run("Delete series fail", func(t *testing.T) {
		seriesRepo.On("DeleteSeries", mock.AnythingOfType("int")).Return(errors.New("Delete series fail")).Once()
		err := seriesBusiness.DeleteSeries(0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Delete series fail")
	})

}

func TestUpdateSeriesById(t *testing.T) {

	t.Run("Update series success", func(t *testing.T) {
		seriesRepo.On("UpdateSeriesById", mock.AnythingOfType("int"), mock.AnythingOfType("series.SeriesCore")).Return(nil).Once()
		err := seriesBusiness.UpdateSeriesById(1, seriesData)
		assert.Nil(t, err)
	})

	t.Run("Update series fail", func(t *testing.T) {
		seriesRepo.On("UpdateSeriesById", mock.AnythingOfType("int"), mock.AnythingOfType("series.SeriesCore")).Return(errors.New("Update series fail")).Once()
		err := seriesBusiness.UpdateSeriesById(0, seriesData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Update series fail")
	})

}

func TestAddArticleSeries(t *testing.T) {

	t.Run("Add article to series success", func(t *testing.T) {
		seriesRepo.On("AddArticleSeries", mock.AnythingOfType("series.ArticlesSeriesCore")).Return(nil).Once()
		err := seriesBusiness.AddArticleSeries(articlesSeriesData)
		assert.Nil(t, err)
	})

	t.Run("Add article to series fail", func(t *testing.T) {
		seriesRepo.On("AddArticleSeries", mock.AnythingOfType("series.ArticlesSeriesCore")).Return(errors.New("Add article to series fail")).Once()
		err := seriesBusiness.AddArticleSeries(articlesSeriesData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Add article to series fail")
	})

}
func TestDeleteArticleSeries(t *testing.T) {

	t.Run("Delete article from series success", func(t *testing.T) {
		seriesRepo.On("DeleteArticleSeries", mock.AnythingOfType("series.ArticlesSeriesCore")).Return(nil).Once()
		err := seriesBusiness.DeleteArticleSeries(articlesSeriesData)
		assert.Nil(t, err)
	})

	t.Run("Delete article from series fail", func(t *testing.T) {
		seriesRepo.On("DeleteArticleSeries", mock.AnythingOfType("series.ArticlesSeriesCore")).Return(errors.New("Delete article from series fail")).Once()
		err := seriesBusiness.DeleteArticleSeries(articlesSeriesData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Delete article from series fail")
	})

}
func TestGetAllSeries(t *testing.T) {

	t.Run("Get all series success", func(t *testing.T) {
		seriesRepo.On("GetAllSeries").Return(seriesDatas, nil).Once()
		resp, err := seriesBusiness.GetAllSeries()
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get all series fail", func(t *testing.T) {
		seriesRepo.On("GetAllSeries").Return(nil, errors.New("Get all series fail")).Once()
		resp, err := seriesBusiness.GetAllSeries()
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Get all series fail")
	})

}
func TestGetSeriesById(t *testing.T) {

	t.Run("Get series by id success", func(t *testing.T) {
		seriesRepo.On("GetSeriesById", mock.AnythingOfType("int")).Return(seriesData, nil).Once()
		resp, err := seriesBusiness.GetSeriesById(1)
		assert.Nil(t, err)
		assert.Equal(t, resp.ID, 1)
	})

	t.Run("Get series by id fail", func(t *testing.T) {
		seriesRepo.On("GetSeriesById", mock.AnythingOfType("int")).Return(series.SeriesCore{}, errors.New("Get series by id fail")).Once()
		resp, err := seriesBusiness.GetSeriesById(0)
		assert.NotNil(t, err)
		assert.Empty(t, resp)
		assert.Equal(t, err.Error(), "Get series by id fail")
	})

}
func TestGetAllArticleSeries(t *testing.T) {

	t.Run("Get all articles from series success", func(t *testing.T) {
		seriesRepo.On("GetAllArticleSeries", mock.AnythingOfType("int")).Return(articlesData, nil).Once()
		resp, err := seriesBusiness.GetAllArticleSeries(1)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("Get all articles from series fail", func(t *testing.T) {
		seriesRepo.On("GetAllArticleSeries", mock.AnythingOfType("int")).Return(nil, errors.New("Get all articles from series fail")).Once()
		resp, err := seriesBusiness.GetAllArticleSeries(0)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Get all articles from series fail")
	})

}
func TestVerifySeriesOwner(t *testing.T) {

	t.Run("Verify series owner success", func(t *testing.T) {
		seriesRepo.On("VerifySeriesOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()
		err := seriesBusiness.VerifySeriesOwner(1, 1)
		assert.Nil(t, err)
	})

	t.Run("Verify series owner fail", func(t *testing.T) {
		seriesRepo.On("VerifySeriesOwner", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("Verify series owner fail")).Once()
		err := seriesBusiness.VerifySeriesOwner(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Verify series owner fail")
	})

}

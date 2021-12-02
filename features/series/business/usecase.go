package business

import (
	"prog/features/articles"
	"prog/features/series"
)

type seriesUsecase struct {
	SeriesData series.Data
}

func NewSeriesBusiness(seriesData series.Data) series.Business {
	return &seriesUsecase{SeriesData: seriesData}
}

func (su *seriesUsecase) CreateSeries(data series.SeriesCore) error {
	err := su.SeriesData.CreateSeries(data)
	if err != nil {
		return err
	}
	return nil
}

func (uu *seriesUsecase) DeleteSeries(seriesId int) error {
	err := uu.SeriesData.DeleteSeries(seriesId)
	if err != nil {
		return err
	}

	return nil
}

func (us *seriesUsecase) UpdateSeriesById(seriesId int, data series.SeriesCore) error {
	err := us.SeriesData.UpdateSeriesById(seriesId, data)

	if err != nil {
		return err
	}

	return nil
}

func (uu *seriesUsecase) AddArticleSeries(data series.ArticlesSeriesCore) error {
	err := uu.SeriesData.AddArticleSeries(data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *seriesUsecase) DeleteArticleSeries(data series.ArticlesSeriesCore) error {
	err := uu.SeriesData.DeleteArticleSeries(data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *seriesUsecase) GetAllSeries() ([]series.SeriesCore, error) {
	series, err := uu.SeriesData.GetAllSeries()
	if err != nil {
		return nil, err
	}

	return series, nil
}

func (uu *seriesUsecase) GetSeriesById(seriesId int) (series.SeriesCore, error) {
	seriesData, err := uu.SeriesData.GetSeriesById(seriesId)
	if err != nil {
		return series.SeriesCore{}, err
	}

	return seriesData, nil
}

func (uu *seriesUsecase) GetAllArticleSeries(seriesId int) ([]articles.ArticleCore, error) {
	articles, err := uu.SeriesData.GetAllArticleSeries(seriesId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (uu *seriesUsecase) VerifySeriesOwner(seriesId, userId int) error {
	err := uu.SeriesData.VerifySeriesOwner(seriesId, userId)
	if err != nil {
		return err
	}

	return nil
}

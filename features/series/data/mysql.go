package data

import (
	"prog/features/articles"
	"prog/features/series"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlSeriesRepository struct {
	Conn *gorm.DB
}

func NewMysqlSeriesRepository(conn *gorm.DB) series.Data {
	return &mysqlSeriesRepository{
		Conn: conn,
	}
}

func (ur *mysqlSeriesRepository) CreateSeries(data series.SeriesCore) error {

	recordData := toSeriesRecord(data)
	err := ur.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *mysqlSeriesRepository) DeleteSeries(seriesId int) error {

	err := ur.Conn.Delete(&Series{}, seriesId).Error
	if err != nil {
		return err
	}

	return nil

}

func (ur *mysqlSeriesRepository) UpdateSeriesById(seriesId int, data series.SeriesCore) error {
	var series Series
	err := ur.Conn.First(&series, seriesId).Error
	if err != nil {
		return err
	}

	if data.Title != "" {
		series.Title = data.Title
	}
	if data.Description != "" {
		series.Description = data.Description
	}

	err = ur.Conn.Save(series).Error
	if err != nil {
		return err
	}

	return nil

}

func (ur *mysqlSeriesRepository) GetAllSeries() ([]series.SeriesCore, error) {
	var series []Series

	err := ur.Conn.Joins("User").Find(&series).Error
	if err != nil {
		return nil, err
	}

	return toSeriesCoreList(series), nil

}

func (ur *mysqlSeriesRepository) GetSeriesById(seriesId int) (series.SeriesCore, error) {
	var seriesData Series

	err := ur.Conn.Joins("User").First(&seriesData, seriesId).Error
	if err != nil {
		return series.SeriesCore{}, err
	}

	return toSeriesCore(seriesData), nil
}

func (ur *mysqlSeriesRepository) AddArticleSeries(data series.ArticlesSeriesCore) error {

	articlesSeries := toArticlesSeriesRecord(data)

	err := ur.Conn.Create(&articlesSeries).Error

	if err != nil {
		return err
	}

	return nil

}

func (ur *mysqlSeriesRepository) GetAllArticleSeries(seriesId int) ([]articles.Core, error) {

	var articlesSeries []ArticleSeries

	err := ur.Conn.Preload(clause.Associations).Joins("Article").Where("series_id = ?", seriesId).Find(&articlesSeries).Error

	if err != nil {
		return nil, err
	}

	return ToArticleCoreList(articlesSeries), nil

}

func (ur *mysqlSeriesRepository) VerifySeriesOwner(seriesId, userId int) error {

	err := ur.Conn.Where("id = ? AND user_id = ?", seriesId, userId).First(&Series{}).Error
	if err != nil {
		return err
	}

	return nil

}

package data

import (
	"prog/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	Conn *gorm.DB
}

func NewMysqlArticleRepository(conn *gorm.DB) articles.Data {
	return &mysqlArticleRepository{
		Conn: conn,
	}
}

func (ar *mysqlArticleRepository) CreateTags(tags []articles.TagCore) ([]articles.TagCore, error) {
	tagsTitle := make([]string, 0, len(tags))
	articleTags := []Tag{}
	for _, tag := range tags {
		err := ar.Conn.Where("title = ?", tag.Title).First(&Tag{}).Error
		if err != nil {
			err := ar.Conn.Create(&Tag{Title: tag.Title}).Error
			if err != nil {
				return []articles.TagCore{}, err
			}
		}
		tagsTitle = append(tagsTitle, tag.Title)
	}
	err := ar.Conn.Where("title IN ?", tagsTitle).Find(&articleTags).Error
	if err != nil {
		return []articles.TagCore{}, err
	}
	return toTagsCoreList(articleTags), nil

}

func (ar *mysqlArticleRepository) CreateArticle(data articles.Core, userId int, tags []articles.TagCore) error {

	recordData := toArticleRecord(data)
	err := ar.Conn.Create(&recordData)
	if err != nil {
		return err.Error
	}
	return nil
}

func (ar *mysqlArticleRepository) GetAllArticles() ([]articles.Core, error) {

	articles := []Article{}
	err := ar.Conn.Find(&articles).Error
	if err != nil {
		return toArticleCoreList([]Article{}), err
	}
	return toArticleCoreList(articles), nil
}

func (ar *mysqlArticleRepository) GetArticleById(articleId int) (articles.Core, error) {

	article := Article{}
	err := ar.Conn.First(&article, articleId).Error
	if err != nil {
		return toArticleCore(Article{}), err
	}
	return toArticleCore(article), nil
}

func (ar *mysqlArticleRepository) UpdateArticleById(articleId int, data articles.Core) error {

	article := toArticleRecord(data)
	article.ID = articleId
	err := ar.Conn.Save(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) DeleteArticleById(articleId int) error {

	err := ar.Conn.Delete(&Article{}, articleId).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) VerifyArticleOwner(articleId int, userId int) error {

	err := ar.Conn.Where("id = ? AND userId >= ?", articleId, userId).First(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

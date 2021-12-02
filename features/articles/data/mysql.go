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
	return ToTagsCoreList(articleTags), nil

}

func (ar *mysqlArticleRepository) CreateArticle(data articles.ArticleCore) error {

	recordData := ToArticleRecord(data)
	err := ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) GetAllArticles() ([]articles.ArticleCore, error) {

	articles := []Article{}
	err := ar.Conn.Joins("User").Preload("Tags").Find(&articles).Error
	if err != nil {
		return ToArticleCoreList([]Article{}), err
	}
	return ToArticleCoreList(articles), nil
}

func (ar *mysqlArticleRepository) GetArticleById(articleId int) (articles.ArticleCore, error) {

	article := Article{}
	err := ar.Conn.Joins("User").Preload("Tags").First(&article, articleId).Error
	if err != nil {
		return ToArticleCore(Article{}), err
	}
	return ToArticleCore(article), nil
}

func (ar *mysqlArticleRepository) UpdateArticleById(articleId int, data articles.ArticleCore) error {

	article := ToArticleRecord(data)
	article.ID = articleId

	if data.Title != "" {
		article.Title = data.Title
	}
	if data.Image != "" {
		article.Image = data.Image
	}
	if data.Content != "" {
		article.Content = data.Content
	}

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

	err := ar.Conn.Where("id = ? AND user_id = ?", articleId, userId).First(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) GetAllUserArticles(userId int) ([]articles.ArticleCore, error) {

	articles := []Article{}
	err := ar.Conn.Joins("User").Preload("Tags").Where("user_id = ?", userId).Find(&articles).Error
	if err != nil {
		return ToArticleCoreList([]Article{}), err
	}
	return ToArticleCoreList(articles), nil
}

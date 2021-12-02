package request

import (
	"prog/features/articles"
)

type ArticleRequest struct {
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (requestData *ArticleRequest) ToTagCoreList() []articles.TagCore {
	convertedData := make([]articles.TagCore, 0, len(requestData.Tags))
	for _, v := range requestData.Tags {
		convertedData = append(convertedData, articles.TagCore{Title: v})
	}
	return convertedData
}

func (requestData *ArticleRequest) ToArticleCore(tags []articles.TagCore, userId int) articles.ArticleCore {
	return articles.ArticleCore{
		Title:   requestData.Title,
		Image:   requestData.Image,
		Content: requestData.Content,
		UserID:  userId,
		Tags:    tags,
	}
}

func (requestData *ArticleRequest) ToTagCore() articles.TagCore {
	return articles.TagCore{
		Title: requestData.Title,
	}
}

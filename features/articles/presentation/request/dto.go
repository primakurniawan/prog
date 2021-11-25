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

func (requestData *ArticleRequest) ToArticleCore() articles.Core {
	return articles.Core{
		Title:   requestData.Title,
		Image:   requestData.Image,
		Content: requestData.Content,
	}
}

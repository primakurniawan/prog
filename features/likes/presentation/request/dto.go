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

func toTagCoreList(requestTagsData []string) []articles.TagCore {
	convertedData := make([]articles.TagCore, 0, len(requestTagsData))
	for _, v := range requestTagsData {
		convertedData = append(convertedData, articles.TagCore{Title: v})
	}
	return convertedData
}

func (requestData *ArticleRequest) ToArticleCore() articles.ArticleCore {
	return articles.ArticleCore{
		Title:   requestData.Title,
		Image:   requestData.Image,
		Content: requestData.Content,
		Tags:    toTagCoreList(requestData.Tags),
	}
}

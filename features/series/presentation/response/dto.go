package response

import (
	"prog/features/articles"
	"prog/features/series"
	"prog/features/users"
	"time"
)

type SeriesResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

type ArticleResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Tags      []string     `json:"tags"`
	User      UserResponse `json:"user"`
}

func ToSeriesResponse(series series.SeriesCore) SeriesResponse {
	return SeriesResponse{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		User:        ToUserResponse(series.User),
	}
}

func ToUserResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func ToArticleResponse(article articles.Core) ArticleResponse {
	return ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      ToUserResponse(article.User),
		Tags:      ToTagsResponse(article.Tags),
	}
}

func ToTagsResponse(tags []articles.TagCore) []string {
	convertedTags := make([]string, 0, len(tags))
	for _, v := range tags {
		convertedTags = append(convertedTags, v.Title)
	}
	return convertedTags
}

func ToSeriesResponseList(seriesList []series.SeriesCore) []SeriesResponse {
	convertedSeries := []SeriesResponse{}
	for _, series := range seriesList {
		convertedSeries = append(convertedSeries, ToSeriesResponse(series))
	}

	return convertedSeries
}

func ToArticleResponseList(articleList []articles.Core) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}

package response

import (
	"prog/features/articles"
	"prog/features/users"
	"time"
)

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

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func ToArticleResponse(article articles.ArticleCore) ArticleResponse {
	return ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserResponse(article.User),
		Tags:      toTagsResponse(article.Tags),
	}
}

func toUserResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toTagsResponse(tags []articles.TagCore) []string {
	convertedTags := make([]string, 0, len(tags))
	for _, v := range tags {
		convertedTags = append(convertedTags, v.Title)
	}
	return convertedTags
}

func ToArticleResponseList(articleList []articles.ArticleCore) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}

func ToUserResponseList(userList []users.Core) []UserResponse {
	convertedArticle := []UserResponse{}
	for _, user := range userList {
		convertedArticle = append(convertedArticle, toUserResponse(user))
	}

	return convertedArticle
}

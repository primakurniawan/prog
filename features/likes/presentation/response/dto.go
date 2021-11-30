package response

import (
	"prog/features/likes"
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

func ToArticleResponse(article likes.ArticleCore) ArticleResponse {
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

func toUserResponse(article likes.UserCore) UserResponse {
	return UserResponse{
		ID:       article.ID,
		Email:    article.Email,
		Fullname: article.Fullname,
		Image:    article.Image,
	}
}

func toTagsResponse(tags []likes.TagCore) []string {
	convertedTags := make([]string, 0, len(tags))
	for _, v := range tags {
		convertedTags = append(convertedTags, v.Title)
	}
	return convertedTags
}

func ToArticleResponseList(articleList []likes.ArticleCore) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}

func ToUserResponseList(userList []likes.UserCore) []UserResponse {
	convertedArticle := []UserResponse{}
	for _, user := range userList {
		convertedArticle = append(convertedArticle, toUserResponse(user))
	}

	return convertedArticle
}

package response

import (
	"prog/features/articles"
	"time"
)

type ArticleResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func ToArticleResponse(article articles.Core) ArticleResponse {
	return ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserResponse(article.User),
	}
}

func toUserResponse(article articles.UserCore) UserResponse {
	return UserResponse{
		ID:       article.ID,
		Email:    article.Email,
		Fullname: article.Fullname,
		Image:    article.Image,
	}
}

func ToArticleResponseList(articleList []articles.Core) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}

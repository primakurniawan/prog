package data

import (
	"prog/features/news"
	"time"
)

type Data struct {
	Articles []Article `json: "articles"`
}

type Article struct {
	Source      string    `json: "source"`
	Author      string    `json: "author"`
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Url         string    `json: "url"`
	UrlToImage  string    `json: "url_to_image"`
	PublishedAt time.Time `json: "published_at"`
	Content     string    `json: "content"`
}

func (n *Article) ToCore() news.Core {
	return news.Core{
		Source:      n.Source,
		Author:      n.Author,
		Title:       n.Title,
		Description: n.Description,
		Url:         n.Url,
		UrlToImage:  n.UrlToImage,
		PublishedAt: n.PublishedAt,
		Content:     n.Content,
	}
}

func ToCoreList(data []Article) []news.Core {
	converted := []news.Core{}
	for _, new := range data {
		converted = append(converted, new.ToCore())
	}
	return converted
}

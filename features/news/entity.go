package news

import "time"

type Core struct {
	Source      string
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt time.Time
	Content     string
}

type Business interface {
	GetNews(q string) ([]Core, error)
}

type Data interface {
	GetData(q string) ([]Core, error)
}

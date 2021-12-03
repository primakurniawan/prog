package business

import "prog/features/news"

type ApiService struct {
	newsData news.Data
}

func NewApiService(nr news.Data) news.Business {
	return &ApiService{nr}
}

func (as *ApiService) GetNews(q string) ([]news.Core, error) {
	data, err := as.newsData.GetData(q)
	if err != nil {
		return nil, err
	}

	return data, nil
}

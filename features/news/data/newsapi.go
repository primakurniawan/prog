package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prog/features/news"
)

type NewsApi struct {
	URL    string
	Client http.Client
	ApiKey string
}

func NewNewsApiRepository(url string, apiKey string) news.Data {
	return &NewsApi{
		URL:    url,
		Client: http.Client{},
		ApiKey: apiKey,
	}
}

func (nr *NewsApi) GetData(q string) ([]news.Core, error) {
	url := fmt.Sprintf("%v?apiKey=%v&q=%v", nr.URL, nr.ApiKey, q)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var newsData Data

	response, err := nr.Client.Do(request)
	if err != nil {
		return nil, err
	}
	fmt.Print(response.Body)

	err = json.NewDecoder(response.Body).Decode(&newsData)
	if err != nil {
		return nil, err
	}

	return ToCoreList(newsData.Articles), nil
}

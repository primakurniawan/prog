package response

import (
	"prog/features/series"
	userResponse "prog/features/users/presentation/response"
)

type SeriesResponse struct {
	ID          int                       `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	User        userResponse.UserResponse `json:"user"`
}

func ToSeriesResponse(series series.SeriesCore) SeriesResponse {
	return SeriesResponse{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		User:        userResponse.ToUserResponse(series.User),
	}
}

func ToSeriesResponseList(seriesList []series.SeriesCore) []SeriesResponse {
	convertedSeries := []SeriesResponse{}
	for _, series := range seriesList {
		convertedSeries = append(convertedSeries, ToSeriesResponse(series))
	}

	return convertedSeries
}

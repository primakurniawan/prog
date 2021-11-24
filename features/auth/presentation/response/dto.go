package response

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

func ToAuthResponse(accessToken, refreshToken string) AuthResponse {
	return AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func ToAuthRefreshResponse(accessToken string) AuthRefreshResponse {
	return AuthRefreshResponse{
		AccessToken: accessToken,
	}
}

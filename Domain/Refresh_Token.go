package domain

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

package models

type AuthInfoResponse struct {
	LoginUrl  string `json:"login_url"`
	LogoutUrl string `json:"logout_url"`
}

type AuthResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	TokenType             string `json:"token_type"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

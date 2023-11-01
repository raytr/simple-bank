package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

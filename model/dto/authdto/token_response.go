package authdto

type TokenResponse struct {
	AccessToken string `json:"-"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

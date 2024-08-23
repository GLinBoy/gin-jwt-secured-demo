package response

type TokenResponse struct {
	Token      string `json:"token"`
	Expiration int64  `json:"exp"`
}

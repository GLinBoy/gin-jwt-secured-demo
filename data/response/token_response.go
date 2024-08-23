package response

type TokenResponse struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

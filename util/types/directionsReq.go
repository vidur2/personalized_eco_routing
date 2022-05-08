package types

type DirectionsReq struct {
	User       string `json:"user"`
	Start      string `json:"start"`
	End        string `json:"end"`
	OauthToken string `json:"oauth_token"`
}

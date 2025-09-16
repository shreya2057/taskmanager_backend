package views

type LoginDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenDetails struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}

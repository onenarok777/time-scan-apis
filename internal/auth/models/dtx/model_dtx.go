package dtx

type RequestLogin struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

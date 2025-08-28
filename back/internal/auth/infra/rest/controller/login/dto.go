package login

type Input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Output struct {
	Token string `json:"token"`
}

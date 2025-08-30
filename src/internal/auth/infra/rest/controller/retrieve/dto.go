package retrieve

type Input struct {
	ID string `json:"id"`
}

type Output struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

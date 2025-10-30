package update

type Input struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Breed       string `json:"breed"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
}

type Output struct {
	ID string
}

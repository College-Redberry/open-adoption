package register

type Input struct {
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

type Output struct {
	ID string
}

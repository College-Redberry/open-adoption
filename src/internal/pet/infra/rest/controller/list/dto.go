package list

type Output struct {
	Data  []Pet `json:"data"`
	Count int   `json:"count"`
}

type Pet struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Breed    string   `json:"breed"`
	Age      string   `json:"age"`
	Gender   string   `json:"gender"`
	IsAdoped bool     `json:"is_adoped"`
	Images   []string `json:"images"`
}

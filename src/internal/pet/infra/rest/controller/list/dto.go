package list

type Input struct {
	Name      string `query:"name"`
	Breed     string `query:"breed"`
	Age       string `query:"age"`
	Gender    string `query:"gender"`
	IsAdopted *bool  `query:"is_adopted"`
	Offset    int    `query:"offset"`
	Limit     int    `query:"limit"`
}

type Output struct {
	Data  []Pet `json:"data"`
	Count int64 `json:"count"`
}

type Pet struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Breed     string   `json:"breed"`
	Age       string   `json:"age"`
	Gender    string   `json:"gender"`
	IsAdopted bool     `json:"is_adopted"`
	Images    []string `json:"images"`
}

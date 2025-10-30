package update

type Input struct {
	Name      string
	Breed     string
	Age       string
	Gender    string
	IsAdopted *bool
	Offset    int
	Limit     int
}

type Output struct {
	Data  []Pet
	Count int64
}

type Pet struct {
	ID          string
	Name        string
	Description string
	Breed       string
	Age         string
	Gender      string
	IsAdoped    bool
	Images      []string
}

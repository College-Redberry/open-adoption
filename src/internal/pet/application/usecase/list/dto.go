package update

type Input any

type Output []Pet

type Pet struct {
	ID       string
	Name     string
	Breed    string
	Age      string
	Gender   string
	IsAdoped bool
	Images   []string
}

package pet

import "github.com/google/uuid"

type PetProps struct {
	Name        string
	Description string
	Breed       string
	Age         string
	Gender      string
	IsAdoped    bool
}

type Pet struct {
	ID string
	PetProps
}

func New(pet PetProps) Pet {
	return Pet{
		ID:       uuid.New().String(),
		PetProps: pet,
	}
}

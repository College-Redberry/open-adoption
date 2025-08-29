package request

import (
	"time"

	"github.com/google/uuid"
)

type Props struct {
	PetID                             string
	Name                              string
	Email                             string
	Phone                             string
	ApprovedAt                        time.Time
	Age                               int
	HouseHoldAgreed                   bool
	AlreadyPets                       int
	AlreadyPetsCastratedAndVaccinated bool
	Property                          string
	OwnProperty                       bool
	Address                           string
	Income                            int
	SuitableLocation                  string
	AccessToTheStreet                 *bool
	CreatedAt                         time.Time
}

type Request struct {
	ID string
	Props
}

func New(pet Props) Request {
	return Request{
		ID:    uuid.New().String(),
		Props: pet,
	}
}

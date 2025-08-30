package list

import "time"

type Input any

type Output []Request

type Request struct {
	ID                                string
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

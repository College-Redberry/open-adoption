package list

import "time"

type Output []Request

type Request struct {
	ID                                string    `json:"id"`
	PetID                             string    `json:"pet_id"`
	Name                              string    `json:"name"`
	Email                             string    `json:"email"`
	Phone                             string    `json:"phone"`
	ApprovedAt                        time.Time `json:"approved_at"`
	Age                               int       `json:"age"`
	HouseHoldAgreed                   bool      `json:"house_hold_agreed"`
	AlreadyPets                       int       `json:"already_pets"`
	AlreadyPetsCastratedAndVaccinated bool      `json:"already_pets_castrated_and_vaccinated"`
	Property                          string    `json:"property"`
	OwnProperty                       bool      `json:"own_property"`
	Address                           string    `json:"address"`
	Income                            int       `json:"income"`
	SuitableLocation                  string    `json:"suitable_location"`
	AccessToTheStreet                 *bool     `json:"access_to_the_street"`
	CreatedAt                         time.Time `json:"created_at"`
}

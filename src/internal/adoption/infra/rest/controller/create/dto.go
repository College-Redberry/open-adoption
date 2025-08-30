package create

type Input struct {
	PetID                             string `json:"pet_id"`
	Name                              string `json:"name"`
	Email                             string `json:"email"`
	Phone                             string `json:"phone"`
	Age                               int    `json:"age"`
	HouseHoldAgreed                   bool   `json:"house_hold_agreed"`
	AlreadyPets                       int    `json:"already_pets"`
	AlreadyPetsCastratedAndVaccinated bool   `json:"already_pets_castrated_and_vaccinated"`
	Property                          string `json:"property"`
	OwnProperty                       bool   `json:"own_property"`
	Address                           string `json:"address"`
	Income                            int    `json:"income"`
	SuitableLocation                  string `json:"suitable_location"`
	AccessToTheStreet                 *bool  `json:"access_to_the_street"`
}

type Output struct {
	ID string
}

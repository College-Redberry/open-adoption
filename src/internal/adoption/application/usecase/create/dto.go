package create

type Input struct {
	PetID                             string
	Name                              string
	Email                             string
	Phone                             string
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
}

type Output struct {
	ID string
}

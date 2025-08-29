package create

import (
	"encoding/json"
	"log"

	"github.com/college-redberry/open-adoption/internal/adoption/application/service"
	adoption "github.com/college-redberry/open-adoption/internal/adoption/domain/request"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/constants"
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
)

type Create struct {
	petRepo      pet.PetRepo
	requestRepo  adoption.Repo
	emailService service.EmailService
}

func New(petRepo pet.PetRepo, requestRepo adoption.Repo, emailService service.EmailService) *Create {
	return &Create{
		petRepo:      petRepo,
		requestRepo:  requestRepo,
		emailService: emailService,
	}
}

func (usecase *Create) Execute(input Input) (Output, error) {
	pet, err := usecase.petRepo.GetByID(input.PetID)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to get pet", err)
	}

	if pet.ID == "" {
		return Output{}, errs.ErrInvalidData.Wrap("pet does not exist", nil)
	}

	if pet.IsAdoped {
		return Output{}, errs.ErrInternal.Wrap("pet already adopted", nil)
	}

	newRequest := adoption.New(
		adoption.Props{
			PetID:                             input.PetID,
			Name:                              input.Name,
			Email:                             input.Email,
			Phone:                             input.Phone,
			Age:                               input.Age,
			HouseHoldAgreed:                   input.HouseHoldAgreed,
			AlreadyPets:                       input.AlreadyPets,
			AlreadyPetsCastratedAndVaccinated: input.AlreadyPetsCastratedAndVaccinated,
			Property:                          input.Property,
			OwnProperty:                       input.OwnProperty,
			Address:                           input.Address,
			Income:                            input.Income,
			SuitableLocation:                  input.SuitableLocation,
			AccessToTheStreet:                 input.AccessToTheStreet,
		},
	)

	err = usecase.requestRepo.Create(newRequest)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to create request", err)
	}

	go usecase.sendWarning(newRequest)

	return Output{
		ID: newRequest.ID,
	}, nil
}

func (usecase *Create) sendWarning(message any) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	err = usecase.emailService.Send(constants.EMAIL_TO, "New adoption request", string(jsonData))
	if err != nil {
		log.Println(err)
	}
}

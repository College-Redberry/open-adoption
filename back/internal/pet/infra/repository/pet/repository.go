package pet

import (
	"context"
	"database/sql"
	"errors"

	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
	querier "github.com/college-redberry/open-adoption/internal/pet/infra/persist/pet"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PetRepo struct {
	querier *querier.Queries
}

func New(db querier.DBTX) *PetRepo {
	return &PetRepo{
		querier: querier.New(db),
	}
}

func (repo *PetRepo) Create(p pet.Pet) error {
	err := repo.querier.CreatePet(context.Background(), querier.CreatePetParams{
		ID:        uuid.MustParse(p.ID),
		Name:      p.Name,
		Breed:     p.Breed,
		Age:       p.Age,
		Gender:    querier.PetGender(p.Gender), // sqlc enum mapping
		IsAdopted: p.IsAdoped,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *PetRepo) Update(id string, props pet.PetProps) error {
	err := repo.querier.UpdatePet(context.Background(), querier.UpdatePetParams{
		ID:     uuid.MustParse(id),
		Name:   props.Name,
		Breed:  props.Breed,
		Age:    props.Age,
		Gender: querier.PetGender(props.Gender),
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetRepo) AdoptById(id string) error {
	return repo.querier.AdoptPetById(context.Background(), uuid.MustParse(id))
}

func (repo *PetRepo) GetByID(id string) (pet.Pet, error) {
	result, err := repo.querier.GetPetById(context.Background(), uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pet.Pet{}, nil
		}
		return pet.Pet{}, err
	}

	return pet.Pet{
		ID: result.ID.String(),
		PetProps: pet.PetProps{
			Name:     result.Name,
			Breed:    result.Breed,
			Age:      result.Age,
			Gender:   string(result.Gender),
			IsAdoped: result.IsAdopted,
		},
	}, nil
}

func (repo *PetRepo) List() ([]pet.Pet, error) {
	results, err := repo.querier.ListPets(context.Background())
	if err != nil {
		return nil, err
	}

	var pets []pet.Pet
	for _, r := range results {
		pets = append(pets, pet.Pet{
			ID: r.ID.String(),
			PetProps: pet.PetProps{
				Name:     r.Name,
				Breed:    r.Breed,
				Age:      r.Age,
				Gender:   string(r.Gender),
				IsAdoped: r.IsAdopted,
			},
		})
	}

	return pets, nil
}

func (repo *PetRepo) ListImagesBtId(id string) ([]string, error) {
	results, err := repo.querier.ListImagesById(context.Background(), pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true})
	if err != nil {
		return nil, err
	}

	return results, nil
}

package request

import (
	"context"
	"database/sql"
	"errors"

	adoption "github.com/college-redberry/open-adoption/internal/adoption/domain/request"
	querier "github.com/college-redberry/open-adoption/internal/adoption/infra/persist/request"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RequestRepo struct {
	querier *querier.Queries
}

func New(db querier.DBTX) *RequestRepo {
	return &RequestRepo{
		querier: querier.New(db),
	}
}

func (repo *RequestRepo) Create(req adoption.Request) error {
	accessToTheStreet := pgtype.Bool{Bool: false, Valid: false}
	if req.AccessToTheStreet != nil {
		accessToTheStreet = pgtype.Bool{Bool: *req.AccessToTheStreet, Valid: true}
	}

	return repo.querier.CreateRequest(context.Background(), querier.CreateRequestParams{
		ID:                                uuid.MustParse(req.ID),
		PetID:                             uuid.MustParse(req.PetID),
		Name:                              req.Name,
		Email:                             req.Email,
		Phone:                             req.Phone,
		Age:                               int32(req.Age),
		HouseHoldAgreed:                   req.HouseHoldAgreed,
		Alreadypets:                       int32(req.AlreadyPets),
		AlreadyPetsCastratedAndVaccinated: req.AlreadyPetsCastratedAndVaccinated,
		Property:                          querier.PropertyType(req.Property), // enum mapping
		OwnProperty:                       req.OwnProperty,
		Address:                           req.Address,
		Income:                            int32(req.Income),
		SuitableLocation:                  req.SuitableLocation,
		AccessToTheStreet:                 accessToTheStreet,
	})
}

func (repo *RequestRepo) Aprove(id string) error {
	return repo.querier.ApproveRequest(context.Background(), uuid.MustParse(id))
}

func (repo *RequestRepo) List() ([]adoption.Request, error) {
	rows, err := repo.querier.ListRequests(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []adoption.Request{}, nil
		}
		return nil, err
	}

	var requests []adoption.Request
	for _, r := range rows {
		requests = append(requests, adoption.Request{
			ID: r.ID.String(),
			Props: adoption.Props{
				PetID:                             r.PetID.String(),
				Name:                              r.Name,
				Email:                             r.Email,
				Phone:                             r.Phone,
				ApprovedAt:                        r.ApprovedAt.Time,
				Age:                               int(r.Age),
				HouseHoldAgreed:                   r.HouseHoldAgreed,
				AlreadyPets:                       int(r.Alreadypets),
				AlreadyPetsCastratedAndVaccinated: r.AlreadyPetsCastratedAndVaccinated,
				Property:                          string(r.Property),
				OwnProperty:                       r.OwnProperty,
				Address:                           r.Address,
				Income:                            int(r.Income),
				SuitableLocation:                  r.SuitableLocation,
				AccessToTheStreet:                 &r.AccessToTheStreet.Bool,
				CreatedAt:                         r.CreatedAt.Time,
			},
		})
	}

	return requests, nil
}

func (repo *RequestRepo) GetByID(id string) (adoption.Request, error) {
	r, err := repo.querier.GetRequestById(context.Background(), uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return adoption.Request{}, nil
		}
		return adoption.Request{}, err
	}

	return adoption.Request{
		ID: r.ID.String(),
		Props: adoption.Props{
			PetID:                             r.PetID.String(),
			Name:                              r.Name,
			Email:                             r.Email,
			Phone:                             r.Phone,
			ApprovedAt:                        r.ApprovedAt.Time,
			Age:                               int(r.Age),
			HouseHoldAgreed:                   r.HouseHoldAgreed,
			AlreadyPets:                       int(r.Alreadypets),
			AlreadyPetsCastratedAndVaccinated: r.AlreadyPetsCastratedAndVaccinated,
			Property:                          string(r.Property),
			OwnProperty:                       r.OwnProperty,
			Address:                           r.Address,
			Income:                            int(r.Income),
			SuitableLocation:                  r.SuitableLocation,
			AccessToTheStreet:                 &r.AccessToTheStreet.Bool,
			CreatedAt:                         r.CreatedAt.Time,
		},
	}, nil
}

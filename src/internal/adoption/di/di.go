package di

import (
	aproveUC "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/aprove"
	createUC "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/create"
	listUC "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/list"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/persist"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/repository/request"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/rest/controller/aprove"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/rest/controller/create"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/rest/controller/list"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/service/email"
	"github.com/college-redberry/open-adoption/internal/pet/infra/repository/pet"
)

type Container struct {
	ListController   *list.List
	CreateController *create.Create
	AproveController *aprove.Aprove
}

func Initialize() *Container {
	db := persist.Connect()
	petRepo := pet.New(db)
	requestRepo := request.New(db)

	emailService := email.New()

	listUc := listUC.New(requestRepo)
	createUc := createUC.New(petRepo, requestRepo, emailService)
	aproveUc := aproveUC.New(requestRepo, petRepo)

	return &Container{
		ListController:   list.New(listUc),
		CreateController: create.New(createUc),
		AproveController: aprove.New(aproveUc),
	}
}

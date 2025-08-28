package di

import (
	adoptUC "github.com/college-redberry/open-adoption/internal/pet/application/usecase/adopt"
	listUC "github.com/college-redberry/open-adoption/internal/pet/application/usecase/list"
	registerUC "github.com/college-redberry/open-adoption/internal/pet/application/usecase/register"
	updateUC "github.com/college-redberry/open-adoption/internal/pet/application/usecase/update"
	"github.com/college-redberry/open-adoption/internal/pet/infra/persist"
	"github.com/college-redberry/open-adoption/internal/pet/infra/repository/pet"
	"github.com/college-redberry/open-adoption/internal/pet/infra/rest/controller/adopt"
	"github.com/college-redberry/open-adoption/internal/pet/infra/rest/controller/list"
	"github.com/college-redberry/open-adoption/internal/pet/infra/rest/controller/register"
	"github.com/college-redberry/open-adoption/internal/pet/infra/rest/controller/update"
)

type Container struct {
	AdoptController    *adopt.Adopt
	ListController     *list.List
	RegisterController *register.Register
	UpdateController   *update.Update
}

func Initialize() *Container {
	db := persist.Connect()
	petRepo := pet.New(db)

	adoptUc := adoptUC.New(petRepo)
	listUc := listUC.New(petRepo)
	registerUc := registerUC.New(petRepo)
	updateUc := updateUC.New(petRepo)

	return &Container{
		AdoptController:    adopt.New(adoptUc),
		ListController:     list.New(listUc),
		RegisterController: register.New(registerUc),
		UpdateController:   update.New(updateUc),
	}
}

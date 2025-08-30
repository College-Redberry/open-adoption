package routes

import (
	"net/http"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	"github.com/college-redberry/open-adoption/internal/pet/di"
)

func Initialize(container di.Container) []bootstrap.Route {
	return []bootstrap.Route{
		{
			Path: "/pets",
			Children: []bootstrap.Route{
				{
					Path:       "/{id}/adopt",
					Method:     http.MethodPatch,
					Controller: container.AdoptController.Handle,
				},
				{
					Path:       "/",
					Method:     http.MethodGet,
					Controller: container.ListController.Handle,
				},
				{
					Path:        "/",
					Method:      http.MethodPost,
					RequireAuth: true,
					Controller:  container.RegisterController.Handle,
				},
				{
					Path:        "/{id}",
					Method:      http.MethodPut,
					RequireAuth: true,
					Controller:  container.UpdateController.Handle,
				},
				{
					Path:        "/{id}/images",
					Method:      http.MethodPost,
					RequireAuth: true,
					Controller:  container.SaveImagesController.Handle,
				},
			},
		},
	}
}

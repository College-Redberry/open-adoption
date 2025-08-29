package routes

import (
	"net/http"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	"github.com/college-redberry/open-adoption/internal/adoption/di"
)

func Initialize(container di.Container) []bootstrap.Route {
	return []bootstrap.Route{
		{
			Path: "/adoption/requests",
			Children: []bootstrap.Route{
				{
					Path:        "/",
					Method:      http.MethodGet,
					RequireAuth: true,
					Controller:  container.ListController.Handle,
				},
				{
					Path:       "/",
					Method:     http.MethodPost,
					Controller: container.CreateController.Handle,
				},
				{
					Path:        "/{id}/approve",
					Method:      http.MethodPatch,
					RequireAuth: true,
					Controller:  container.AproveController.Handle,
				},
			},
		},
	}
}

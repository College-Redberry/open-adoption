package routes

import (
	"net/http"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	"github.com/college-redberry/open-adoption/internal/auth/di"
)

func Initialize(container di.Container) []bootstrap.Route {
	return []bootstrap.Route{
		{
			Path: "/auth",
			Children: []bootstrap.Route{
				{
					Path:       "/login",
					Method:     http.MethodPost,
					Controller: container.LoginController.Handle,
				},
				{
					Path:       "/logout",
					Method:     http.MethodPost,
					Controller: container.LogoutController.Handle,
				},
				{
					Path:       "/refresh",
					Method:     http.MethodPost,
					Controller: container.RefreshController.Handle,
				},
			},
		},
		{
			Path: "/users",
			Children: []bootstrap.Route{
				{
					Path:       "/",
					Method:     http.MethodPost,
					Controller: container.RegisterController.Handle,
				},
				{
					Path:       "/{id}",
					Method:     http.MethodGet,
					Controller: container.RetrieveController.Handle,
				},
			},
		},
	}
}

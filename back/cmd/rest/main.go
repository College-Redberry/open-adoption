package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	authDi "github.com/college-redberry/open-adoption/internal/auth/di"
	authRoutes "github.com/college-redberry/open-adoption/internal/auth/infra/rest"
)

func main() {
	authContainer := authDi.Initialize()
	authRoutes := authRoutes.Initialize(*authContainer)

	routesV1 := bootstrap.Route{
		Path:     "/v1",
		Children: slices.Concat(authRoutes),
	}

	router := bootstrap.Bootstrap(routesV1, authContainer.AuthService)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}

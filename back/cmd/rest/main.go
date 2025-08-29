package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	adoptionDi "github.com/college-redberry/open-adoption/internal/adoption/di"
	adoptionRoutes "github.com/college-redberry/open-adoption/internal/adoption/infra/rest"
	authDi "github.com/college-redberry/open-adoption/internal/auth/di"
	authRoutes "github.com/college-redberry/open-adoption/internal/auth/infra/rest"
	petDi "github.com/college-redberry/open-adoption/internal/pet/di"
	petRoutes "github.com/college-redberry/open-adoption/internal/pet/infra/rest"
)

func main() {
	authContainer := authDi.Initialize()
	authRoutes := authRoutes.Initialize(*authContainer)

	petContainer := petDi.Initialize()
	petRoutes := petRoutes.Initialize(*petContainer)

	adoptionContainer := adoptionDi.Initialize()
	adoptionRoutes := adoptionRoutes.Initialize(*adoptionContainer)

	routesV1 := bootstrap.Route{
		Path:     "/v1",
		Children: slices.Concat(authRoutes, petRoutes, adoptionRoutes),
	}

	router := bootstrap.Bootstrap(routesV1, authContainer.AuthService)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}

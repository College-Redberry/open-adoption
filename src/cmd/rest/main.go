package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/college-redberry/open-adoption/cmd/rest/bootstrap"
	adoptionDi "github.com/college-redberry/open-adoption/internal/adoption/di"
	"github.com/college-redberry/open-adoption/internal/adoption/infra/constants"
	adoptionRoutes "github.com/college-redberry/open-adoption/internal/adoption/infra/rest"
	authDi "github.com/college-redberry/open-adoption/internal/auth/di"
	authRoutes "github.com/college-redberry/open-adoption/internal/auth/infra/rest"
	petDi "github.com/college-redberry/open-adoption/internal/pet/di"
	petRoutes "github.com/college-redberry/open-adoption/internal/pet/infra/rest"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	m, err := migrate.New("file://../../migrations", constants.DB_URL)
	if err != nil {
		log.Fatalf("failed to init migrate: %v", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No new migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		fmt.Println("Migrations applied successfully!")
	}

	log.Println("Starting server on :80")
	http.ListenAndServe(":80", router)
}

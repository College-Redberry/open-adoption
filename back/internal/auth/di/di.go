package di

import (
	"time"

	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	loginUsecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/login"
	refreshUsecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/refresh"
	registerUsecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/register"
	retrieveUsecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/retrieve"
	"github.com/college-redberry/open-adoption/internal/auth/infra/constants"
	"github.com/college-redberry/open-adoption/internal/auth/infra/persist"
	"github.com/college-redberry/open-adoption/internal/auth/infra/repository/user"
	"github.com/college-redberry/open-adoption/internal/auth/infra/rest/controller/login"
	"github.com/college-redberry/open-adoption/internal/auth/infra/rest/controller/logout"
	"github.com/college-redberry/open-adoption/internal/auth/infra/rest/controller/refresh"
	"github.com/college-redberry/open-adoption/internal/auth/infra/rest/controller/register"
	"github.com/college-redberry/open-adoption/internal/auth/infra/rest/controller/retrieve"
	"github.com/college-redberry/open-adoption/internal/auth/infra/service/auth"
	"github.com/college-redberry/open-adoption/internal/auth/infra/service/hash"
)

type Container struct {
	AuthService        service.AuthService
	LoginController    *login.Login
	RegisterController *register.Register
	RetrieveController *retrieve.Retrieve
	LogoutController   *logout.Logout
	RefreshController  *refresh.Refresh
}

func Initialize() *Container {
	authService := auth.New(
		constants.JWT_SECRET_KEY,
		constants.ISSUER,
		time.Duration(constants.EXPIRATION_TIME)*time.Minute,
		time.Duration(constants.REFRESH_EXPIRATION_TIME)*time.Minute,
	)

	db := persist.Connect()
	hashService := hash.New()
	userRepo := user.New(db)

	loginUC := loginUsecase.New(userRepo, authService, hashService)
	registerUC := registerUsecase.New(userRepo, hashService)
	retrieveUC := retrieveUsecase.New(userRepo)
	refreshUC := refreshUsecase.New(userRepo, authService)

	return &Container{
		AuthService:        authService,
		LoginController:    login.New(loginUC),
		RegisterController: register.New(registerUC),
		RetrieveController: retrieve.New(retrieveUC),
		LogoutController:   logout.New(),
		RefreshController:  refresh.New(refreshUC),
	}
}

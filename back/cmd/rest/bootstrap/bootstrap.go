package bootstrap

import (
	"net/http"

	"github.com/college-redberry/open-adoption/cmd/rest/middleware"
	corsMiddleware "github.com/college-redberry/open-adoption/cmd/rest/middleware/cors"
	errorMiddleware "github.com/college-redberry/open-adoption/cmd/rest/middleware/error"
	permissionMiddleware "github.com/college-redberry/open-adoption/cmd/rest/middleware/permission"
	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Path        string
	Method      string
	RequireAuth bool
	Controller  middleware.HandlerFunc
	Children    []Route
}

func Bootstrap(routes Route, authService service.AuthService) http.Handler {
	mux := chi.NewMux()

	mux.Use(chiMiddleware.RequestID)
	mux.Use(chiMiddleware.RealIP)
	mux.Use(chiMiddleware.Logger)
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(corsMiddleware.Middleware)

	permMiddleware := permissionMiddleware.New(authService)
	errorMiddleware := errorMiddleware.New()

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handle(mux, routes, permMiddleware, errorMiddleware)

	return mux
}

func handle(router chi.Router, route Route, permMiddleware *permissionMiddleware.Middleware, errorMiddleware *errorMiddleware.Middleware) {
	if route.RequireAuth {
		router = router.With(permMiddleware.Handle())
	}

	if route.Method != "" {
		router.MethodFunc(route.Method, route.Path, errorMiddleware.Handle(route.Controller))
	}

	if len(route.Children) > 0 {
		router.Route(route.Path, func(r chi.Router) {
			for _, subRoute := range route.Children {
				handle(r, subRoute, permMiddleware, errorMiddleware)
			}
		})
	}
}

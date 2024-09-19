package app

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/router"
	"github.com/go-chi/chi/v5"
)

func LoadRouter(route chi.Router) {
	// # Load All Service
	bootstrap.LoadRouter(
		route,

		// # List Of Router
		router.NewRouterAuth(),
	)
}

package middleware

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/go-chi/chi/v5"
)

func Global(req *registry.AppRegistry, router chi.Router) {
	// Assign your global middleware on here ...
}

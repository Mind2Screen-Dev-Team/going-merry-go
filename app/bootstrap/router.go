package bootstrap

import "github.com/go-chi/chi/v5"

type LoaderRouter interface {
	Loader(router chi.Router) LoaderRouter
	Route()
}

func LoadRouter(router chi.Router, loaders ...LoaderRouter) {
	if loaders == nil {
		return
	}

	for _, l := range loaders {
		l.Loader(router).Route()
	}
}

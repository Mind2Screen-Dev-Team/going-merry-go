package app

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

func AppProviderLoader(reg *registry.AppRegistry) {
	// # Load All Provider
	bootstrap.LoadProvider(
		context.Background(),

		// Link Dependency
		reg,

		// # List Of Provider

		// add more on here...
	)
}

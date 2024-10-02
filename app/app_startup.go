package app

import (
	"fmt"
	"slices"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
)

func MustLoadDependencyAtStartup(service string, reg *registry.AppRegistry) error {

	if slices.Contains([]string{"rest-api", "grpc-api"}, service) {
		// # Load MySQL First
		reg.Dependency.MySqlDB.Value()
		if err := reg.Dependency.MySqlDB.Error(); err != nil {
			return fmt.Errorf("failed to load depedency mysql db: %w", err)
		}
	}

	return nil
}

package app

import (
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
)

func MustLoadDependencyAtStartup(service string, cfg *appconfig.AppConfig, dep *registry.AppDependency) error {

	if service == "restapi" {
		// # Load MySQL First
		dep.MySqlDB.Value()
		if err := dep.MySqlDB.Error(); err != nil {
			return fmt.Errorf("failed to load depedency mysql db: %w", err)
		}
	}

	return nil
}

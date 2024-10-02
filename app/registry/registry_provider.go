package registry

// # PROVIDER

type AppProvider struct{}

func NewAppProvider() *AppProvider {
	return &AppProvider{}
}

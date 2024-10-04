package registry

import (
	repo_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/api"
)

// # REPOSITORY

type RepositoryRegistry struct {
	// register your repository on here
	User repo_api.UserRepoAPI
}

func NewRepositoryRegistry() *RepositoryRegistry {
	return &RepositoryRegistry{}
}

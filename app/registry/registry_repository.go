package registry

import (
	repo_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/api"
)

// # REPOSITORY

type AppRepository struct {
	// register your repository on here
	User repo_api.UserRepoAPI
}

func NewAppRepository() *AppRepository {
	return &AppRepository{}
}

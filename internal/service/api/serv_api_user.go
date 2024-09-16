package service_api

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/entity"
	service_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/attribute"
)

type UserServiceAPI interface {
	Find(ctx context.Context, p service_attribute.UserFindAttribute) (*entity.User, error)
}

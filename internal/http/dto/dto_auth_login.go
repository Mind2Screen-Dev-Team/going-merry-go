package dto

import (
	"context"
	"errors"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	repo_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/attribute"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xtracer"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xvalidate"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/guregu/null.v4"
)

type AuthLoginPayloadReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginReqDTO struct {
	Payload AuthLoginPayloadReqDTO `in:"body=json" json:"payload"`
}

func (a *AuthLoginReqDTO) ValidateWithContext(ctx context.Context) error {
	ctx, span := xtracer.Start(ctx, "dto.auth.login.validation")
	defer span.End()

	return xvalidate.WrapperValidation(validation.ValidateStructWithContext(ctx, &a.Payload,
		// Email cannot be empty, email must exist on database
		validation.Field(&a.Payload.Email, validation.Required, a.IsEmailExists(ctx)),
		// Password cannot be empty, and the length min length is 8
		validation.Field(&a.Payload.Password, validation.Required, validation.Length(8, 0)),
	))
}

func (a *AuthLoginReqDTO) IsEmailExists(ctx context.Context) validation.Rule {
	return validation.By(func(value any) error {
		ctx, span := xtracer.Start(ctx, "dto.auth.login.validation.check.email.existence")
		defer span.End()

		v, _ := value.(string)
		repo := config.LoadRepositoryRegistry(ctx)
		if repo == nil || repo.User == nil {
			err := errors.New("invalid load user repository")
			span.RecordError(err)
			return validation.NewInternalError(err)
		}

		count, err := repo.User.Count(ctx, repo_attribute.UserFindAttribute{
			Email: null.NewString(v, true),
		})
		if err != nil || count <= 0 {
			if err != nil {
				span.RecordError(err)
			}
			return validation.NewError("email_not_exists", "user email not exists")
		}

		return nil
	})
}

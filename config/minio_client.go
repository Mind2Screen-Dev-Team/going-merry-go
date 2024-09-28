package config

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioClient struct{}

func NewMinioClient() *minioClient {
	return &minioClient{}
}

func (minioClient) Loader(ctx context.Context, cfg *appconfig.AppConfig, app *registry.AppDependency) {
	app.Storage = xlazy.New(func() (*minio.Client, error) {
		return minio.New(cfg.Minio.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(cfg.Minio.Credential.AccessKeyId, cfg.Minio.Credential.SecretAccessKey, cfg.Minio.Credential.Token),
			Secure: cfg.Minio.UseSSL,
		})
	})
}

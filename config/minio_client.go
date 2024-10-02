package config

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioClient struct{}

func NewMinioClient() *minioClient {
	return &minioClient{}
}

func (minioClient) Loader(ctx context.Context, reg *registry.AppRegistry) {
	reg.Dependency.Storage = xlazy.New(func() (*minio.Client, error) {
		return minio.New(reg.Config.Minio.Endpoint, &minio.Options{
			Creds: credentials.NewStaticV4(
				reg.Config.Minio.Credential.AccessKeyId,
				reg.Config.Minio.Credential.SecretAccessKey,
				reg.Config.Minio.Credential.Token,
			),
			Secure: reg.Config.Minio.UseSSL,
		})
	})
}

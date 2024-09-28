// Code generated from Pkl module `MinioConfig`. DO NOT EDIT.
package minioconfig

type Credential struct {
	// Credential Enabled
	Enabled bool `pkl:"enabled"`

	// Credential access key id
	AccessKeyId string `pkl:"accessKeyId"`

	// Credential secret access key
	SecretAccessKey string `pkl:"secretAccessKey"`

	// Credential secret token
	Token string `pkl:"token"`
}

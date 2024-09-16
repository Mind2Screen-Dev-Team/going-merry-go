// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

type AppHttp struct {
	Port int `pkl:"port"`

	IdleTimeout int `pkl:"idleTimeout"`

	ReadHeaderTimeout int `pkl:"readHeaderTimeout"`

	ReadTimeout int `pkl:"readTimeout"`

	WriteTimeout int `pkl:"writeTimeout"`

	HandlerTimeout int `pkl:"handlerTimeout"`
}

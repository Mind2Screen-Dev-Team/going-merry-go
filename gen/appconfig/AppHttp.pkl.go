// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

type AppHttp struct {
	// App HTTP Port
	Port int `pkl:"port"`

	// App HTTP Idle Timeout
	IdleTimeout int `pkl:"idleTimeout"`

	// App HTTP Read Header Timeout
	ReadHeaderTimeout int `pkl:"readHeaderTimeout"`

	// App HTTP Read Timeout
	ReadTimeout int `pkl:"readTimeout"`

	// App HTTP Write Timeout
	WriteTimeout int `pkl:"writeTimeout"`

	// App HTTP Handler Timeout
	HandlerTimeout int `pkl:"handlerTimeout"`
}

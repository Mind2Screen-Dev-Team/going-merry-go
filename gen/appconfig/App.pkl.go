// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

type App struct {
	// App Name
	Name string `pkl:"name"`

	// App Domain
	Domain string `pkl:"domain"`

	// App Host
	Host string `pkl:"host"`

	// App JWT Configuration
	Jwt *Jwt `pkl:"jwt"`

	// App Log Configuration
	Log *Log `pkl:"log"`

	// App Http Configuration
	Http *Http `pkl:"http"`
}

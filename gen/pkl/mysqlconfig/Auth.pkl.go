// Code generated from Pkl module `MySQLConfig`. DO NOT EDIT.
package mysqlconfig

type Auth struct {
	// MySql authorization is enabled
	Enabled bool `pkl:"enabled"`

	// MySql authorization username
	Username string `pkl:"username"`

	// MySql authorization password
	Password string `pkl:"password"`
}

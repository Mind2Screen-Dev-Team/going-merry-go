package config

// # DSN (Data Source Name)
//   - is a data structure containing information about a specific database to which an Open Database Connectivity (ODBC) driver needs to connect.
type DSN struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
}

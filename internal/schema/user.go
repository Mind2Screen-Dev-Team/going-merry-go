package schema

import "gopkg.in/guregu/null.v4"

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	CreatedAt null.Time
	UpdatedAt null.Time
}

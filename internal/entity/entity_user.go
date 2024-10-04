package entity

import "gopkg.in/guregu/null.v4"

type User struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt null.Time `db:"created_at"`
	UpdatedAt null.Time `db:"updated_at"`
}

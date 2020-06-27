package models

import "database/sql"

type Users struct {
	ID          string         `db:"id"`
	FullName    string         `db:"full_name"`
	Email       string         `db:"email"`
	Password    string         `db:"password"`
	MobilePhone string         `db:"mobile_phone"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}

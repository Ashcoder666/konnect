package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	UserName  string    `gorm:"unique; not null" json:"username"`
	Email     string    `gorm:"unique; not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Role      UserRole  `gorm:"type:user_role;not null" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

type UserRole string

const (
	Customer UserRole = "customer"
	Partner  UserRole = "partner"
)

func IsDuplicateKeyError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}

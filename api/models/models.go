// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type RefreshToken struct {
	ID        int32
	UserID    pgtype.Int8
	TokenHash string
	IssuedAt  pgtype.Timestamp
	ExpiresAt pgtype.Timestamp
	Revoked   pgtype.Bool
	RevokedAt pgtype.Timestamp
}

type User struct {
	ID              int32
	FirstName       string
	LastName        string
	Password        string
	RefreshToken    pgtype.Text
	Email           string
	IsEmailVerified pgtype.Bool
	IsActive        pgtype.Bool
	GithubHandle    pgtype.Text
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email, is_email_verified, is_active, github_handle, password
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, first_name, last_name, password, email, is_email_verified, is_active, github_handle
`

type CreateUserParams struct {
	FirstName       string
	LastName        string
	Email           string
	IsEmailVerified pgtype.Bool
	IsActive        pgtype.Bool
	GithubHandle    pgtype.Text
	Password        string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.IsEmailVerified,
		arg.IsActive,
		arg.GithubHandle,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Email,
		&i.IsEmailVerified,
		&i.IsActive,
		&i.GithubHandle,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, password, email, is_email_verified, is_active, github_handle FROM users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Email,
		&i.IsEmailVerified,
		&i.IsActive,
		&i.GithubHandle,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, password, email, is_email_verified, is_active, github_handle FROM users 
ORDER BY first_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Password,
			&i.Email,
			&i.IsEmailVerified,
			&i.IsActive,
			&i.GithubHandle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
  set first_name = $2,
  last_name = $3
WHERE id = $1
`

type UpdateUserParams struct {
	ID        int32
	FirstName string
	LastName  string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.FirstName, arg.LastName)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auth.sql

package repository

import (
	"context"
)

const registerUser = `-- name: RegisterUser :exec
INSERT INTO "user" (
 id, email, password_hash 
) VALUES ( uuid_generate_v4(), $1, $2 )
`

type RegisterUserParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) error {
	_, err := q.db.Exec(ctx, registerUser, arg.Email, arg.PasswordHash)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: user_quries.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    
username,email,password,phone_number
)VALUES(
   
$1,$2,$3,$4
)
RETURNING id
`

type CreateUserParams struct {
	Username    string
	Email       string
	Password    string
	PhoneNumber sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.PhoneNumber,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, uuid, username, email, password, phone_number FROM users
WHERE email=$1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
	)
	return i, err
}
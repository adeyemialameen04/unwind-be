// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `binding:"required" json:"title"`
	Author    string    `binding:"required" json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID         uuid.UUID   `json:"id"`
	Name       pgtype.Text `json:"name"`
	Username   pgtype.Text `json:"username"`
	Email      string      `binding:"required,email" json:"email"`
	Password   string      `binding:"required" json:"password"`
	ProfilePic pgtype.Text `json:"profilePic"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

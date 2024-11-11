// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `binding:"required,uuid" json:"userId"`
	ProfilePic *string   `json:"profilePic"`
	Name       *string   `json:"name"`
	Username   string    `binding:"required,min=8" json:"username"`
	CoverPic   *string   `json:"coverPic"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `binding:"required,email" json:"email"`
	Password  string    `binding:"required" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WatchList struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	AnilistID *string   `binding:"required" json:"anilistId"`
	HianimeID *string   `binding:"required" json:"hianimeId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

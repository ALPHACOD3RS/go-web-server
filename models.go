package main

import (
	"time"

	"github.com/ALPHACOD3RS/go-web-server/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.NullUUID `json:"id"`
	Name      string		`json:"name"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}


func dbUserToUser(dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}

}
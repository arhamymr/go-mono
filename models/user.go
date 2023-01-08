package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             uint `gorm:"primaryKey"`
	Email          string
	Name           string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func Hash(password string) string {
	return "hashed password"
}

func VerifyPassword(HashedPassword, password string) bool {
	return Hash(password) == "hashed password"
}

func CreateUser(db *gorm.DB, payload User) *gorm.DB {

	user := &User{
		Email:          payload.Email,
		Name:           payload.Name,
		HashedPassword: payload.HashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	result := db.Create(user)
	return result
}

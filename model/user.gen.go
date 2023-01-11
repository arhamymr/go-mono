// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "User"

// User mapped from table <User>
type User struct {
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email          string    `gorm:"column:email;not null" json:"email"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	HashedPassword string    `gorm:"column:hashed_password;not null" json:"hashed_password"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Deleted        bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

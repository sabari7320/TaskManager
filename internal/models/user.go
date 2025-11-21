package models

import "gorm.io/gorm"

// User example
// @Description User model
type User struct {
	gorm.Model `swaggerignore:"true"`

	Email    string `gorm:"unique" json:"email" validate:"email"`
	Password string `json:"-" validate:"min=4"`
}

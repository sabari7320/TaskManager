package models

import "gorm.io/gorm"

// Task example
// @Description Task model
type Task struct {
	gorm.Model `swaggerignore:"true"`
	Title      string `json:"title" binding:"required"`
	Done       bool   `json:"done"`
	UserID     uint   `json:"-"`                                    // Foreign key
	User       User   `gorm:"constraint:OnDelete:CASCADE" json:"-"` // Belongs To
}

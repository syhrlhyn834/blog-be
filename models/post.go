package models

import "time"

type Post struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"unique"`
	Slug        string    `json:"slug" gorm:"unique"`
	CategoryID  int       `json:"category_id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	Image       string    `json:"image"`
	Status      string    `json:"status" gorm:"size:50;default:archive"`
	Views       uint64    `json:"views" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

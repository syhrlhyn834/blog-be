package models

import (
	"fmt"
	"net/http"
	"time"
)

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
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
}

// Fungsi Getter untuk properti ImageURL
func (post *Post) GetImageURL(r *http.Request) string {
	return fmt.Sprintf("%s://%s/%s", getScheme(r), r.Host, post.Image)
}

// Fungsi bantu untuk menentukan skema (http atau https)
func getScheme(r *http.Request) string {
	if r.Header.Get("X-Forwarded-Proto") != "" {
		return r.Header.Get("X-Forwarded-Proto")
	}
	return "http"
}

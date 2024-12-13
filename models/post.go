package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

func (post *Post) GetImageURL(c *gin.Context) string {
	return fmt.Sprintf("http://%s/%s", c.Request.Host, post.Image)
}

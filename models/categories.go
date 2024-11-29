package models

import "time"

type Category struct {
	Id        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"unique"`
	Slug      string     `json:"slug" gorm:"unique"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

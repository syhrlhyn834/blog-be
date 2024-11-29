package models

import "time"

type Menu struct {
	Id        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

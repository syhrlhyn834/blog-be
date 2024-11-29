package models

import "time"

type Sosmed struct {
	Id        int        `json:"id" gorm:"primary_key"`
	Logo      string     `json:"logo"`
	Url       string     `json:"url"`
	Username  string     `json:"username" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

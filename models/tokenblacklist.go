package models

import "time"

type TokenBlacklist struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Token     string    `json:"token" gorm:"type:text;unique"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

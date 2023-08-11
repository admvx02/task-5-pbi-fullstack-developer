package models

import (
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"-"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

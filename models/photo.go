package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Comments  []Comment
	DeletedAt gorm.DeletedAt
}

type CreatePhotoRes struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoRes struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      UserEtc
}

type UpdatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentPhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   string `json:"user_id"`
}

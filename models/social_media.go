package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" gorm:"not null"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"not null"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User
	DeletedAt      gorm.DeletedAt
}

type CreateSosmedRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetSosmedRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserEtc   `json:"user"`
}

type UpdateSosmedRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         string    `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

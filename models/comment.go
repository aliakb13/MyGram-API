package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int       `json:"comment_id"`
	UserID    string    `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Photo     Photo
	DeletedAt gorm.DeletedAt
}

type PostCommentRes struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentRes struct {
	ID        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoId   int          `json:"photo_id"`
	UserId    string       `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	User      UserEtc      `json:"user"`
	Photo     CommentPhoto `json:"photo"`
}

type UpdateCommentRes struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

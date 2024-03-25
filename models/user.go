package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `json:"user_id" gorm:"not null;unique"`
	Username    string         `json:"username" gorm:"not null;unique"`
	Email       string         `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password    string         `json:"password" gorm:"not null" validate:"required,min=6"`
	Age         int            `json:"age" gorm:"not null" validate:"required,gt=8"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Photos      []Photo        `json:"photos"`
	Comments    []Comment      `json:"comments"`
	SocialMedia []SocialMedia  `json:"social_media"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	// u.CreatedAt = time.Now()
	return nil
}

func (u *User) Validate() error {
	// var validate *validator.Validate
	validate := validator.New()
	return validate.Struct(u)
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type UserRegisterResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       string `json:"user_id"`
	Username string `json:"username"`
}

type UserUpdateResponse struct {
	ID        string    `json:"user_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserEtc struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

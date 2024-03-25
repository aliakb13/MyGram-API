package interfaces

import "final-project/models"

type UserInterface interface {
	RegisterUser(user models.User) (models.UserRegisterResponse, error)
	GetByEmail(email string) (models.User, error)
	UpdateUser(id string, email string, username string) (models.UserUpdateResponse, error)
	DeleteUser(id string) (bool, error)
}

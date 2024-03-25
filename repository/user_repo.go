package repository

import (
	"final-project/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) RegisterUser(user models.User) (models.UserRegisterResponse, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return models.UserRegisterResponse{}, err
	}

	// var user models.User

	ur.db.Where("email = ?", user.Email).First(&user)

	return models.UserRegisterResponse{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (ur *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := ur.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (ur *userRepository) UpdateUser(id string, email string, username string) (models.UserUpdateResponse, error) {
	user := models.User{
		Email:    email,
		Username: username,
	}

	err := ur.db.Model(&user).Where("id = ? AND deleted_at IS NULL", id).Updates(user).Error

	if err != nil {
		return models.UserUpdateResponse{}, err
	}

	userField := models.User{}

	err = ur.db.First(&userField, "id = ?", id).Error

	if err != nil {
		return models.UserUpdateResponse{}, err
	}

	userRes := models.UserUpdateResponse{
		ID:        userField.ID,
		Email:     userField.Email,
		Username:  userField.Username,
		Age:       userField.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes, nil
}

func (ur *userRepository) DeleteUser(id string) (bool, error) {
	user := models.User{
		ID: id,
	}
	err := ur.db.Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ur *userRepository) GetById(id string) (models.UserEtc, error) {
	var user models.User
	err := ur.db.First(&user, id).Error

	if err != nil {
		return models.UserEtc{}, err
	}

	return models.UserEtc{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

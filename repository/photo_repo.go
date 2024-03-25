package repository

import (
	"errors"
	"final-project/models"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) CreatePhoto(reqPhoto models.Photo) (models.CreatePhotoRes, error) {
	var photo models.Photo

	err := pr.db.Create(&reqPhoto).Error

	if err != nil {
		return models.CreatePhotoRes{}, err
	}

	err = pr.db.Order("id desc").First(&photo).Error

	if err != nil {
		return models.CreatePhotoRes{}, err
	}

	return models.CreatePhotoRes{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}, nil
}

func (pr *photoRepository) GetAllPhotos(userId string) ([]models.Photo, error) {
	photos := []models.Photo{}
	err := pr.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Find(&photos).Error

	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (pr *photoRepository) GetPhotoById(id int, userId string) (models.Photo, error) {
	var photo models.Photo
	err := pr.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).First(&photo, id).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (pr *photoRepository) UpdatePhoto(photoId int, reqPhoto models.Photo) (models.UpdatePhotoResponse, error) {
	reqPhoto.ID = photoId

	err := pr.db.Model(&reqPhoto).Where("user_id = ? ", reqPhoto.UserID).Updates(&reqPhoto).Error

	if err != nil {
		return models.UpdatePhotoResponse{}, err
	}

	var updatedPhoto models.Photo

	err = pr.db.Where("id = ? AND user_id = ?", photoId, reqPhoto.UserID).First(&updatedPhoto).Error

	if err != nil {
		return models.UpdatePhotoResponse{}, errors.New("failed to catching updated photo")
	}

	return models.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoUrl:  updatedPhoto.PhotoUrl,
		UserID:    updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}, nil
}

func (pr *photoRepository) DeletePhoto(photoId int, userId string) error {
	err := pr.db.Where("id = ? AND user_id = ?", photoId, userId).Delete(&models.Photo{}).Error

	if err != nil {
		return err
	}

	return nil
}

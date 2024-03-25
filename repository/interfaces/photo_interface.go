package interfaces

import "final-project/models"

type PhotoInterface interface {
	CreatePhoto(reqPhoto models.Photo) (models.CreatePhotoRes, error)
	GetAllPhotos(userId string) ([]models.Photo, error)
	GetPhotoById(id int, userId string) (models.Photo, error)
	UpdatePhoto(photoId int, reqPhoto models.Photo) (models.UpdatePhotoResponse, error)
	DeletePhoto(photoId int, userId string) error
}

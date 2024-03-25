package interfaces

import "final-project/models"

type SocialMediaInterface interface {
	CreateSocialMedia(reqSosmed models.SocialMedia) (models.CreateSosmedRes, error)
	GetAllSocialMedia(userId string) ([]models.SocialMedia, error)
	GetSocialMediaById(id int, userId string) (models.SocialMedia, error)
	UpdateSocialMedia(sosmedId int, reqSosmed models.SocialMedia) (models.UpdateSosmedRes, error)
	DeleteSocialMedia(sosmedId int, userId string) error
}

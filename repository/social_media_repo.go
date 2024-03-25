package repository

import (
	"final-project/models"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (scm *socialMediaRepository) CreateSocialMedia(reqSosmed models.SocialMedia) (models.CreateSosmedRes, error) {
	var createdSosmed models.SocialMedia
	err := scm.db.Create(&reqSosmed).Error

	if err != nil {
		return models.CreateSosmedRes{}, nil
	}

	err = scm.db.Order("id desc").First(&createdSosmed).Error

	if err != nil {
		return models.CreateSosmedRes{}, nil
	}

	// fmt.Println(createdSosmed)

	return models.CreateSosmedRes{
		ID:             createdSosmed.ID,
		Name:           createdSosmed.Name,
		SocialMediaUrl: createdSosmed.SocialMediaUrl,
		UserId:         createdSosmed.UserID,
		CreatedAt:      createdSosmed.CreatedAt,
	}, nil
}

func (scm *socialMediaRepository) GetAllSocialMedia(userId string) ([]models.SocialMedia, error) {
	socialMedias := []models.SocialMedia{}
	err := scm.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Find(&socialMedias).Error

	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

func (scm *socialMediaRepository) GetSocialMediaById(id int, userId string) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := scm.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).First(&socialMedia, id).Error

	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (scm *socialMediaRepository) UpdateSocialMedia(sosmedId int, reqSosmed models.SocialMedia) (models.UpdateSosmedRes, error) {
	sosmed := models.SocialMedia{
		ID: sosmedId,
	}

	err := scm.db.Model(&sosmed).Where("user_id = ? ", reqSosmed.UserID).Updates(&reqSosmed).Error

	if err != nil {
		return models.UpdateSosmedRes{}, err
	}

	// var updatedSosmed = models.SocialMedia{
	// 	ID: sosmedId,
	// }

	err = scm.db.Where("user_id = ?", reqSosmed.UserID).First(&sosmed).Error

	if err != nil {
		return models.UpdateSosmedRes{}, err
	}

	return models.UpdateSosmedRes{
		ID:             sosmed.ID,
		Name:           sosmed.Name,
		SocialMediaUrl: sosmed.SocialMediaUrl,
		UserId:         sosmed.UserID,
		UpdatedAt:      sosmed.UpdatedAt,
	}, nil
}

func (scm *socialMediaRepository) DeleteSocialMedia(sosmedId int, userId string) error {
	socialMedia := models.SocialMedia{
		ID: sosmedId,
	}

	err := scm.db.Where("user_id = ?", userId).Delete(&socialMedia).Error

	if err != nil {
		return err
	}

	return nil

}

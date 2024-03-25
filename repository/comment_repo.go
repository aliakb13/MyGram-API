package repository

import (
	"errors"
	"final-project/models"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) PostComment(reqComment models.Comment) (models.PostCommentRes, error) {
	var comment models.Comment
	err := cr.db.Create(&reqComment).Error

	if err != nil {
		return models.PostCommentRes{}, err
	}

	err = cr.db.Order("id desc").First(&comment).Error

	if err != nil {
		return models.PostCommentRes{}, errors.New("error getting create comment")
	}

	return models.PostCommentRes{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}, nil
}

func (cr *commentRepository) GetAllComments(userId string) ([]models.Comment, error) {
	comments := []models.Comment{}
	err := cr.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, caption, photo_url, user_id")
	}).Find(&comments).Error

	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (cr *commentRepository) GetCommentById(id int, userId string) (models.Comment, error) {
	var comment models.Comment
	err := cr.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, caption, photo_url, user_id")
	}).First(&comment, id).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (cr *commentRepository) EditComment(commentId int, reqComment models.Comment) (models.UpdateCommentRes, error) {
	comment := models.Comment{
		ID: commentId,
	}

	err := cr.db.Model(&comment).Where("user_id = ?", reqComment.UserID).Update("message", reqComment.Message).Error

	if err != nil {
		return models.UpdateCommentRes{}, nil
	}

	var updatedComment models.Comment

	err = cr.db.Where("id = ? AND user_id = ?", commentId, reqComment.UserID).First(&updatedComment).Error

	if err != nil {
		return models.UpdateCommentRes{}, errors.New("failed to catching updated comment")
	}

	return models.UpdateCommentRes{
		ID:        updatedComment.ID,
		Message:   updatedComment.Message,
		PhotoId:   updatedComment.PhotoID,
		UserId:    updatedComment.UserID,
		UpdatedAt: updatedComment.UpdatedAt,
	}, nil
}

func (cr *commentRepository) DeleteComment(commentId int, userId string) (bool, error) {
	comment := models.Comment{
		ID: commentId,
	}

	err := cr.db.Where("user_id = ?", userId).Delete(&comment).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

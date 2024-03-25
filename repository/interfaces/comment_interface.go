package interfaces

import "final-project/models"

type CommentInterface interface {
	PostComment(reqComment models.Comment) (models.PostCommentRes, error)
	GetAllComments(userId string) ([]models.Comment, error)
	GetCommentById(id int, userId string) (models.Comment, error)
	EditComment(commentId int, reqComment models.Comment) (models.UpdateCommentRes, error)
	DeleteComment(commentId int, userId string) (bool, error)
}

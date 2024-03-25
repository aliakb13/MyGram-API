package controller

import (
	"final-project/models"
	"final-project/repository/interfaces"
	"final-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository interfaces.CommentInterface
}

func NewCommentController(commentRepository interfaces.CommentInterface) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

func (cc *commentController) PostComment(ctx *gin.Context) {
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to binding",
			"error":   err.Error(),
		})
		return
	}

	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims does not exist",
		})
		return
	}

	userId, err := util.GetIdFromClaims(claims)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims not valid",
			"error":   err.Error(),
		})
		return
	}

	comment.UserID = userId

	commentRes, err := cc.commentRepository.PostComment(comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to creating comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, commentRes)
}

func (cc *commentController) GetAllComments(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims does not exist",
		})
		return
	}

	userId, err := util.GetIdFromClaims(claims)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims not valid",
			"error":   err.Error(),
		})
		return
	}

	comments, err := cc.commentRepository.GetAllComments(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error getting comments",
			"error":   err.Error(),
		})
		return
	}

	fetchComments := []models.GetCommentRes{}

	for _, comment := range comments {
		fetchComments = append(fetchComments, models.GetCommentRes{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoId:   comment.PhotoID,
			UserId:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			User: models.UserEtc{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: models.CommentPhoto{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	ctx.JSON(http.StatusOK, fetchComments)
}

func (cc *commentController) GetCommentById(ctx *gin.Context) {
	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	comment, err := cc.commentRepository.GetCommentById(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
			"error":   err.Error(),
		})
		return
	}

	requestedComment := models.GetCommentRes{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		User: models.UserEtc{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: models.CommentPhoto{
			ID:       comment.Photo.ID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoUrl: comment.Photo.PhotoUrl,
			UserID:   comment.Photo.UserID,
		},
	}

	ctx.JSON(http.StatusOK, requestedComment)
}

func (cc *commentController) EditComment(ctx *gin.Context) {
	var reqComment models.Comment

	err := ctx.ShouldBindJSON(&reqComment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error binding json",
			"error":   err.Error(),
		})
		return
	}

	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	reqComment.UserID = objParam.UserId

	updatedComment, err := cc.commentRepository.EditComment(objParam.Param, reqComment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Comment not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	isDeleted, err := cc.commentRepository.DeleteComment(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting data",
			"error":   err,
		})
		return
	}

	if !isDeleted {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

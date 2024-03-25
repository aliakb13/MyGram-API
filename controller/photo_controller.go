package controller

import (
	"final-project/models"
	"final-project/repository/interfaces"
	"final-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoRepository interfaces.PhotoInterface
	// userRepository  *repository.userRepository
}

func NewPhotoController(photoRepository interfaces.PhotoInterface) *photoController {
	return &photoController{
		photoRepository: photoRepository,
		// userRepository:  userRepository,
	}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	err := ctx.ShouldBindJSON(&photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "binding not success",
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

	id, err := util.GetIdFromClaims(claims)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims not valid",
			"error":   err.Error(),
		})
		return
	}

	photo.UserID = id

	createdPhoto, err := pc.photoRepository.CreatePhoto(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Record not found while create",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success creating new photo",
		"data":    createdPhoto,
	})

}

func (pc *photoController) GetAllPhotos(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims does not exist",
		})
		return
	}

	id, err := util.GetIdFromClaims(claims)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims not valid",
			"error":   err.Error(),
		})
		return
	}

	photos, err := pc.photoRepository.GetAllPhotos(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error getting photos",
			"error":   err.Error(),
		})
		return
	}

	fetchPhotos := []models.GetPhotoRes{}

	for _, photo := range photos {
		fetchPhotos = append(fetchPhotos, models.GetPhotoRes{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			User: models.UserEtc{
				ID:       photo.User.ID,
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, fetchPhotos)
}

func (pc *photoController) GetPhotoById(ctx *gin.Context) {

	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	photo, err := pc.photoRepository.GetPhotoById(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error getting photos",
			"error":   err.Error(),
		})
		return
	}

	resPhoto := models.GetPhotoRes{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		User: models.UserEtc{
			ID:       photo.User.ID,
			Email:    photo.User.Email,
			Username: photo.User.Username,
		},
	}

	ctx.JSON(http.StatusOK, resPhoto)
}

func (pc *photoController) UpdatePhoto(ctx *gin.Context) {
	var photo models.Photo

	err := ctx.ShouldBindJSON(&photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error binding json",
			"error":   err.Error(),
		})
		return
	}

	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	photo.UserID = objParam.UserId

	updatedPhoto, err := pc.photoRepository.UpdatePhoto(objParam.Param, photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Photo not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)

}

func (pc *photoController) DeletePhoto(ctx *gin.Context) {
	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	err := pc.photoRepository.DeletePhoto(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "error deleting",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

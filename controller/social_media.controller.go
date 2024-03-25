package controller

import (
	"final-project/models"
	"final-project/repository/interfaces"
	"final-project/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaRepository interfaces.SocialMediaInterface
}

func NewSocialMediaController(socialMediaRepository interfaces.SocialMediaInterface) *socialMediaController {
	return &socialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (smc *socialMediaController) CreatedSocialMedia(ctx *gin.Context) {
	var sosmed models.SocialMedia

	err := ctx.ShouldBind(&sosmed)

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

	sosmed.UserID = userId

	createdSosmed, err := smc.socialMediaRepository.CreateSocialMedia(sosmed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to getting creating data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdSosmed)
}

func (smc *socialMediaController) GetAllSosmed(ctx *gin.Context) {
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

	socialMedias, err := smc.socialMediaRepository.GetAllSocialMedia(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error getting social media",
			"error":   err.Error(),
		})
		return
	}

	fetchSocialMedia := []models.GetSosmedRes{}

	for _, socialMedia := range socialMedias {
		fetchSocialMedia = append(fetchSocialMedia, models.GetSosmedRes{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserId:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: models.UserEtc{
				ID:       socialMedia.User.ID,
				Email:    socialMedia.User.Email,
				Username: socialMedia.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, fetchSocialMedia)

}

func (smc *socialMediaController) GetSosmedById(ctx *gin.Context) {
	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	sosmed, err := smc.socialMediaRepository.GetSocialMediaById(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
			"error":   err.Error(),
		})
		return
	}

	requestedSosmed := models.GetSosmedRes{
		ID:             sosmed.ID,
		Name:           sosmed.Name,
		SocialMediaUrl: sosmed.SocialMediaUrl,
		UserId:         sosmed.UserID,
		CreatedAt:      sosmed.CreatedAt,
		UpdatedAt:      sosmed.UpdatedAt,
		User: models.UserEtc{
			ID:       sosmed.User.ID,
			Email:    sosmed.User.Email,
			Username: sosmed.User.Username,
		},
	}
	ctx.JSON(http.StatusOK, requestedSosmed)
}

func (smc *socialMediaController) UpdateSosmed(ctx *gin.Context) {

	var reqSosmed models.SocialMedia

	err := ctx.ShouldBindJSON(&reqSosmed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error binding json",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(reqSosmed)

	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	reqSosmed.UserID = objParam.UserId

	updated, err := smc.socialMediaRepository.UpdateSocialMedia(objParam.Param, reqSosmed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "sosmed not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (smc *socialMediaController) DeleteSosmed(ctx *gin.Context) {
	objParam := ctx.MustGet("paramUserId").(util.ParamAndUserId)

	err := smc.socialMediaRepository.DeleteSocialMedia(objParam.Param, objParam.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting data",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}

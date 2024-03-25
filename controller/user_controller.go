package controller

import (
	"final-project/models"
	"final-project/repository/interfaces"
	"final-project/util"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository interfaces.UserInterface
}

func NewUserController(userRepo interfaces.UserInterface) *userController {
	return &userController{
		userRepository: userRepo,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "binding not success",
			"error":   err,
		})
		return
	}

	if err := user.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encrpytPass, err := util.HashPassword(user.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error encrypting password",
			"error":   err,
		})
		return
	}

	user.Password = encrpytPass

	userResponse, err := uc.userRepository.RegisterUser(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "created not success",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success creating user",
		"data":    userResponse,
	})
}

func (uc *userController) Login(ctx *gin.Context) {
	var reqLogin models.UserLogin
	err := ctx.ShouldBindJSON(&reqLogin)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "error read or bad request",
			"error":   err,
		})
		return
	}
	// fmt.Println(reqLogin)

	user, err := uc.userRepository.GetByEmail(reqLogin.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "email not found, please check your email",
			"error":   err.Error(),
		})
		return
	}

	isTrue := util.ComparedPassword(user.Password, reqLogin.Password)

	if !isTrue {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "password not match",
		})
		return
	}

	token, err := util.GenerateToken(user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error generating token",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var user models.UserUpdate
	err := ctx.ShouldBindJSON(&user)
	// userId := ctx.Param("id")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "error read from body",
			"error":   err,
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
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "error from getting id from claims",
			"error":   err.Error(),
		})
		return
	}

	updateUser, err := uc.userRepository.UpdateUser(id, user.Email, user.Username)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, updateUser)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	// userId := ctx.Param("id")

	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "claims does not exist",
		})
		return
	}

	id, err := util.GetIdFromClaims(claims)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "error from getting id from claims",
			"error":   err.Error(),
		})
		return
	}

	isDelete, err := uc.userRepository.DeleteUser(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting data",
			"error":   err,
		})
		return
	}

	if !isDelete {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Record successfully deleted",
	})
}

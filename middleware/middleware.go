package middleware

import (
	"final-project/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	authorizationVal := ctx.GetHeader("Authorization")
	separatedAut := strings.Split(authorizationVal, " ")

	if len(separatedAut) == 1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, token not found",
		})
		return
	}

	jwtToken := separatedAut[1]

	claims, err := util.VerifyToken(jwtToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, verifying token error",
			"error":   err.Error(),
		})
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()

}

func UserAuthorization(ctx *gin.Context) {
	stringId := ctx.Param("id")

	realId, err := strconv.Atoi(stringId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error converting",
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

	obj := util.ParamAndUserId{
		Param:  realId,
		UserId: userId,
	}

	ctx.Set("paramUserId", obj)

	ctx.Next()

}

// func PhotoAuthorization(ctx *gin.Context) {

// }

// func CommentAuthorization(ctx *gin.Context) {

// }

// func SocialMediaAuthorization(ctx *gin.Context) {

// }

package router

import (
	"final-project/controller"
	"final-project/middleware"
	"final-project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(engine *gin.Engine, db *gorm.DB) {
	engine.GET("/testing", middleware.AuthMiddleware, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello from server")
	})
	// user route
	// controller and repo user

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	usersRoute := engine.Group("/users")
	usersRoute.POST("/register", userController.RegisterUser)
	usersRoute.POST("/login", userController.Login)
	usersRoute.PUT("", middleware.AuthMiddleware, userController.UpdateUser)
	usersRoute.DELETE("", middleware.AuthMiddleware, userController.DeleteUser)

	// photos route
	// controller and repo photo
	photoRepository := repository.NewPhotoRepository(db)
	photoController := controller.NewPhotoController(photoRepository)

	photosRoute := engine.Group("/photos", middleware.AuthMiddleware)
	photosRoute.POST("", photoController.CreatePhoto)
	photosRoute.GET("", photoController.GetAllPhotos)
	photosRoute.GET("/:id", middleware.UserAuthorization, photoController.GetPhotoById) //get one / by id
	photosRoute.PUT("/:id", middleware.UserAuthorization, photoController.UpdatePhoto)
	photosRoute.DELETE("/:id", middleware.UserAuthorization, photoController.DeletePhoto)

	// comments route
	// controller and repo comments
	commentRepository := repository.NewCommentRepository(db)
	commentController := controller.NewCommentController(commentRepository)

	commentsRoute := engine.Group("/comments", middleware.AuthMiddleware)
	commentsRoute.POST("", commentController.PostComment)
	commentsRoute.GET("", commentController.GetAllComments)
	commentsRoute.GET("/:id", middleware.UserAuthorization, commentController.GetCommentById) //get one / by id
	commentsRoute.PUT("/:id", middleware.UserAuthorization, commentController.EditComment)
	commentsRoute.DELETE("/:id", middleware.UserAuthorization, commentController.DeleteComment)

	// social media route
	// controller and repo comments
	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaController := controller.NewSocialMediaController(socialMediaRepository)

	sosmedRoute := engine.Group("/socialmedias", middleware.AuthMiddleware)
	sosmedRoute.POST("", socialMediaController.CreatedSocialMedia)
	sosmedRoute.GET("", socialMediaController.GetAllSosmed)
	sosmedRoute.GET("/:id", middleware.UserAuthorization, socialMediaController.GetSosmedById)
	sosmedRoute.PUT("/:id", middleware.UserAuthorization, socialMediaController.UpdateSosmed)
	sosmedRoute.DELETE("/:id", middleware.UserAuthorization, socialMediaController.DeleteSosmed)
}

package routes

import (
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/services"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, blogController controller.BlogController, jwtService services.JWTService) {
	userRoutes := router.Group("/blogs")
	{
		userRoutes.POST("", middleware.Authenticate(jwtService, "user"), blogController.UploadBlog)
		userRoutes.GET("/get", blogController.GetBlog)
		userRoutes.POST("/like", blogController.Like)
		userRoutes.POST("/comment", blogController.GetCommentBlog)
	}
}

package routes

import (
	"tugas4/controller"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, blogController controller.BlogController) {
	userRoutes := router.Group("/blogs")
	{
		userRoutes.POST("", blogController.UploadBlog)
		userRoutes.GET("/get", blogController.GetBlogComment)
		userRoutes.POST("/like", blogController.Like)
	}
}

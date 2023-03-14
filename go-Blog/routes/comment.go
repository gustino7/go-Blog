package routes

import (
	"tugas4/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, commentController controller.CommentController) {
	userRoutes := router.Group("/comment")
	{
		userRoutes.POST("", commentController.AddComment)
	}
}

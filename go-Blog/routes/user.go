package routes

import (
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController controller.UserController, jwtService services.JWTService) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/update", userController.Update)
		userRoutes.DELETE("/delete", userController.Delete)
		userRoutes.GET("/getUser", middleware.Authenticate(jwtService, "user"), userController.GetUserData)
	}
}

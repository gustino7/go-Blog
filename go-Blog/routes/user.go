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
		userRoutes.GET("/getUser", middleware.Authenticate(jwtService, "user"), userController.GetUserData)
		userRoutes.PUT("/update", middleware.Authenticate(jwtService, "user"), userController.Update)
		userRoutes.DELETE("/delete", middleware.Authenticate(jwtService, "user"), userController.Delete)
	}
}

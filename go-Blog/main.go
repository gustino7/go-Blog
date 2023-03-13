package main

import (
	"tugas4/config"
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/repository"
	"tugas4/routes"
	"tugas4/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDataBaseConnection()

	jwtService := services.NewJWTService()
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)

	blogRepository := repository.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepository)
	blogController := controller.NewBlogController(blogService)
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.UserRoutes(server, userController, jwtService)
	routes.BlogRoutes(server, blogController)

	port := "8080"
	server.Run(":" + port)
}

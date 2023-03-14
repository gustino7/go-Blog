package main

import (
	"os"
	"tugas4/config"
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/repository"
	"tugas4/routes"
	"tugas4/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "Production" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	db := config.SetUpDataBaseConnection()

	jwtService := services.NewJWTService()
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)

	blogRepository := repository.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepository)
	blogController := controller.NewBlogController(blogService, jwtService)

	commentRepo := repository.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService)
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.UserRoutes(server, userController, jwtService)
	routes.BlogRoutes(server, blogController, jwtService)
	routes.CommentRoutes(server, commentController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}

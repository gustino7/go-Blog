package controller

import (
	"fmt"
	"net/http"
	"tugas4/dto"
	"tugas4/entity"
	"tugas4/services"
	"tugas4/utils"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	blogService services.BlogService
}

type BlogController interface {
	UploadBlog(ctx *gin.Context)
	GetBlogComment(ctx *gin.Context)
	Like(ctx *gin.Context)
}

func NewBlogController(blogService services.BlogService) BlogController {
	return &blogController{
		blogService: blogService,
	}
}

func (c *blogController) UploadBlog(ctx *gin.Context) {
	//upload blog
	var dtoBlog dto.UpBlog
	err := ctx.ShouldBind(&dtoBlog)
	fmt.Println(dtoBlog)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	UpBlog, err := c.blogService.UploadBlog(ctx, dtoBlog)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to upload blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Blog uploaded", http.StatusCreated, UpBlog)
	ctx.JSON(http.StatusCreated, response)

}

func (c *blogController) GetBlogComment(ctx *gin.Context) {
	//upload blog
	var blog []entity.Blog

	blog, err := c.blogService.GetBlogComment(ctx, entity.Blog{})
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get all blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success get all blog", http.StatusCreated, blog)
	ctx.JSON(http.StatusCreated, response)

}

func (c *blogController) Like(ctx *gin.Context) {
	//like blog
	var blog entity.Blog
	err := ctx.ShouldBind(&blog) //just req title

	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	blog, err = c.blogService.Like(ctx, blog.Title)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to like blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to like blog", http.StatusCreated, blog)
	ctx.JSON(http.StatusCreated, response)

}

// func (c *blogController) Comment(ctx *gin.Context) {
// 	//upload blog
// 	var comment string

// 	blog, err := c.blogService.Comment(ctx, comment)
// 	if err != nil {
// 		response := utils.BuildErrorResponse("Failed to comment", http.StatusBadRequest)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := utils.BuildSuccessResponse("Success to comment", http.StatusCreated, blog)
// 	ctx.JSON(http.StatusCreated, response)

// }

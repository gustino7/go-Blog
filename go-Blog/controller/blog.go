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
	jwtService  services.JWTService
}

type BlogController interface {
	UploadBlog(ctx *gin.Context)
	GetBlog(ctx *gin.Context)
	Like(ctx *gin.Context)
	GetCommentBlog(ctx *gin.Context)
}

func NewBlogController(blogService services.BlogService, jwtService services.JWTService) BlogController {
	return &blogController{
		blogService: blogService,
		jwtService:  jwtService,
	}
}

// upload blog
func (c *blogController) UploadBlog(ctx *gin.Context) {
	var dtoBlog dto.UpBlog
	err := ctx.ShouldBind(&dtoBlog)
	fmt.Println(dtoBlog)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//cek token return user id
	token := ctx.MustGet("token").(string)
	userID, err := c.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find token", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	UpBlog, err := c.blogService.UploadBlog(ctx, dtoBlog, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to upload blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Blog uploaded", http.StatusCreated, UpBlog)
	ctx.JSON(http.StatusCreated, response)

}

func (c *blogController) GetBlog(ctx *gin.Context) {
	//upload blog
	var blog []entity.Blog

	blog, err := c.blogService.GetBlog(ctx, entity.Blog{})
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

// get comment by title blog
func (c *blogController) GetCommentBlog(ctx *gin.Context) {
	var blog entity.Blog
	err := ctx.ShouldBind(&blog)

	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//get blog data and comment
	commentBlog, err := c.blogService.GetCommentBlog(ctx, blog.Title)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get all blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success get all blog", http.StatusCreated, commentBlog)
	ctx.JSON(http.StatusCreated, response)

}

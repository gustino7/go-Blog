package controller

import (
	"github.com/gin-gonic/gin"
)

// type commentController struct {
// 	commentService services.CommentService
// }

type CommentController interface {
	AddComment(ctx *gin.Context, title string, comment string)
}

// func (c *commentController) AddComment(ctx *gin.Context) {
// 	//upload blog
// 	var dtoComment dto.AddComm
// 	err := ctx.ShouldBind(&dtoComment)
// 	fmt.Println(dtoComment)
// 	if err != nil {
// 		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	comment, err := c.commentService.AddComment(ctx, dtoComment.Title, dtoComment.Comment)
// 	if err != nil {
// 		response := utils.BuildErrorResponse("Failed to upload blog", http.StatusBadRequest)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := utils.BuildSuccessResponse("Blog uploaded", http.StatusCreated, comment)
// 	ctx.JSON(http.StatusCreated, response)

//}

package controller

import (
	"net/http"
	"tugas4/dto"

	"tugas4/services"
	"tugas4/utils"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUserData(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewUserController(userService services.UserService, jwt services.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var userDTO dto.RegisterUser

	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(ctx, userDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to register user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("User created", http.StatusCreated, user)
	ctx.JSON(http.StatusCreated, response)

}

func (c *userController) Login(ctx *gin.Context) {
	var userDTO dto.LoginUser
	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.FindUser(ctx, userDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	_, err = c.userService.Login(ctx, userDTO.Email, userDTO.Password)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to verif", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID, user.Role)

	response := utils.BuildSuccessResponse("Login sucessfully", http.StatusAccepted, token)
	ctx.JSON(http.StatusAccepted, response)
}

func (c *userController) GetUserData(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := c.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find token", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.userService.GetUserByID(ctx, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Sucess Get User", http.StatusAccepted, res)
	ctx.JSON(http.StatusAccepted, response)

}

func (c *userController) Update(ctx *gin.Context) {
	var userDTO dto.UserUpdateDTO
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process update", http.StatusBadRequest)
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

	//update name using id
	updUser, err := c.userService.Update(ctx, userDTO.Nama, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Update sucessfully", http.StatusAccepted, updUser)
	ctx.JSON(http.StatusAccepted, response)

}

func (c *userController) Delete(ctx *gin.Context) {
	//cek token return user id
	token := ctx.MustGet("token").(string)
	userID, err := c.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find token", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//delete user
	delUser, err := c.userService.Delete(ctx, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Delete sucessfully", http.StatusAccepted, delUser)
	ctx.JSON(http.StatusAccepted, response)

}

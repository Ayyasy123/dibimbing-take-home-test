package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/response"
	"github.com/Ayyasy123/dibimbing-take-home-test/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req entity.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := c.userService.RegisterUser(&req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusCreated, "User registered successfully", user)
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var req entity.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	userRes, err := c.userService.LoginUser(&req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to login", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Login successful", userRes)
}

func (c *UserController) FindUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	userRes, err := c.userService.FindUserByID(id)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "User data retrieved successfully", userRes)
}

func (c *UserController) FindAllUsers(ctx *gin.Context) {
	usersRes, err := c.userService.FindAllUsers()
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Users data retrieved successfully", usersRes)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	var req entity.UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err = c.userService.UpdateUser(id, &req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "User updated successfully", nil)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}

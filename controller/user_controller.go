package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/helper"
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
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := c.userService.RegisterUser(&req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusCreated, "User registered successfully", user)
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var req entity.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	userRes, err := c.userService.LoginUser(&req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to login", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Login successful", userRes)
}

func (c *UserController) FindUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	userRes, err := c.userService.FindUserByID(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "User data retrieved successfully", userRes)
}

func (c *UserController) FindAllUsers(ctx *gin.Context) {
	var paginationReq helper.PaginationRequest
	if err := ctx.ShouldBindQuery(&paginationReq); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid pagination parameters", err)
		return
	}

	// Set default values if page or limit is not provided
	if paginationReq.Page == 0 {
		paginationReq.Page = 1 // Default page is 1
	}
	if paginationReq.Limit == 0 {
		paginationReq.Limit = 10 // Default limit is 10
	}

	usersRes, err := c.userService.FindAllUsers()
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}

	var data []interface{}
	for _, user := range usersRes {
		data = append(data, user)
	}

	paginatedResponse := helper.Paginate(data, paginationReq.Page, paginationReq.Limit)

	helper.SendSuccessResponse(ctx, http.StatusOK, "Users data retrieved successfully", paginatedResponse)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	var req entity.UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err = c.userService.UpdateUser(id, &req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "User updated successfully", nil)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}

func (c *UserController) RegisterAsAdmin(ctx *gin.Context) {
	var req entity.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	admin, err := c.userService.RegisterAsAdmin(&req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to register as admin", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusCreated, "Admin registered successfully", admin)
}

func (c *UserController) GetUserReport(ctx *gin.Context) {
	// Ambil start_date dan end_date dari query string (opsional)
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	var startDate, endDate time.Time
	var err error

	// Parse start_date jika ada
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid start_date format (YYYY-MM-DD)", err)
			return
		}
	}

	// Parse end_date jika ada
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid end_date format (YYYY-MM-DD)", err)
			return
		}
	}

	// Pastikan start_date tidak lebih besar dari end_date jika keduanya ada
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "start_date must be before or equal to end_date", nil)
		return
	}

	// Panggil service untuk menghasilkan laporan
	report, err := c.userService.GetUserReport(startDate, endDate)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate user report", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "User report generated successfully", report)
}

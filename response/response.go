package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func SendSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}

func SendErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	}
	ctx.JSON(statusCode, response)
}

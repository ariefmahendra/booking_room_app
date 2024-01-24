package common

import (
	"booking-room/shared/shared_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, &shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendSingleResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, &shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &shared_model.Status{
		Code:    code,
		Message: message,
	})
}

func SendNoContentResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func SendListResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, &shared_model.ListResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

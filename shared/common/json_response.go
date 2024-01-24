package common

import (
	"booking-room/shared/shared_model"
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, shared_model.Status{
		Code:    code,
		Message: message,
	})
}

func SendSuccessResponse(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    code,
			Message: "success",
		},
		Data: data,
	})
}

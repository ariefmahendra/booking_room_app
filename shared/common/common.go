package common

import (
	"net/http"

	"booking-room/shared/shared_model"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, message string){
	ctx.JSON(code, shared_model.Status{
		Code: code,
		Message: message,
	})
}

func SendPagedResponse(ctx *gin.Context, data interface{}, paging shared_model.Paging, message string) {
	ctx.JSON(http.StatusOK, shared_model.PagedResponse{
		Status: shared_model.Status{
			Code: http.StatusOK,
			Message: message,
		},
		Data: data,
		Paging: paging,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
    ctx.JSON(http.StatusOK, shared_model.SingleResponse{
        Code:    http.StatusOK,
        Message: message,
        Data:    data,
    })
}


func SendCreatedResponse(ctx *gin.Context, data interface{}, message string){
	ctx.JSON(http.StatusOK, shared_model.SingleResponse{
		Code:    http.StatusCreated,
        Message: message,
        Data:    data,
	})
}
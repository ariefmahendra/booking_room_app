package common

import (
	"net/http"

	"booking-room/shared/shared_model"

	"github.com/gin-gonic/gin"
)

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
        Data:    data,
    })
}


func SendCreatedResponse(ctx *gin.Context, data interface{}, message string){
	ctx.JSON(http.StatusOK, shared_model.SingleResponse{
        Data:    data,
	})
}

func SendCreateResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, &shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}
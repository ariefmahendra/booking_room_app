package common

import (
	"booking-room/shared/shared_model"
	"net/http"

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

func SendSuccessPagedResponse(ctx *gin.Context, code int, data any, paging shared_model.Paging) {
	ctx.JSON(code, shared_model.PagedResponse{
		Status: shared_model.Status{
			Code:    code,
			Message: "success",
		},
		Data:   data,
		Paging: paging,
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

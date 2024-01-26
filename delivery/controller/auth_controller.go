package controller

import (
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthControllerImpl struct {
	authUC usecase.AuthUC
	rg     *gin.RouterGroup
}

func NewAuthController(authUC usecase.AuthUC, rg *gin.RouterGroup) *AuthControllerImpl {
	return &AuthControllerImpl{authUC: authUC, rg: rg}
}

func (a *AuthControllerImpl) Route() {
	a.rg.POST("/login", a.LoginController)
	a.rg.POST("/register", a.RegisterController)
}

func (a *AuthControllerImpl) LoginController(ctx *gin.Context) {
	var payload dto.AuthRequest
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Printf("failed to bind json : %v", err)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "bad request")
		return
	}

	response, err := a.authUC.Login(payload)
	if err != nil {
		log.Printf("failed to login : %v", err)
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, response)
}

func (a *AuthControllerImpl) RegisterController(ctx *gin.Context) {
	var request dto.EmployeeCreateRequest
	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("failed to bind json : %v", err)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "bad request")
		return
	}

	response, err := a.authUC.Register(request)
	if err != nil {
		log.Printf("failed to register : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusCreated, response)
}

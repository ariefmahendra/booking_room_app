package middleware

import (
	"booking-room/shared/common"
	"booking-room/shared/service"
	"booking-room/shared/shared_model"
	"booking-room/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	authUC     usecase.AuthUC
	jwtService service.JwtService
}

func NewMiddleware(authUC usecase.AuthUC, jwtService service.JwtService) *Middleware {
	return &Middleware{authUC: authUC, jwtService: jwtService}
}

func (m *Middleware) NewAuth(ctx *gin.Context) {
	if !(ctx.FullPath() == "/api/v1/auth/login") {
		fullToken := ctx.GetHeader("Authorization")

		if fullToken == "" {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		tokens := strings.Split(fullToken, " ")

		if len(tokens) != 2 {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "token length invalid")
			ctx.Abort()
			return
		}

		if tokens[0] != "Bearer" {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "bearer token invalid")
			ctx.Abort()
			return
		}

		customClaims, err := m.jwtService.ValidateToken(tokens[1])
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "token invalid")
			ctx.Abort()
			return
		}

		role := strings.ToUpper(customClaims.Role)
		if (customClaims == nil) || (role != "ADMIN" && role != "EMPLOYEE" && role != "GA") {
			common.SendErrorResponse(ctx, http.StatusForbidden, "role forbidden")
			ctx.Abort()
			return
		}

		ctx.Set("ctx", customClaims)
		ctx.Next()
	}
}

func (m *Middleware) GetUser(ctx *gin.Context) *shared_model.CustomClaims {
	user, exists := ctx.Get("ctx")
	if !exists {
		return nil
	}
	claims := user.(*shared_model.CustomClaims)
	return claims
}

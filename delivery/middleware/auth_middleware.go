package middleware

import (
	"booking-room/shared/common"
	"booking-room/shared/service"
	"booking-room/shared/shared_model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Middleware struct {
	jwtService service.JwtService
}

func NewMiddleware(jwtService service.JwtService) *Middleware {
	return &Middleware{jwtService: jwtService}
}

func (m *Middleware) NewAuth(ctx *gin.Context) {
	if !(ctx.FullPath() == "/api/v1/auth/login" || ctx.FullPath() == "/api/v1/auth/register") {
		fullToken := ctx.GetHeader("Authorization")

		if fullToken == "" {
			log.Printf("token is nil value : %v", fullToken)
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		tokens := strings.Split(fullToken, " ")

		if len(tokens) != 2 {
			log.Printf("token length invalid : %v", tokens)
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "token length invalid")
			ctx.Abort()
			return
		}

		if tokens[0] != "Bearer" {
			log.Printf("key is not bearer : %v", tokens[0])
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "bearer token invalid")
			ctx.Abort()
			return
		}

		customClaims, err := m.jwtService.ValidateToken(tokens[1])
		if err != nil {
			log.Printf("failed to validate token : %v", err)
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "token invalid")
			ctx.Abort()
			return
		}

		role := strings.ToUpper(customClaims.Role)
		if (customClaims == nil) || (role != "ADMIN" && role != "EMPLOYEE" && role != "GA") {
			log.Printf("role forbidden : %v", customClaims)
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

package common

import (
	"booking-room/shared/shared_model"
	"strings"
)

func AuthorizationAdmin(claims *shared_model.CustomClaims) bool {
	if claims == nil || strings.ToUpper(claims.Role) != "ADMIN" {
		return false
	}

	return false
}

func AuthorizationGaAdmin(claims *shared_model.CustomClaims) bool {
	if claims == nil || strings.ToUpper(claims.Role) != "ADMIN" || strings.ToUpper(claims.Role) != "GA" {
		return false
	}

	return false
}

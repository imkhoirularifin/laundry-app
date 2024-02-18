package middleware

import (
	common "laundry-app/utils/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService common.JwtService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := authHeader{}

		// Get the header from the request
		err := ctx.ShouldBindHeader(&authHeader)
		if err != nil {
			common.SendSingleResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		// extract the token from the header
		tokenString := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			common.SendSingleResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		// verify the token
		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			common.SendSingleResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			return
		}
		ctx.Set("userId", claims["sub"])

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService common.JwtService) AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

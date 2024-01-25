package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/shared/service"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader AuthHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			log.Printf("RequireToken: Error binding header: %v \n", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenHeader := strings.TrimPrefix(authHeader.AuthorizationHeader, "Bearer ")
		if tokenHeader == "" {
			log.Println("RequireToken: Missing token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken: Error parsing token: %v \n", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("customer", claims["customerId"])

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/shared/service"
)

type AuthMiddleware interface {
	RequireToken(roles string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader AuthHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			log.Printf("RequireToken: Error binding header: %v \n", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Error binding header"})
			ctx.Abort()
			return
		}

		tokenHeader := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenHeader == "" {
			log.Println("RequireToken: Missing token")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
			ctx.Abort()
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken: Error parsing token: %v \n", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			ctx.Abort()
			return
		}

		if _, ok := claims["customerId"]; !ok {
			log.Println("RequireToken: Missing 'customerId' claim in token")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing 'customerId' claim in token"})
			ctx.Abort()
			return
		}

		// a.jwtService.SaveCustomerSession(ctx, claims)

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

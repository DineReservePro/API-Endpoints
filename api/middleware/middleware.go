package middleware

import (
	"api-gateway/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)


func JWTMiddleware()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth != ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"ERROR" : "Authorization header required",
			})
			return
		}

		valid,err := token.ValidateToken(auth)
		if err != nil || !valid{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"ERROR" : "Invalid token",
			})
			return
		}

		claims,err := token.ExtractClaims(auth)
		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"ERROR" : "Invalid Token Claims",
			})
			return
		}
		ctx.Set("claims",claims)
		ctx.Next()
	}
}
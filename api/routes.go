package api

import (
	"api-gateway/api/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.DELETE("/logout/:user-id", handle.LogoutUserHandler)
		auth.GET("/profile", handle.GetUserProfileHandler)
		auth.PUT("/profile/:user-id", handle.UpdateUserProfile)
	}

	payment := router.Group("/payments")
	{
		payment.POST("/", handle.CreatePaymentHandler)
		payment.GET("/:payment-id", handle.GetPaymentHandler)
		payment.PUT("/:payment-id", handle.UpdatePaymentHandler)
	}

	return router
}

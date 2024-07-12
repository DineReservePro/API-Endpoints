package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample server for Auth Service.
// @host localhost:8081
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.LoggerMiddleware())

	auth := router.Group("/auth")
	{
		auth.DELETE("/logout/:user-id", handle.LogoutUserHandler)
		auth.GET("/profile", handle.GetUserProfileHandler)
		auth.PUT("/profile/:user-id", handle.UpdateMenuItemHandler)
	}

	restaurant := router.Group("/restaurant")
	{
		restaurant.POST("/",handle.CreateRestaurantHandler)
		restaurant.GET("/",handle.ListRestaurantsHandler)
		restaurant.GET("/:restaurant-id",handle.GetRestaurantHandler)
		restaurant.PUT("/:restaurant-id",handle.UpdateRestaurantHandler)
		restaurant.DELETE("/:restaurant_id",handle.DeleteRestaurantHandler)
		
	}

	menu := router.Group("/menu")
	{
		menu.POST("/",handle.CreateMenuItemHandler)
		menu.GET("/",handle.ListMenuItemsHandler)
		menu.GET("/:menu-id",handle.GetMenuItemHandler)
		menu.PUT("/:menu-id",handle.UpdateMenuItemHandler)
		menu.DELETE("/:menu-id",handle.DeleteMenuItemHandler)
	}

	payment := router.Group("/payments")
	{
		payment.POST("/", handle.CreatePaymentHandler)
		payment.GET("/:payment-id", handle.GetPaymentHandler)
		payment.PUT("/:payment-id", handle.UpdatePaymentHandler)
	}

	reservation := router.GET("/reservations")
	{
		reservation.POST("/",handle.CreateReservationHandler)
		reservation.GET("/",handle.ListReservationHandler)
		reservation.GET("/:reservation-id",handle.GetReservationHandler)
		reservation.PUT("/:reservation-id",handle.UpdateReservationHandler)
		reservation.DELETE("/:reservation-id",handle.DeleteReservationHandler)
		reservation.POST("/check",handle.CheckReservationHandler)
		reservation.POST("/order",handle.OrderMealsHandler)
	}

	return router
}

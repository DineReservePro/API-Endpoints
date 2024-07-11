package handler

import (
	pb "api-gateway/generated/reservation_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRestaurantHandler(ctx *gin.Context) {
	var restaurant pb.CreateRestaurantRequest

	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Reservation.CreateRestaurant(ctx, &restaurant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListRestaurantsHandler(ctx *gin.Context) {
	var filter pb.ListRestaurantsRequest

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


}

func (h *Handler) GetRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("restaurant-id")

	resp, err := h.Reservation.GetRestaurant(ctx, &pb.GetRestaurantRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("restaurant-id")
	var restaurant pb.UpdateRestaurantRequest

	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	restaurant.Id = id

	resp, err := h.Reservation.UpdateRestaurant(ctx, &restaurant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("restaurant-id")
	
	resp, err := h.Reservation.DeleteRestaurant(ctx, &pb.DeleteRestaurantRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

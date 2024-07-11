package handler

import (
	pb "api-gateway/generated/reservation_service"
	"net/http"

	"github.com/gin-gonic/gin"
)
//@Summary Create Reataurant
//@Description  Create a new Restaurant
//@Tags Restaurant
//@Accept json
//@Security ApiKeyAuth
//@Produce json
//@Param Restaurant body reservation_service.CreateRestaurantRequest true "Create Restaurant"
//@Success 200 {object} reservation_service.CreateRestaurantResponse
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/restaurant [post]
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

//@Summary List Restaurant
//@Description List all Restaurant
//@Tags Restaurant
//@Accept json
//@Security ApiKeyAuth
//@Produce json
//@Succes 200 {object} reservation_service.ListRestaurantResponse
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/restaurant [get]
func (h *Handler) ListRestaurantsHandler(ctx *gin.Context) {
	var filter pb.ListRestaurantsRequest

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if filter.Limit == 0{
		filter.Limit = 10
	}

	res,err := h.Reservation.ListRestaurants(ctx,&filter)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"ERROR" : err.Error(),
		})
		return
	}
	ctx.JSON(200,res)
}

//@Summary Get Restaurant
//@Description Get Restaurant By ID
//@Tags Restaurant
//@Accept json
//@Security ApiKeyAuth
//@Produce json
//@Param restaurant-id path string true "Restaurant ID"
//@Success 200 {object} reservation_service.GetRestaurantResponse
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/restaurant/{restaurant-id} [get]
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

//@Summary Update Restaurant
//@Description Update an existing restaurant 
//@Tags Restaurant
//@Accept json
//@Security ApiKeyAuth
//@Produce json
//@Param restaurant-id path string true "Restaurant ID"
//@Param Restaurant body reservation_service.UpdateRestaurantRequest true "Update Restaurant"
//@Success 200 {object} reservation_service.UpdateRestaurantResponse
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/restaurant{restaurant-id} [put]
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
//@Summary Delete Restaurant
//@Description Delete a specific restaurant by ID
//@Tags Restaurant
//@Accept json
//@Security ApiKeyAuth
//@Produce json
//@Param restaurant-id path string true "Restaurant ID"
//@Success 200 {object} reservation_service.DeleteRestaurantResponse
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Interval Server Error"
//@Router /api/restaurant/{restaurant-id} [delete]
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

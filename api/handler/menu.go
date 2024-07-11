package handler

import (
	pb "api-gateway/generated/reservation_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//@Summary Create MeniItem
//@Description Create a new MenuItem
//@Tags MenuItem
//@Accept json
//@Security BearerAuth
//@Produce json
//@Param MenuItem body reservation_service.CreateMenuItemRequest true "Create MenuItem"
//@Succes 200 {object} reservation_service.CreateMenuItemResponce
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal server Error"
//@Router /api/menu [post]
func (h *Handler) CreateMenuItemHandler(ctx *gin.Context) {
	menuItem := pb.CreateMenuItemRequest{}

	if err := ctx.ShouldBindJSON(&menuItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Reservation.CreateMenuItem(ctx, &menuItem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


//@Summary List MenuItem
//@Description List all MenuItem
//@Tags MenuItem
//@Accept json
//@Security BearerAuth
//@Produce json
//@Success 200 {object} reservation_service.ListMenuItemRequest
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/menu [get]
func (h *Handler) ListMenuItemsHandler(ctx *gin.Context) {
	var filter pb.ListMenuItemsRequest

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}

}
//@Summary Get MenuItem
//@Description Get a specific MenuItem by ID
//@Tags MenuItem
//@Accept json
//@Security BearerAuth
//@Produce json
//@Param menu-id path string true "Menu ID"
//@Success 200 {object} reservation_service.GetMenuItemResponce
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/menu/{menu-id} [get]
func (h *Handler) GetMenuItemHandler(ctx *gin.Context) {
	id := ctx.Param("menu-id")

	resp, err := h.Reservation.GetMenuItem(ctx, &pb.GetMenuItemRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

//@Summary Update MenuItem
//@Descripton Update an existing reservation
//@Tags MenuItem
//@Accept json
//@Security BearerAuth
//@Produce json
//@Param menu-id path string true "Menu ID"
//@Param Menu body reservation_service.UpdateMenuItemRequest true "Update MenuItem"
//@Success 200 {object} reservation_service.UpdateMenuItemResponce
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/menu/{menu-id} [put]
func (h *Handler) UpdateMenuItemHandler(ctx *gin.Context) {
	id := ctx.Param("menu-id")

	var menuItem pb.UpdateMenuItemRequest

	if err := ctx.ShouldBindJSON(&menuItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	menuItem.Id = id

	resp, err := h.Reservation.UpdateMenuItem(ctx, &menuItem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

//@Summary Delete MenuItem
//@Description Delete a specific MenuItem by ID
//@Tags MenuItem
//@Accept json
//@Security BearerAuth
//@Produce json
//@Param menu-id path string true "Menu ID"
//@Success 200 {object} reservation_service.DeleteMenuItemResponce
//@Failure 400 {object} string "Bad Request"
//@Failure 500 {object} string "Internal Server Error"
//@Router /api/menu/{menu-id} [delete]
func (h *Handler) DeleteMenuItemHandler(ctx *gin.Context) {
	id := ctx.Param("menu-id")

	resp, err := h.Reservation.DeleteMenuItem(ctx, &pb.DeleteMenuItemRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}	
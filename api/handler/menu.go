package handler

import (
	pb "api-gateway/generated/reservation_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) ListMenuItemsHandler(ctx *gin.Context) {
	var filter pb.ListMenuItemsRequest

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}

}

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
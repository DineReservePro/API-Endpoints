package handler

import (
	pb "api-gateway/generated/reservation_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateReservationHandler(ctx *gin.Context) {
	reservation := pb.CreateReservationRequest{}

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Reservation.CreateReservation(ctx, &reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Reservation)
}

func (h *Handler) ListReservationHandler(ctx *gin.Context) {
	resp, err := h.Reservation.ListReservations(ctx, &pb.ListReservationsRequest{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Reservations)
}

func (h *Handler) GetReservationHandler(ctx *gin.Context) {
	id := ctx.Param("reservation-id")

	resp, err := h.Reservation.GetReservation(ctx, &pb.GetReservationRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Reservation)
}

func (h *Handler) UpdateReservationHandler(ctx *gin.Context) {
	id := ctx.Param("reservation-id")

	reservation := pb.UpdateReservationRequest{}

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	reservation.Id = id

	resp, err := h.Reservation.UpdateReservation(ctx, &reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Reservation)
}

func (h *Handler) DeleteReservationHandler(ctx *gin.Context) {
	id := ctx.Param("reservation-id")

	reservation := pb.DeleteReservationRequest{}

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	reservation.Id = id

	resp, err := h.Reservation.DeleteReservation(ctx, &reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Message)
}
func (h *Handler) CheckReservationHandler(ctx *gin.Context) {
	checkReq := pb.CheckReservationRequest{}

	if err := ctx.ShouldBindJSON(&checkReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Reservation.CheckReservation(ctx, &checkReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderMealsHandler(ctx *gin.Context) {
	orderReq := pb.OrderMealsRequest{}

	if err := ctx.ShouldBindJSON(&orderReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Reservation.OrderMeals(ctx, &orderReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

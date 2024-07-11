package handler

import (
	pb "api-gateway/generated/reservation_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateReservationHandler handles the creation of a new reservation.
// @Summary Create Reservation
// @Description Create a new reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Reservation body pb.CreateReservationRequest true "Create Reservation"
// @Success 200 {object} pb.CreateReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations [post]
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

// ListReservationHandler lists all reservations.
// @Summary List Reservations
// @Description List all reservations
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {object} pb.ListReservationsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations [get]
func (h *Handler) ListReservationHandler(ctx *gin.Context) {
	var filter pb.ListReservationsRequest
	if err := ctx.ShouldBindQuery(&filter);err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if filter.Limit == 0{
		filter.Limit = 10
	}
	resp, err := h.Reservation.ListReservations(ctx, &filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetReservationHandler retrieves a specific reservation by ID.
// @Summary Get Reservation
// @Description Get a specific reservation by ID
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Success 200 {object} pb.GetReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [get]
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

// UpdateReservationHandler updates an existing reservation.
// @Summary Update Reservation
// @Description Update an existing reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Param Reservation body pb.UpdateReservationRequest true "Update Reservation"
// @Success 200 {object} pb.UpdateReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [put]
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

// DeleteReservationHandler deletes a specific reservation by ID.
// @Summary Delete Reservation
// @Description Delete a specific reservation by ID
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Success 200 {object} pb.DeleteReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [delete]
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

// CheckReservationHandler checks the status of a reservation.
// @Summary Check Reservation
// @Description Check the status of a reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Check body pb.CheckReservationRequest true "Check Reservation"
// @Success 200 {object} pb.CheckReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/check [post]
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

// OrderMealsHandler handles ordering meals for a reservation.
// @Summary Order Meals
// @Description Order meals for a reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Order body pb.OrderMealsRequest true "Order Meals"
// @Success 200 {object} pb.OrderMealsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/order [post]
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

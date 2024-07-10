package handler

import (
	pb "api-gateway/generated/payment_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
	payment := pb.CreatePaymentRequest{}

	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Payment.CreatePayment(ctx, &payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Payment)
}

func (h *Handler) GetPaymentHandler(ctx *gin.Context) {
	id := ctx.Param("payment-id")

	resp, err := h.Payment.GetPayment(ctx, &pb.GetPaymentRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Payment)
}


func (h *Handler) UpdatePaymentHandler(ctx *gin.Context) {
	id := ctx.Param("payment-id")

	payment := pb.UpdatePaymentRequest{}

	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	payment.Id = id

	resp, err := h.Payment.UpdatePayment(ctx, &payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Payment)
}
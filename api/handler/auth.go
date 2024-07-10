package handler

import (
	pb "api-gateway/generated/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterHandler(ctx *gin.Context) {
	user := pb.RegisterRequest{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Auth.RegisterUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": resp.Message,
	})
}

func (h *Handler) LoginHandler(ctx *gin.Context) {
	login := pb.LoginRequest{}

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Auth.LoginUser(ctx, &login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) LogoutUserHandler(ctx *gin.Context) {
	id := ctx.Param("user-id")

	resp, err := h.Auth.LogoutUser(ctx, &pb.LogoutRequest{UserId: id})
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserProfileHandler(ctx *gin.Context) {
	username := ctx.Param("username")

	resp, err := h.Auth.GetUserProfile(ctx, &pb.GetUserProfileRequest{Username: username})
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
	id := ctx.Param("user-id")

	profile := pb.UpdateUserProfileRequest{}
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	profile.UserId = id

	resp, err := h.Auth.UpdateUserProfile(ctx, &profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
package handler

import (
	pb "api-gateway/generated/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)


// LogoutUserHandler handles the logout of a user.
// @Summary Logout User
// @Description Logout the authenticated user
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} pb.LogoutResponse
// @Failure 400 {object} string "Bad Request"
// @Router /api/auth/logout/{user-id} [post]
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

// GetUserProfileHandler retrieves the user profile.
// @Summary Get User Profile
// @Description Get profile of the authenticated user
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pb.GetUserProfileResponse
// @Failure 400 {object} string "Bad Request"
// @Router /api/auth/profile/{username} [get]
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

// UpdateUserProfile updates the user profile.
// @Summary Update User Profile
// @Description Update the profile of the authenticated user
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param user-id path string true "User ID"
// @Param profile body pb.UpdateUserProfileRequest true "Profile"
// @Success 200 {object} pb.UpdateUserProfileResponse
// @Failure 400 {object} string "Bad Request"
// @Router /api/auth/profile/{user-id} [put]
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
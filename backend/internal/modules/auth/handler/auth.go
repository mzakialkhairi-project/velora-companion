// Package handler provides authentication HTTP handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/service"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// AuthHandler handles HTTP requests for authentication operations
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 201 {object} shared.Response
// @Failure 400 {object} shared.Response
// @Failure 409 {object} shared.Response
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, apperrors.ErrAlreadyExists) {
			shared.ErrorResponse(c, http.StatusConflict, "Email already exists")
			return
		}
		if errors.Is(err, apperrors.ErrValidation) {
			shared.ErrorResponse(c, http.StatusBadRequest, "Validation failed")
			return
		}
		shared.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Return success response
	shared.Success(c, http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// Login handles user login
// Login godoc
// @Summary Login a user
// @Description Authenticate user and return user info
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} shared.Response
// @Failure 400 {object} shared.Response
// @Failure 401 {object} shared.Response
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) || errors.Is(err, apperrors.ErrUnauthorized) {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		if errors.Is(err, apperrors.ErrValidation) {
			shared.ErrorResponse(c, http.StatusBadRequest, "Validation failed")
			return
		}
		shared.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	shared.Success(c, http.StatusOK, result)
}

// Logout handles user logout
// Logout godoc
// @Summary Logout a user
// @Description Invalidate refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} shared.Response
// @Failure 401 {object} shared.Response
// @Router /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := h.authService.Logout(c.Request.Context(), req.RefreshToken)
	if err != nil {
		if errors.Is(err, apperrors.ErrValidation) {
			shared.ErrorResponse(c, http.StatusBadRequest, "Refresh token required")
			return
		}
		shared.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	shared.Success(c, http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// RefreshToken handles token refresh
// RefreshToken godoc
// @Summary Refresh access token
// @Description Get new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} shared.Response
// @Failure 401 {object} shared.Response
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.authService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, apperrors.ErrValidation) {
			shared.ErrorResponse(c, http.StatusBadRequest, "Refresh token required")
			return
		}
		if errors.Is(err, apperrors.ErrUnauthorized) {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired refresh token")
			return
		}
		shared.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	shared.Success(c, http.StatusOK, result)
}

// Package handler provides workspace HTTP handlers.
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/mapper"
	"github.com/mzakiaklhairi/velora/internal/ai/service"
	"github.com/mzakiaklhairi/velora/internal/infrastructure"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// WorkspaceHandler handles HTTP requests for workspace operations
type WorkspaceHandler struct {
	workspaceService service.WorkspaceService
}

// NewWorkspaceHandler creates a new WorkspaceHandler
func NewWorkspaceHandler(workspaceService service.WorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{
		workspaceService: workspaceService,
	}
}

// Create handles workspace creation
// Create godoc
// @Summary Create a new workspace
// @Description Create a new workspace
// @Tags workspaces
// @Accept json
// @Produce json
// @Param request body dto.CreateWorkspaceRequest true "Create workspace request"
// @Success 201 {object} dto.WorkspaceResponse
// @Failure 400 {object} map[string]string
// @Router /workspaces [post]
func (h *WorkspaceHandler) Create(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	var req dto.CreateWorkspaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	workspace, err := h.workspaceService.Create(c.Request.Context(), userID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusCreated, mapper.ToResponse(workspace))
}

// Update handles workspace update
// Update godoc
// @Summary Update a workspace
// @Description Update an existing workspace
// @Tags workspaces
// @Accept json
// @Produce json
// @Param id path int true "Workspace ID"
// @Param request body dto.UpdateWorkspaceRequest true "Update workspace request"
// @Success 200 {object} dto.WorkspaceResponse
// @Failure 400 {object} map[string]string
// @Router /workspaces/{id} [put]
func (h *WorkspaceHandler) Update(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	var req dto.UpdateWorkspaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	workspace, err := h.workspaceService.Update(c.Request.Context(), userID, id, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, mapper.ToResponse(workspace))
}

// Delete handles workspace deletion
// Delete godoc
// @Summary Delete a workspace
// @Description Delete a workspace (soft delete)
// @Tags workspaces
// @Param id path int true "Workspace ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /workspaces/{id} [delete]
func (h *WorkspaceHandler) Delete(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	if err := h.workspaceService.Delete(c.Request.Context(), userID, id); err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusNoContent, nil)
}

// GetByID handles getting a workspace by ID
// GetByID godoc
// @Summary Get a workspace by ID
// @Description Get a workspace by its ID
// @Tags workspaces
// @Produce json
// @Param id path int true "Workspace ID"
// @Success 200 {object} dto.WorkspaceResponse
// @Failure 404 {object} map[string]string
// @Router /workspaces/{id} [get]
func (h *WorkspaceHandler) GetByID(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	workspace, err := h.workspaceService.GetByID(c.Request.Context(), userID, id)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, mapper.ToResponse(workspace))
}

// List handles listing workspaces
// List godoc
// @Summary List workspaces
// @Description Get all workspaces for the authenticated user
// @Tags workspaces
// @Produce json
// @Success 200 {object} dto.WorkspaceListResponse
// @Router /workspaces [get]
func (h *WorkspaceHandler) List(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaces, err := h.workspaceService.List(c.Request.Context(), userID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, dto.WorkspaceListResponse{
		Workspaces: mapper.ToResponseList(workspaces),
		Total:      len(workspaces),
	})
}

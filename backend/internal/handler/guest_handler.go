package handler

import (
	"net/http"
	"rsvp-system/internal/dto"
	"rsvp-system/internal/models"
	"rsvp-system/internal/service"

	"github.com/gin-gonic/gin"
)

type GuestHandler struct {
	service service.GuestService
}

func NewGuestHandler(s service.GuestService) *GuestHandler {
	return &GuestHandler{service: s}
}

func (h *GuestHandler) InviteGuest(c *gin.Context) {

	var req dto.InviteGuestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guest := models.Guest{
		Name:    req.Name,
		Email:   req.Email,
		EventID: req.EventID,
	}

	if err := h.service.InviteGuest(&guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.GuestResponse{
		ID:     guest.ID,
		Name:   guest.Name,
		Email:  guest.Email,
		Status: guest.Status,
	}

	c.JSON(http.StatusCreated, res)
}

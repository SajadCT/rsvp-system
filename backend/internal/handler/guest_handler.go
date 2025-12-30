package handler

import (
	"net/http"
	"strconv"

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

func (h *GuestHandler) GetGuests(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("event_id"))
	guests, _ := h.service.GetGuestsForEvent(uint(id))

	response := []dto.GuestResponse{}
	for _, g := range guests {
		response = append(response, dto.GuestResponse{
			ID:     g.ID,
			Name:   g.Name,
			Email:  g.Email,
			Status: g.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *GuestHandler) UpdateRSVP(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.UpdateRSVPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status"})
		return
	}

	if err := h.service.RSVPGuest(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "RSVP updated successfully"})
}

func (h *GuestHandler) GetGuestDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	guest, err := h.service.GetGuestDetails(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest invitation not found"})
		return
	}

	res := dto.GuestDetailResponse{
		ID:            guest.ID,
		Name:          guest.Name,
		Status:        guest.Status,
		EventTitle:    guest.Event.Title,
		EventDate:     guest.Event.Date,
		EventLocation: guest.Event.Location,
	}

	c.JSON(http.StatusOK, res)
}

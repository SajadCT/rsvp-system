package handler

import (
	"net/http"
	"strconv"

	"rsvp-system/internal/dto"
	"rsvp-system/internal/models"
	"rsvp-system/internal/service"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service service.EventService
}

func NewEventHandler(s service.EventService) *EventHandler {
	return &EventHandler{service: s}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
		Location:    req.Location,
	}

	if err := h.service.CreateEvent(&event, userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		Location:    event.Location,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	events, _ := h.service.GetAllEvents(userID.(uint))

	var response []dto.EventResponse
	for _, e := range events {
		response = append(response, dto.EventResponse{
			ID:          e.ID,
			Title:       e.Title,
			Description: e.Description,
			Date:        e.Date,
			Location:    e.Location,
		})
	}

	if response == nil {
		response = []dto.EventResponse{}
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteEvent(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func (h *EventHandler) GetStats(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	statsMap, _ := h.service.GetEventStats(uint(id))

	response := dto.EventStatsResponse{
		Total:    statsMap["total"],
		Accepted: statsMap["accepted"],
		Declined: statsMap["declined"],
		Pending:  statsMap["pending"],
	}

	c.JSON(http.StatusOK, response)
}

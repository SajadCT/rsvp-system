package handler

import (
	"net/http"
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

	if err := h.service.CreateEvent(&event); err != nil {
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
	events, _ := h.service.GetAllEvents()

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

	c.JSON(http.StatusOK, gin.H{"data": response})
}

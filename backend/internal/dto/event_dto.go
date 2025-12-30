package dto

type CreateEventRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Date        string `json:"date" binding:"required"`
	Location    string `json:"location" binding:"required"`
}

type EventResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}

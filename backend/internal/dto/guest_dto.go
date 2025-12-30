package dto

type InviteGuestRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	EventID uint   `json:"event_id" binding:"required"`
}

type UpdateRSVPRequest struct {
	Status string `json:"status" binding:"required,oneof=Accepted Declined Pending"`
}

type GuestResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type GuestDetailResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	EventTitle    string `json:"event_title"`
	EventDate     string `json:"event_date"`
	EventLocation string `json:"event_location"`
}

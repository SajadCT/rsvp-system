package routes

import (
	"rsvp-system/internal/handler"
	"rsvp-system/internal/middleware"
	"rsvp-system/internal/repository"
	"rsvp-system/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	// 2. Repositories (Data Layer)
	eventRepo := repository.NewEventRepository(db)
	guestRepo := repository.NewGuestRepository(db)

	// 3. Services (Business Logic Layer)
	eventService := service.NewEventService(eventRepo, guestRepo)
	guestService := service.NewGuestService(guestRepo, eventRepo)

	// 4. Handlers (HTTP Layer)
	eventHandler := handler.NewEventHandler(eventService)
	guestHandler := handler.NewGuestHandler(guestService)

	r.GET("/hai", handler.Hello)

	api := r.Group("/api")
	{
		// Event Routes
		api.POST("/events", eventHandler.CreateEvent)
		api.GET("/events", eventHandler.GetEvents)
		// api.DELETE("/events/:id", eventHandler.DeleteEvent)
		// api.GET("/events/:id/stats", eventHandler.GetStats)

		// Guest Routes
		api.POST("/invite", guestHandler.InviteGuest)
		// api.GET("/guests/:event_id", guestHandler.GetGuests)s
		// api.PAsTCH("/guests/:id/rsvp", guestHandler.UpdateRSVP)
		// api.GET("/guests/details/:id", guestHandler.GetGuestDetails)

		// Export (Optional)
		// api.GET("/events/:id/export", eventHandler.ExportGuests)
	}

	return r
}

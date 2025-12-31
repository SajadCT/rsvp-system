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
	r.Use(middleware.CORS())

	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)
	guestRepo := repository.NewGuestRepository(db)

	authService := service.NewAuthService(userRepo)
	eventService := service.NewEventService(eventRepo, guestRepo)
	guestService := service.NewGuestService(guestRepo, eventRepo)

	authHandler := handler.NewAuthHandler(authService)
	eventHandler := handler.NewEventHandler(eventService)
	guestHandler := handler.NewGuestHandler(guestService)

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		api.GET("/guests/details/:id", guestHandler.GetGuestDetails)
		api.PATCH("/guests/:id/rsvp", guestHandler.UpdateRSVP)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/events", eventHandler.CreateEvent)
			protected.GET("/events", eventHandler.GetEvents)
			protected.DELETE("/events/:id", eventHandler.DeleteEvent)
			protected.GET("/events/:id/stats", eventHandler.GetStats)

			protected.POST("/invite", guestHandler.InviteGuest)
			protected.GET("/guests/:event_id", guestHandler.GetGuests)
		}
	}

	return r
}

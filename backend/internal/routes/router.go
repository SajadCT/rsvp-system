package routes

import (
	"rsvp-system/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/hello", handler.Hello)

	return r
}

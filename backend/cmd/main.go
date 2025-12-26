package main

import (
	"log"
	"rsvp-system/internal/database"
	"rsvp-system/internal/routes"
)

func main() {

	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	r := routes.SetUpRoutes(db)
	log.Println("Server running on http://localhost:8080")

	Ginerr := r.Run(":8080")
	if Ginerr != nil {
		log.Fatal(err)
	}

}

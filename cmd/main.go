package main

import (
	"log"
	"net/http"

	"github.com/arpitproject2/meeting_schedular/config"
	"github.com/arpitproject2/meeting_schedular/internal/controllers"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize routes
	http.HandleFunc("/meetings", controllers.CreateMeeting)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

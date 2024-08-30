package main

import (
	"log"
	"meeting_scheduler/controllers"
	"meeting_scheduler/internal/db"
	"net/http"
)

func main() {
	// Simple route handler
	db.Connect()
	http.HandleFunc("/meeting", controllers.GetMeeting)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

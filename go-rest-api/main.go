package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models" // Import the models package

	"github.com/gin-gonic/gin"
)


func  main() {
	db.InitDB()	// Initialize the database connection

	server := gin.Default()	// Create an http server with default middleware (Logger and Recovery)
	
	// First parameter is the path, second parameter is the function to be executed when the path is hit
	server.GET("/events", getEvents )	
	server.POST("/events", createEvent)
	
	server.Run(":8080");	// localhost:8080

}

// getEvents handles the retrieval of all events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events) 
}

// createEvent handles the creation of a new event
// context is the gin context that contains the request and response objects
// It binds the JSON request body to the event struct and saves it
// If there is an error, it returns a bad request response
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Call Save method to store the event in the database
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

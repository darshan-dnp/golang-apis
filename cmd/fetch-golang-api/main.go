package main

import (
	"fetch-golang-api/api"
	"fetch-golang-api/api/handler"
	"fetch-golang-api/db"
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/receipts/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Setup Database
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	pointsVariables := cache.NewCache(database)

	// Setup service
	receiptService := service.NewService(database, pointsVariables)
	handler.SetReceiptService(receiptService)

	// Routes
	routes := gin.Default()
	api.SetupRoutes(routes)

	routes.Run()
}

package api

import (
	"fetch-golang-api/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/receipts/:id/points", handler.GetReceiptPoints)
	r.POST("/receipts/process", handler.CreateReceipts)
}

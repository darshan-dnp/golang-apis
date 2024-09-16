package handler

import (
	"fetch-golang-api/internal/model"
	"fetch-golang-api/internal/receipts"
	"fetch-golang-api/internal/receipts/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptService *service.Service

func SetReceiptService(service *service.Service) {
	receiptService = service
}

func CreateReceipts(c *gin.Context) {
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := receipt.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReceipt, err := receiptService.CreateReceipt(receipt)
	if err != nil {
		if receiptErr, ok := err.(*receipts.ReceiptExistsError); ok {
			c.JSON(http.StatusConflict, gin.H{"error": receiptErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": createdReceipt.ID})
}

func GetReceiptPoints(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(body) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	id := c.Param("id")
	receiptUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	points, err := receiptService.GetPointsByUUID(receiptUUID.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

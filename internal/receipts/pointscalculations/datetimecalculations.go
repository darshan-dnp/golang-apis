package pointscalculations

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type DateTimeCalculations struct{}

func (i *DateTimeCalculations) PointsCalculation(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error) {
	count := 0

	// Count Date Points
	if receipt.PurchaseDate != "" {
		parts := strings.Split(receipt.PurchaseDate, "-")
		if len(parts) != 3 {
			return 0, fmt.Errorf("invalid date format: date %s", receipt.PurchaseDate)
		}

		purchaseDay, err := strconv.Atoi(parts[2])
		if err != nil {
			return 0, fmt.Errorf("invalid date format: date %s", receipt.PurchaseDate)
		}

		if purchaseDay%2 != 0 {
			count += pointsVariables.PurchaseDatePoint
		}
	}

	// Count Time Points
	if receipt.PurchaseTime != "" {
		layout := "15:04"
		parsedTime, err := time.Parse(layout, receipt.PurchaseTime)
		if err != nil {
			return 0, fmt.Errorf("invalid time format: time %s", receipt.PurchaseTime)
		}

		validStartTime, _ := time.Parse(layout, pointsVariables.PurchaseTimeAfter)
		validEndTime, _ := time.Parse(layout, pointsVariables.PurchaseTimeBefore)

		if parsedTime.After(validStartTime) && parsedTime.Before(validEndTime) {
			count += pointsVariables.PurchaseTimePoints
		}
	}

	log.Println("Date Time Points" + strconv.Itoa(count))

	return count, nil
}

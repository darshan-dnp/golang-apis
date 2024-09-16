package pointscalculations

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"log"
	"strconv"
	"unicode"
)

type RetailerNameCalculations struct{}

func (i *RetailerNameCalculations) PointsCalculation(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error) {
	count := 0
	for _, retailerChar := range receipt.Retailer {
		if unicode.IsLetter(retailerChar) || unicode.IsNumber(retailerChar) {
			count++
		}
	}

	earnedPoints := count * pointsVariables.RetailerNamePoints

	log.Println("Retailer Name Points" + strconv.Itoa(earnedPoints))
	return earnedPoints, nil
}

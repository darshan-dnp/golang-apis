package pointscalculations

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"fmt"
	"log"
	"math"
	"strconv"
)

type TotalPriceCalculations struct{}

func (i *TotalPriceCalculations) PointsCalculation(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error) {
	count := 0

	totalPrice, err := strconv.ParseFloat(receipt.Total, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid total price: price %s", receipt.Total)
	}

	if math.Mod(totalPrice, 1) == 0 {
		count += pointsVariables.RoundAmountPoints
	}

	if math.Mod(totalPrice, 0.25) == 0 {
		count += pointsVariables.QuarterMultiplePoints
	}

	log.Println("Total Price Points" + strconv.Itoa(count))
	return count, nil
}

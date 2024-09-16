package pointscalculations

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type ItemsCalculations struct{}

func (i *ItemsCalculations) PointsCalculation(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error) {

	count := 0

	//Count number of pairs of Items
	pairs := len(receipt.Items) / 2
	count += pairs * pointsVariables.ItemPairPoints

	for _, item := range receipt.Items {
		shortDescriptionLen := len(strings.TrimSpace(item.ShortDescription))
		if shortDescriptionLen%pointsVariables.ItemDescLenDivisor == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				return 0, fmt.Errorf("invalid item price: description: %s, price %s", item.ShortDescription, item.Price)
			}
			count += int(math.Ceil(itemPrice * pointsVariables.ItemDescPriceMultiplier))
		}
	}

	log.Println("Items Points" + strconv.Itoa(count))
	return count, nil
}

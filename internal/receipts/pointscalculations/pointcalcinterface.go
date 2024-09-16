package pointscalculations

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
)

type PointsCalculation interface {
	PointsCalculation(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error)
}

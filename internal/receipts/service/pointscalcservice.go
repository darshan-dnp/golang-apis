package service

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"fetch-golang-api/internal/receipts/pointscalculations"
)

type PointsCalculationService struct {
	calculations []pointscalculations.PointsCalculation
}

func NewPointCalculationService() *PointsCalculationService {
	return &PointsCalculationService{
		calculations: []pointscalculations.PointsCalculation{
			&pointscalculations.ItemsCalculations{},
			&pointscalculations.DateTimeCalculations{},
			&pointscalculations.RetailerNameCalculations{},
			&pointscalculations.TotalPriceCalculations{},
		},
	}
}

func (s *PointsCalculationService) ApplyPointsCalculations(receipt *model.Receipt, pointsVariables cache.PointsVariables) (int, error) {
	totalPoints := 0
	for _, calc := range s.calculations {
		points, err := calc.PointsCalculation(receipt, pointsVariables)
		if err != nil {
			return 0, err
		}
		totalPoints += points
	}
	return totalPoints, nil
}

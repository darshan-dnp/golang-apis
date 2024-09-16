package service

import (
	"fetch-golang-api/internal/cache"
	"fetch-golang-api/internal/model"
	"fetch-golang-api/internal/receipts"
	"fetch-golang-api/pkg/utils"

	"gorm.io/gorm"
)

type Service struct {
	DB                *gorm.DB
	Cache             *cache.Cache
	PointsCalculation *PointsCalculationService
}

func NewService(db *gorm.DB, cache *cache.Cache) *Service {
	return &Service{
		DB:                db,
		Cache:             cache,
		PointsCalculation: NewPointCalculationService()}
}

func (s *Service) CreateReceipt(receipt model.Receipt) (*model.Receipt, error) {
	// Check if the receipt has already been uploaded
	// A User should not be able to resubmit the receipt. As there is no mapping for a User yet,
	// this query needs to be updated to work with User

	var count int64
	dataHash, err := utils.GenerateHashFromJSON(receipt)
	if err != nil {
		return nil, err
	}
	receipt.DataHash = dataHash

	if err := s.DB.Model(&model.Receipt{}).Where("data_hash = ?", receipt.DataHash).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, receipts.NewReceiptExistsError("receipt has already been submitted")
	}

	// Get Points Variables from cache
	pointsVariables, err := s.Cache.GetPointsVariables()
	if err != nil {
		return nil, err
	}

	// Calculte all points earned by the receipt
	if err := s.CalculateReceiptPoints(&receipt, pointsVariables); err != nil {
		return nil, err
	}

	// Store Receipt in DB
	if err := s.DB.Create(&receipt).Error; err != nil {
		return nil, err
	}
	return &receipt, nil

}

func (s *Service) CalculateReceiptPoints(receipt *model.Receipt, pointsVariables cache.PointsVariables) error {

	earnedPoints, err := s.PointsCalculation.ApplyPointsCalculations(receipt, pointsVariables)
	if err != nil {
		return err
	}

	// Assign points to receipt object and save in DB
	receipt.Points = earnedPoints
	return nil
}

func (s *Service) GetPointsByUUID(id string) (int, error) {

	var receiptPoints int
	if err := s.DB.Model(&model.Receipt{}).
		Select("points").
		Where("id = ?", id).
		Scan(&receiptPoints).Error; err != nil {
		return 0, err
	}

	return receiptPoints, nil
}

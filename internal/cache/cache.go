package cache

import (
	"time"

	"gorm.io/gorm"
)

type PointsVariables struct {
	RoundAmountPoints       int
	QuarterMultiplePoints   int
	RetailerNamePoints      int
	ItemPairPoints          int
	ItemDescLenDivisor      int
	ItemDescPriceMultiplier float64
	PurchaseDatePoint       int
	PurchaseTimeAfter       string
	PurchaseTimeBefore      string
	PurchaseTimePoints      int
	Valid_till              time.Time
}

type Cache struct {
	DB    *gorm.DB
	Cache PointsVariables
}

func NewCache(db *gorm.DB) *Cache {
	return &Cache{DB: db}
}

func (c *Cache) GetPointsVariables() (PointsVariables, error) {
	if c.Cache.Valid_till.After(time.Now()) {
		return c.Cache, nil
	}
	return c.loadPointsVariablesFromDB()
}

func (c *Cache) loadPointsVariablesFromDB() (PointsVariables, error) {
	var pointsVariables PointsVariables

	err := c.DB.Raw(`
        SELECT *
        FROM points_validity_parameters
		WHERE valid_till > DATE('now')
		ORDER BY valid_till ASC
		LIMIT 1
    `).Scan(&pointsVariables).Error

	if err != nil {
		return PointsVariables{}, err
	}

	c.Cache = pointsVariables
	return pointsVariables, nil
}

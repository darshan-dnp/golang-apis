package db

import (
	"fetch-golang-api/internal/model"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&model.Receipt{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&model.Item{})
	if err != nil {
		return nil, err
	}

	query := `
    CREATE TABLE IF NOT EXISTS points_validity_parameters (
        round_amount_points INTEGER,
        quarter_multiple_points INTEGER,
        retailer_name_points INTEGER,
        item_pair_points INTEGER,
        item_desc_len_divisor INTEGER,
		item_desc_price_multiplier REAL,
		purchase_date_point INTEGER,
		purchase_time_after TEXT,
		purchase_time_before TEXT,
		purchase_time_points INTEGER,
        Valid_till DATETIME
    );`

	if err := DB.Exec(query).Error; err != nil {
		log.Fatalf("failed to create points_validity_parameters table: %v", err)
	}

	insertQuery := `
    INSERT INTO points_validity_parameters
    VALUES (50, 25, 1, 5, 3, 0.2, 6, '14:00', '16:00', 10, ?);`

	if err := DB.Exec(insertQuery, time.Now().Add(24*time.Hour)).Error; err != nil {
		log.Fatalf("failed to insert points validity parameters: %v", err)
	}

	return DB, nil
}

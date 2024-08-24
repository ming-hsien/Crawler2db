package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (db *gorm.DB, err error) {
	dsn := "Your postgres"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

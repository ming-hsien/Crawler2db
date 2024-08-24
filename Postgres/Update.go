package postgres

import (
	"Crawler2db/crawler"
	"log"
)

type dbInfo struct {
	stationName string
	
}

func Update(MetroInfo map[string]crawler.MetroInfo) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&crawler.MetroInfo{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	for _, v := range(MetroInfo) {
		result := db.Create(v)
		if result.Error != nil {
			log.Fatal("failed to insert data", result.Error)
		}
	}
	
	
	return nil
}

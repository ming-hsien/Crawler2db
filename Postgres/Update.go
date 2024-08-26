package postgres

import (
	"Crawler2db/crawler"
	"log"
)

func Update(MetroInfo map[string]crawler.MetroInfo) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DBInfo{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	for k, v := range(MetroInfo) {
		for _, station := range v.StationName {
			newdb := &DBInfo{
				Line: k,
				StationName: station,
				ArrivalTimes: v.StationInfo[station].ArrivalTimeList,
			}
			result := db.Create(newdb)
			if result.Error != nil {
				log.Fatal("failed to insert data", result.Error)
				break
			}
		}
	}

	return nil
}

package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type stringList []string

func (s stringList) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *stringList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, s)
}

type DBInfo struct {
	Line         string     `gorm:"type:varchar(30)"`
	StationName  string     `gorm:"type:varchar(30)"`
	ArrivalTimes stringList `gorm:"type:jsonb"`
}

func (DBInfo) TableName() string {
	return "metroinfo"
}

func connectDB() (db *gorm.DB, err error) {
	dsn := "Your postgres"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

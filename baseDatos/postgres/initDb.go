package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func NewPostgres() (*gorm.DB, error) {
	connectionString := "postgres://postgres:F4mC4sh3r:2016@localhost:5432/postgres?sslmode=disable"
	DB, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetConnMaxLifetime(time.Second * 60)
	DB.LogMode(true)

	return DB, nil
}

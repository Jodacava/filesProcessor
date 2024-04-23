package baseDatos

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func NewPostgres() (*gorm.DB, error) {
	connectionString := "postgresql://postgres:gqMgnayv9nR7u7i@tizon.cjxcsq7qqmym.us-east-1.rds.amazonaws.com:5432/stori"
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

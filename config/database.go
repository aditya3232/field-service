package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"time"
)

func InitDatabase() (*gorm.DB, error) {
	config := Config
	encodedPassword := url.QueryEscape(config.Database.Password)
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		encodedPassword,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxLifeTimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime) * time.Second)
	return db, nil
}

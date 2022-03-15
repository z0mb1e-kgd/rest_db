package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < dbConnectAttempts; i++ {
			time.Sleep(dbRetryPeriod)
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return db
			}
		}
		panic(err)
	}
	return db
}

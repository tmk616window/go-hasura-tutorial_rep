package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	dsn := "host=postgresql user=user password=password dbname=db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

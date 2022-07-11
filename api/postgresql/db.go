package postgresql

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	dsn := fmt.Sprintf("host=10.125.208.7 user=user password=password port=5432 database=db")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return db
}

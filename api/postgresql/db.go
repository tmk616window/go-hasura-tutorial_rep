package postgresql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s  port=%s database=%s", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("PORT"), os.Getenv("DBNAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return db
}

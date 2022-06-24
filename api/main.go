package main

import (
	"api_test/models"
	"fmt"
)


func main() {
	db := models.DbConnect()
	fmt.Println(db)
	// defer db.Close()
	// db.LogMode(true)
}

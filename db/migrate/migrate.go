package main

import (
	"fmt"

	"github.com/lawnmower-74/storage_app_server/db"
	"github.com/lawnmower-74/storage_app_server/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
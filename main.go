package main

import "gorm.io/gorm"

var db *gorm.DB

func main() {
	router := SetRouter()
	db = ConnectDB()
	router.Run(serverPort)
}

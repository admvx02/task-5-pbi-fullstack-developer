package main

import (
	"app/config"
	"app/models"
)

func main() {
	config.SetupDatabase()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Photo{})
}

package db

import (
	"gocrawler/app/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.PageData{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migration completed!")
}

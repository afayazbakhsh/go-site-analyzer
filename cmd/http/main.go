package main

import (
	"gocrawler/app/db"
	"gocrawler/app/httpserver"
	"log"
)

func main() {
	log.Println("Starting Go app...")

	// 1️⃣ Initialize database
	if err := db.Init(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.Migrate()

	log.Println("DB connected successfully")

	// 2️⃣ Start HTTP server
	if err := httpserver.Run(":8282"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

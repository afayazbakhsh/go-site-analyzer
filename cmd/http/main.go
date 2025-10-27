package main

import (
	"gocrawler/app/httpserver"
	"log"
)

func main() {
	if err := httpserver.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

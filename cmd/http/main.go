package main

import (
	"gocrawler/app/commands"
	"gocrawler/app/db"
	"gocrawler/app/httpserver"
	"gocrawler/app/httpserver/requests"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	log.Println("Starting Go app...")

	//********* Initialize database *********** //

	if err := db.Init(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.Migrate()
	log.Println("DB connected successfully")

	//********* Initialize Custom Validators *********** //

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pwd", requests.PasswordStrength)
	}

	//********* Initialize CLI *********** //

	commands.Execute()

	// 2️⃣ Start HTTP server
	if err := httpserver.Run(":8282"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

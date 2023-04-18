package main

import (
	"log"
	"os"

	"gym/server"
	"gym/server/db"
	"gym/server/model"
	"gym/server/services/authentication"
	"gym/server/services/slots"

	"github.com/joho/godotenv"
)

//	@title			Gym Managment App
//	@version		1.0.0
//	@description	This is a API documnetation for gym management
//	@host			localhost:3000
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	authentication.TwilioInit(os.Getenv("TWILIO_AUTH_TOKEN"))
	connection := db.InitDB()
	db.Transfer(connection)

	if connection.Migrator().HasTable(&model.Slot{}) {

		var slot []model.Slot
		query := "SELECT * FROM slots ORDER BY slot_id ASC;"
		db.QueryExecutor(query, &slot)
		if slot == nil {
			slots.SlotDistribution()
		}

	}
	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}

}

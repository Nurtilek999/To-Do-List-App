package main

import (
	"Pet_1/api"
	"Pet_1/pkg/config"
	"Pet_1/pkg/database"
	"log"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	port := ":8080"
	app := api.SetupRouter(db)
	app.Run(port)
}

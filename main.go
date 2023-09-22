package main

import (
	"api/database"
	"api/routes"
)

func main() {
	db, err := database.InitDB()

	if err != nil {
		panic("Failed to initialize the database: " + err.Error())
	}

	r := routes.SetupRouter(db)
	r.Run(":8000")
}

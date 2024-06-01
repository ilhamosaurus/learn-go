package main

import (
	"belajar/database"
	"belajar/router"
)

func main() {
	PORT := ":5003"

	database.StartDB()
	router.StartServer().Run(PORT)
}

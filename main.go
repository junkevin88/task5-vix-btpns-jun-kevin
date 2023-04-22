package main

import (
	"btpn-backend-go/database"
	"btpn-backend-go/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}

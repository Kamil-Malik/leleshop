package main

import (
	"leleshop/db"
	"leleshop/router"
)

func main() {
	db.StartDB()
	router.StartServer()
}

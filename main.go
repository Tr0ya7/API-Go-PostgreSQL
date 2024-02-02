package main

import (
	"api/database"
	"api/router"
)

func main() {
	database.DbConect()
	router.HandleRequest()
}

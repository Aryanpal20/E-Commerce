package main

import (
	"e-Commerce/database"
	routing "e-Commerce/routings"
)

func main() {
	database.DataMigration()
	routing.HandlerRouting()
}

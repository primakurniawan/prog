package main

import (
	"prog/db"
	"prog/migrate"
	"prog/routes"
)

func main() {
	db.InitDB()
	migrate.AutoMigrate()
	e := routes.New()
	e.Start(":8000")
}

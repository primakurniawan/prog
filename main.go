package main

import (
	"prog/db"
	"prog/routes"
)

func main() {
	db.InitDB()
	e := routes.New()
	e.Start(":8000")
}

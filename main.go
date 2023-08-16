package main

import (
	"1program.com/tournament_app/conn"
	"1program.com/tournament_app/db"
	"1program.com/tournament_app/routes"
)

func main() {
	routes.AttachRoutes()
	db.Connect()
	defer db.Close()
	conn.Start()
}

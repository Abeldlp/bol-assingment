package main

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/route"
)

func main() {

	// Initial setup
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()

	// Migrations
	migration.PlayerMigration()
	migration.GameMigration()

	// Initialize server
	r := config.InitializeServer()

	//Routes
	route.AppendMancalaGameRoute(r)

	// Run server
	r.Run(":8000")
}

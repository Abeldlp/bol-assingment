package main

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/route"
)

func main() {

	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()

	migration.PlayerMigration()
	migration.GameMigration()

	r := config.InitializeServer()

	route.AppendMancalaGameRoute(r)

	r.Run(":8000")
}

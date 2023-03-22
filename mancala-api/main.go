package main

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/route"
)

//	@title			Mancala API
//	@version		1.0
//	@description	Mancala api for bol assignment

//	@contact.email	abel45991690@gmail.com

//	@host		localhost:5000
//	@BasePath	/v1/api

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()

	migration.PlayerMigration()
	migration.GameMigration()

	r := config.InitializeServer()

	route.AppendMancalaGameRoute(r)
	route.AppendDocumentationRoute(r)

	r.Run(":8000")
}

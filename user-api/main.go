package main

import (
	"github.com/Abeldlp/bol-assignment/user-api/config"
	"github.com/Abeldlp/bol-assignment/user-api/route"
)

func main() {
	//Initial setup
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()
	r := config.InitializeServer()

	//Routes
	route.AppendUserRoute(r)

	r.Run(":8000")
}

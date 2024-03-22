package main

import (
	repo "MyGramAtta/repo"
	routers "MyGramAtta/routers"
)

// @title My Gram API
// @version 1.0
// @description This is a final project Hacktiv8
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header -> Bearer
// @name Authorization
func main() {
    repo.CreateDatabase()
    repo.StartDB()
    defer repo.CloseDB()
    repo.EnsureTablesCreated()
	routers.StartServer().Run(":8080")
}

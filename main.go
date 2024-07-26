package main

import (
	"be-empower-hr/app/config"
	"be-empower-hr/app/databases"
	"be-empower-hr/app/migrations"
	"be-empower-hr/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbPostgresql := databases.InitDBpostgre(cfg)
	migrations.RunMigrations(dbPostgresql)

	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// CORS configuration
	e.Use(middleware.CORS())

	// Routes
	routes.InitRouter(e, dbPostgresql)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

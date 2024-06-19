package main

import (
	"github.com/Maniexie/crud-fiber-postgres/database"
	"github.com/Maniexie/crud-fiber-postgres/database/migration"
	"github.com/Maniexie/crud-fiber-postgres/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()
	migration.RunMigrate()
	app := fiber.New()

	routers.RouterApp(app)
	app.Listen(":3000")
}

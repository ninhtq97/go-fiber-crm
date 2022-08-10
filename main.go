package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ninhtq97/go-fiber-crm/database"
	"github.com/ninhtq97/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/lead", lead.GetLeads)
	api.Get("/lead/:id", lead.GetLead)
	api.Post("/lead", lead.NewLead)
	api.Delete("/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	database.Connect()
	fmt.Println("Connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":8000")
}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ninhtq97/go-fiber-crm/database"
	"github.com/ninhtq97/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Group("/api/v1")

	app.Get("/lead", lead.GetLeads)
	app.Get("/lead/:id", lead.GetLead)
	app.Post("/lead", lead.NewLead)
	app.Delete("/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	db := database.Connect()

	fmt.Println("Connection opened to database")
	db.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":8000")
}

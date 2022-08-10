package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ninhtq97/go-fiber-crm/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBconn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).SendString("No lead found with Id")
		return nil
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}

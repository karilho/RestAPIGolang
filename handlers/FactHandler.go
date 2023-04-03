package handlers

import (
	"RestAPIGolang/database"
	"RestAPIGolang/models"
	"github.com/jeffotoni/quick"
)

func Home(c *quick.Ctx) error {
	return c.SendString("Bem vindo ao APP")
}

func ListAll(c *quick.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	c.Set("Content-Type", "application/json")
	return c.Status(200).JSON(facts)
}

func CreateFact(c *quick.Ctx) error {
	fact := new(models.Fact)
	err := c.BodyParser(&fact)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

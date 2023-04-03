package routes

import (
	"RestAPIGolang/handlers"
	"github.com/jeffotoni/quick"
)

func SetupRoutes(c *quick.Quick) {
	c.Get("/", handlers.Home)

	c.Get("/facts", handlers.ListAll)

	c.Post("/new", handlers.CreateFact)
}

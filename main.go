package main

import (
	"RestAPIGolang/database"
	"RestAPIGolang/routes"
	"fmt"
	"github.com/jeffotoni/quick"
)

func main() {
	database.ConnectDb()
	app := quick.New()
	fmt.Println("Aplicação iniciada")
	routes.SetupRoutes(app)
	app.Listen(":3000")
}

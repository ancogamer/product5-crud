package main

import (
	"fmt"
	"log"

	"github.com/ancogamer/product5/api.crud/config"
	route "github.com/ancogamer/product5/api.crud/controller/handler"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	route.DbStruct{PostgreSQLString: db.Connect()}
	route.AllRoutes(app)

	log.Fatal(app.Listen(config.HTTPort))

	fmt.Println("Server Started!")
}

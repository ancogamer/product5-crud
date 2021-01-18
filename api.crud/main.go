package main

import (
	"fmt"
	"log"

	"github.com/ancogamer/product5/api.crud/config"
	route "github.com/ancogamer/product5/api.crud/controller/handler"
	db "github.com/ancogamer/product5/api.crud/pkg/psql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	route.AllRoutes(route.DbStruct{PostgreSQLString: db.Connect()}, app)

	log.Fatal(app.Listen(config.HTTPort))

	fmt.Println("Server Started!")
}

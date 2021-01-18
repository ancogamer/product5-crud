package handler

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"

	mw "github.com/gofiber/fiber/middleware"
)

// AllRoutes , aqui se encontra todas as rotas
func AllRoutes(app *fiber.App) {
	////
	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	///

	// application/json
	// application/x-www-form-urlencoded
	// multipart/form-data

}

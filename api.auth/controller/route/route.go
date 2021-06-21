package route

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"
)

//
func AllRoutes(db DbStruct, app *fiber.App) {
	////
	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	// application/json
	// application/x-www-form-urlencoded
	// multipart/form-data

}

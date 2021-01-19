package handler

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber/v2"

	mw "github.com/gofiber/fiber/middleware"
)

// AllRoutes , aqui se encontra todas as rotas
func AllRoutes(s DbStruct, app *fiber.App) {
	////
	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	table := app.Group("api1/table/")
	///

	// application/json
	// application/x-www-form-urlencoded
	// multipart/form-data
	table.Post("", s.TablePost)
	table.Get(":offset/:limit", s.TableGetAll)
	table.Get(":uuid", s.TableGetOne)
	table.Put("", s.TableUpdate)
	table.Delete(":uuid", s.TableDelete)

}

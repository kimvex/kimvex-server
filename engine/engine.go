package engine

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"

	"../db"
	"../routes"
)

//ServerExecute function of exution api
func ServerExecute() {
	app := fiber.New()
	app.Use(logger.New())

	database := db.MySQLConnect()
	redisC := db.RedisConnect()
	getUser := db.GetUserID("ytr")

	routes.API(app, database, redisC, getUser)
	app.Listen(3000)
}

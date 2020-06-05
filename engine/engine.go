package engine

import (
	"fmt"

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
	getUser := db.GetUserID
	fmt.Println(getUser, "cas")

	routes.API(app, database, redisC, getUser)
	app.Listen(3000)
}

package engine

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"

	"../db"
	"../routes"
)

//ServerExecute function of exution api
func ServerExecute() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
		},
	))

	database := db.MySQLConnect()
	db.RedisConnect()
	mongodb := db.MonoDBConnect()
	getUser := db.GetUserID
	setUser := db.SetUserID
	delUser := db.DeleteUserID

	routes.API(app, database, getUser, setUser, delUser, mongodb)
	app.Listen(3003)
}

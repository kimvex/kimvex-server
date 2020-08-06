package routes

import (
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	apiRoute fiber.Router
	database *sql.DB
	userIDF  func(string) string
	setID    func(string, string)
	delID    func(string)
	mongodb  *mongo.Database
)

//ValidateRoute endpoint for validate users
var ValidateRoute = func(c *fiber.Ctx) {
	if c.Get("token") != "" {
		token, err := jwt.Parse(c.Get("token"), func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token.Valid {
			fmt.Println(c.Get("token"))
			validateRedis := userIDF(c.Get("token"))
			if len(validateRedis) > 0 {
				c.Next()
			} else {
				c.JSON(ErrorResponse{MESSAGE: "Token expired"})
				c.SendStatus(401)
			}
			return
		}

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(ErrorResponse{MESSAGE: "Token structure not is valid"})
				c.Status(401)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.JSON(ErrorResponse{MESSAGE: "Token is expired"})
				c.Status(401)
			} else {
				c.JSON(ErrorResponse{MESSAGE: "Invalid token"})
				c.Status(401)
			}
		}
		return
	}

	c.JSON(ErrorResponse{MESSAGE: "Without token"})
	c.Status(401)
	return
}

//API function pricipal for backboune
func API(app *fiber.App, Database *sql.DB, UserIDC func(string) string, SetIDC func(string, string), DelIDC func(string), Mongodb *mongo.Database) {

	apiRoute = app.Group("/api")
	database = Database
	userIDF = UserIDC
	setID = SetIDC
	delID = DelIDC
	mongodb = Mongodb

	Users()
	Shops()

	apiRoute.Get("/", func(c *fiber.Ctx) {
		userID := userIDF(c.Get("token"))

		var response SuccessResponse
		fmt.Println(userID, "11")
		response.MESSAGE = "Raiz del proyecto"
		c.JSON(response)
	})
}

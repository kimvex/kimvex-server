package routes

import (
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/gomodule/redigo/redis"
)

var (
	apiRoute *fiber.Group
	database *sql.DB
	redisC   redis.Conn
	userIDF  func(string) string
)

//ValidateRoute endpoint for validate users
var ValidateRoute = func(c *fiber.Ctx) {
	if c.Get("token") != "" {
		token, err := jwt.Parse(c.Get("token"), func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token.Valid {
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
func API(app *fiber.App, Database *sql.DB, RedisCl redis.Conn, UserIDC func(string) string) {

	apiRoute = app.Group("/api")
	database = Database
	redisC = RedisCl
	userIDF = UserIDC

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

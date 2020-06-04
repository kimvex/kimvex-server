package routes

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber"
)

//ErrorResponse Structure for response of type error api
type ErrorResponse struct {
	MESSAGE string `json:"message"`
}

//SuccessResponse structure for response of type success api
type SuccessResponse struct {
	MESSAGE string `json:"message"`
}

var (
	app      *fiber.App
	database *sql.DB
	redisC   *redis.Client
	userID   string
)

//API function pricipal for backboune
func API(App *fiber.App, Database *sql.DB, RedisCl *redis.Client, UserIDC string) {
	app = App
	database = Database
	redisC = RedisCl
	userID = UserIDC

	app.Get("/", func(c *fiber.Ctx) {
		var response SuccessResponse

		response.MESSAGE = "Raiz del proyecto"
		c.JSON(response)
	})
}

//ValidateRoute endpoint for validate users
func ValidateRoute(c *fiber.Ctx) {
	if c.Get("token") != "" {
		token, err := jwt.Parse(c.Get("token"), func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token.Valid {
			c.Next()
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

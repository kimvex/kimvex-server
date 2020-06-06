package routes

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

//ReponseProfile struct
type ReponseProfile struct {
	UserID   *string `json:"user_id"`
	Fullname *string `json:"fullname"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Age      *string `json:"age"`
	Gender   *string `json:"gender"`
	Image    *string `json:"image"`
	CreateAt *string `json:"create_at"`
	Code     *string `json:"code"`
}

//Users Namespace for endpoint of users
func Users() {
	apiRoute.Post("/login", Login)
	apiRoute.Get("/profile", Profile)
}

//Login Handler for endpoint
func Login(c *fiber.Ctx) {
	singSecret := []byte("secret")

	var userLogin User
	if err := c.BodyParser(&userLogin); err != nil {
		fmt.Println(err, "Error parsing login")
	}

	var selectedUser User

	errorGetUser := sq.Select("email", "password", "user_id").
		From("usersk").
		Where(sq.Eq{"email": userLogin.Email}).
		RunWith(database).
		QueryRow().
		Scan(&selectedUser.Email, &selectedUser.Password, &selectedUser.UserID)

	if errorGetUser != nil {
		fmt.Println(errorGetUser, "Error get user login")
		error := ErrorResponse{MESSAGE: "Problem with get user"}
		c.JSON(error)
		c.Status(401)
		return
	}

	compare := bcrypt.CompareHashAndPassword([]byte(string(selectedUser.Password)), []byte(userLogin.Password))

	if compare == nil {
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "kimvex",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenResolve, setErr := token.SignedString(singSecret)
		if setErr != nil {
			fmt.Println("Error with generate token", setErr)
			error := ErrorResponse{MESSAGE: "Problem with generating token"}
			c.JSON(error)
			c.Status(500)
			return
		}

		redisC.Do("SET", tokenResolve, selectedUser.UserID)
		c.JSON(TokenResponse{Token: tokenResolve})
		return
	}

	error := ErrorResponse{MESSAGE: "Usuario ó contraseña incorrectos"}
	c.JSON(error)
	c.Status(401)
}

//Profile Handler for endpoint
func Profile(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var profileSQL ProfileStruct
	var response ReponseProfile

	userProfile, err := sq.Select(
		"usersk.user_id",
		"fullname",
		"email",
		"phone",
		"age",
		"gender",
		"image",
		"create_at",
		"code",
	).
		From("usersk").
		LeftJoin("code_reference on code_reference.user_id = usersk.user_id").
		Where(sq.Eq{"usersk.user_id": userID}).
		RunWith(database).
		Query()

	if err != nil {
		fmt.Println(err, "Error to get profile")
		ErrorProfile := ErrorResponse{MESSAGE: "Error to get profile"}
		c.JSON(ErrorProfile)
		c.Status(400)
		return
	}

	for userProfile.Next() {
		_ = userProfile.Scan(
			&profileSQL.UserID,
			&profileSQL.Fullname,
			&profileSQL.Email,
			&profileSQL.Phone,
			&profileSQL.Age,
			&profileSQL.Gender,
			&profileSQL.Image,
			&profileSQL.CreateAt,
			&profileSQL.Code,
		)

		response.UserID = &profileSQL.UserID.String
		response.Fullname = &profileSQL.Fullname.String
		response.Email = &profileSQL.Email.String
		response.Phone = &profileSQL.Phone.String
		response.Age = &profileSQL.Age.String
		response.Gender = &profileSQL.Gender.String
		response.Image = &profileSQL.Image.String
		response.CreateAt = &profileSQL.CreateAt.String
		response.Code = &profileSQL.Code.String
	}

	c.JSON(response)
}

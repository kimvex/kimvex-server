package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../helper"
	sq "github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

//Users Namespace for endpoint of users
func Users() {
	apiRouteUser := apiRoute.Group("/users")
	apiRouteUser.Post("/login", Login)
	apiRouteUser.Get("/profile", Profile)
	apiRouteUser.Post("/register", Register)
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

//Register Handler for endpoint
func Register(c *fiber.Ctx) {
	var UserData RegisterData
	var existUser ValidateExistUser

	rts, te := http.Get("https://donfreddy.kimvex.com")
	fmt.Println(rts, "que paso2", te)

	if errorParse := c.BodyParser(&UserData); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	fmt.Println(UserData.Email, UserData.Password, UserData.Fullname, UserData.Age, UserData.Phone, UserData.Gender)
	if len(UserData.Email) == 0 || len(UserData.Password) == 0 || len(UserData.Fullname) == 0 || len(UserData.Age) == 0 || UserData.Phone == 0 || len(UserData.Gender) == 0 {
		c.JSON(ErrorResponse{MESSAGE: "Incomplete data"})
		c.Status(400)
		return
	}

	sq.Select("email").
		From("usersk").
		Where(sq.Eq{"email": UserData.Email}).
		RunWith(database).
		QueryRow().
		Scan(&existUser.Email)

	if len(existUser.Email) > 0 {
		fmt.Println("El usuario ya esta registrado", existUser.Email)
		error := ErrorResponse{MESSAGE: "Users exist"}
		c.JSON(error)
		c.Status(400)
		return
	}

	// if ErrorGetUser != nil {
	// 	fmt.Println(ErrorGetUser)
	// 	error := ErrorResponse{MESSAGE: "Problem with get user information"}
	// 	c.JSON(error)
	// 	c.Status(400)
	// 	return
	// }

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(UserData.Password), 14)

	id, errorInsert := sq.Insert("usersk").
		Columns(
			"email",
			"password",
			"fullname",
			"age",
			"phone",
			"gender",
			"status",
		).
		Values(
			UserData.Email,
			string(passwordHash),
			UserData.Fullname,
			UserData.Age,
			UserData.Phone,
			UserData.Gender,
			0,
		).
		RunWith(database).
		Exec()

	fmt.Println(errorInsert, "si")
	if errorInsert != nil {
		fmt.Println(errorInsert)
		ErrorI := ErrorResponse{MESSAGE: "No se pudo registrar al usuario"}
		c.JSON(ErrorI)
		c.SendStatus(400)
		return
	}

	r, _ := id.LastInsertId()

	if r > 0 {
		firstHS, _ := helper.RandomCode(2)
		secondHS, _ := helper.RandomCode(2)
		thirdHS, _ := helper.RandomCode(2)
		fourHS, _ := helper.RandomCode(2)
		codeRef := fmt.Sprintf("%v-%v-%v-%v", firstHS, secondHS, thirdHS, fourHS)

		_, errorInsertCode := sq.Insert("code_reference").
			Columns("code", "user_id").
			Values(codeRef, r).
			RunWith(database).
			Exec()

		if errorInsertCode != nil {
			fmt.Println(errorInsertCode, "problem with generate code")
			c.JSON(ErrorResponse{MESSAGE: "Problem with generate code"})
			return
		}

		requestBody, _ := json.Marshal(map[string]string{
			"send_to": UserData.Email,
		})

		rps, e := http.Post("https://process.kimvex.com/api/code_send_mail", "application/json", bytes.NewBuffer(requestBody))

		fmt.Println(rps, "que paso", e, codeRef)

		if e != nil {
			fmt.Println(e, "to post code")
		}

		success := SuccessResponse{MESSAGE: "El usuario se registro con exito"}
		c.JSON(success)
	}
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

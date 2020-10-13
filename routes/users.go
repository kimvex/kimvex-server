package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"../helper"
	sq "github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

//Users Namespace for endpoint of users
func Users() {
	apiRouteUser := apiRoute.Group("/user")
	apiRouteCode := apiRoute.Group("/code")

	//Validations token
	apiRouteUser.Use("/profile", ValidateRoute)
	apiRouteUser.Use("/update/profile", ValidateRoute)
	apiRouteUser.Use("/restore_password", ValidateRoute)
	apiRouteCode.Use("/auth", ValidateRoute)
	// apiRouteUser.Use("/code/auth", ValidateRoute)
	apiRouteUser.Use("/referrals", ValidateRoute)
	apiRouteUser.Use("/my_code", ValidateRoute)
	apiRouteUser.Use("/refferals_fail", ValidateRoute)
	apiRouteUser.Use("/earned_referrals", ValidateRoute)
	apiRouteUser.Use("/earned_referrals_month", ValidateRoute)

	apiRouteUser.Post("/login", Login)
	apiRouteUser.Get("/profile", Profile)
	apiRouteUser.Post("/register", Register)
	apiRouteUser.Put("/update/profile", UpdateProfileEnd)
	apiRouteUser.Post("/logout", Logout)
	apiRouteUser.Post("/restore_password", RestorePassword)
	apiRouteCode.Get("/auth", CodeAuth)
	apiRouteUser.Get("/referrals", Referrals)
	apiRouteUser.Get("/my_code", MyCodeHandler)
	apiRouteUser.Get("/refferals_fail", ReferralsFail)
	apiRouteUser.Get("/earned_referrals", EarnedReferrals)
	apiRouteUser.Get("/earned_referrals_month", EarnedReferralsMonth)
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
		Where(sq.Eq{"email": userLogin.Email, "status": 1}).
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

		fmt.Println(tokenResolve, selectedUser.UserID)
		setID(tokenResolve, selectedUser.UserID)
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
			Values(strings.ToUpper(codeRef), r).
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

//UpdateProfileEnd Handler for endpoint
func UpdateProfileEnd(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var profileSave UpdateProfile
	var baseUser BasicUser

	if errorParse := c.BodyParser(&profileSave); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	ErrorRequest := sq.Select(
		"user_id",
		"password",
	).
		From("usersk").
		Where(sq.Eq{"user_id": userID}).
		RunWith(database).
		QueryRow().
		Scan(&baseUser.UserID, &baseUser.Password)

	if ErrorRequest != nil {
		fmt.Println(ErrorRequest, "Error to get user update profile")
		c.JSON(ErrorResponse{MESSAGE: "Error to get user"})
		c.SendStatus(404)
		return
	}

	if len(profileSave.Password) > 0 && len(profileSave.NewPassword) > 0 {
		compare := bcrypt.CompareHashAndPassword([]byte(baseUser.Password), []byte(string(profileSave.Password)))

		if compare != nil {
			c.JSON(ErrorResponse{MESSAGE: "IncorrectPassword"})
			c.SendStatus(401)
			return
		}

		newPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(profileSave.NewPassword), 14)

		_, ErrorUpdatePassword := sq.Update("usersk").
			Set("password", newPasswordHash).
			Where(sq.Eq{"user_id": userID}).
			RunWith(database).
			Exec()

		if ErrorUpdatePassword != nil {
			fmt.Println(ErrorUpdatePassword, "Error to update new password")
			c.JSON(ErrorResponse{MESSAGE: "Problem to save new password"})
			c.SendStatus(500)
			return
		}

		c.JSON(SuccessResponse{MESSAGE: "Update password"})
		return
	}

	queryUpdateValue := sq.Update("usersk")

	if len(profileSave.Email) > 0 {
		queryUpdateValue = queryUpdateValue.Set("email", profileSave.Email)
	}

	if len(profileSave.Fullname) > 0 {
		queryUpdateValue = queryUpdateValue.Set("fullname", profileSave.Fullname)
	}

	if len(profileSave.Age) > 0 {
		queryUpdateValue = queryUpdateValue.Set("age", profileSave.Age)
	}

	if len(profileSave.Phone) > 0 {
		queryUpdateValue = queryUpdateValue.Set("phone", profileSave.Phone)
	}

	if len(profileSave.Gender) > 0 {
		queryUpdateValue = queryUpdateValue.Set("gender", profileSave.Gender)
	}

	if len(profileSave.Address) > 0 {
		queryUpdateValue = queryUpdateValue.Set("address", profileSave.Address)
	}

	if len(profileSave.URLImage) > 0 {
		queryUpdateValue = queryUpdateValue.Set("image", profileSave.URLImage)
	}

	_, ErrorUpdateProfile := queryUpdateValue.
		Where(sq.Eq{"user_id": userID}).
		RunWith(database).
		Exec()

	if ErrorUpdateProfile != nil {
		fmt.Println(ErrorUpdateProfile, "Problem with update profile")
		c.JSON(ErrorResponse{MESSAGE: "Problem with update profile"})
		c.SendStatus(500)
		return
	}

	c.JSON(SuccessResponse{MESSAGE: "Profile updated"})
}

//Logout Handler for endpoint
func Logout(c *fiber.Ctx) {
	Token := c.Get("token")

	delID(Token)

	c.JSON(SuccessResponse{MESSAGE: "Logout success"})
}

//RestorePassword Handler for endpoint
func RestorePassword(c *fiber.Ctx) {
	var data DataRestorePassword
	var user BasicUser

	if errorParse := c.BodyParser(&data); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	if len(data.NewPassword) > 0 || len(data.OldPassword) > 0 || len(data.Code) > 0 || len(data.Email) > 0 {
		fmt.Println("Parameters are missing")
		c.JSON(ErrorResponse{MESSAGE: "Parameters are missing"})
		c.SendStatus(400)
		return
	}

	ErrorGetCode := sq.Select("user_id", "password").
		From("code_restore").
		LeftJoin("usersk on code_restore.user_id=usersk.user_id").
		Where(sq.Eq{"code": data.Code, "email": data.Email, "active": 0}).
		RunWith(database).
		Scan(&user.UserID, &user.Password)

	if ErrorGetCode != nil {
		fmt.Println(ErrorGetCode, "Problem with get code")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get code"})
		c.SendStatus(500)
		return
	}

	compare := bcrypt.CompareHashAndPassword([]byte(string(data.OldPassword)), []byte(user.Password))

	if compare != nil {
		c.JSON(ErrorResponse{MESSAGE: "IncorrectPassword"})
		c.SendStatus(401)
		return
	}

	newPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 14)

	_, ErrorUpdatePassword := sq.Update("usersk").
		Set("password", newPasswordHash).
		Where(sq.Eq{"user_id": user.UserID}).
		RunWith(database).
		Exec()

	if ErrorUpdatePassword != nil {
		fmt.Println(ErrorUpdatePassword, "Error to update new password")
		c.JSON(ErrorResponse{MESSAGE: "Problem to save new password"})
		c.SendStatus(500)
		return
	}

}

//CodeAuth Handler for endpoint
func CodeAuth(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	code := c.Query("reffer_code")
	var structCode SQLCodeRef
	fmt.Println(code, "sio")
	ErrorGetCode := sq.Select("code", "code_reference_id").
		From("code_reference").
		Where("code = ? AND user_id <> ?", code, userID).
		RunWith(database).
		Scan(&structCode.Code, &structCode.CodeReferenceID)

	if ErrorGetCode != nil {
		fmt.Println(ErrorGetCode, "Problem with get code")
	}

	fmt.Println(strings.ToUpper("47a6-8a54-4310-463b"), "-.-")

	c.JSON(ResponseValidate{Validate: structCode})
}

//Referrals Handler for endpoint
func Referrals(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var refferalsSQL RefferalsStruct
	var listRefferals []RefferalsStruct
	var pounters RefferalsPounters
	listReponse := []RefferalsPounters{}

	ref, errorReff := sq.Select(
		"code_used.refund_id",
		"money_win",
		"day_used",
		"code_reference.user_id",
		"plans_pay.shop_id",
		"shop.shop_name",
	).
		From("code_used").
		LeftJoin("code_reference on code_reference.code_reference_id = code_used.code_reference_id").
		LeftJoin("plans_pay on plans_pay.plans_id = code_used.plans_id").
		LeftJoin("shop on shop.shop_id = plans_pay.shop_id").
		Where("code_reference.user_id = ? AND code_used.plans_id IS NOT NULL AND code_used.refund_id IS NULL AND code_used.paid_out IS NULL", userID).
		RunWith(database).
		Query()

	if errorReff != nil {
		fmt.Println(errorReff, "Error get refferals")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get refferals"})
		c.SendStatus(400)
		return
	}

	for ref.Next() {
		_ = ref.Scan(
			&refferalsSQL.RefundID,
			&refferalsSQL.MoneyWin,
			&refferalsSQL.DayUsed,
			&refferalsSQL.UserID,
			&refferalsSQL.ShopID,
			&refferalsSQL.ShopName,
		)

		listRefferals = append(listRefferals, refferalsSQL)
	}

	for i := 0; i < len(listRefferals); i++ {
		pounters.RefundID = &listRefferals[i].RefundID.String
		pounters.MoneyWin = &listRefferals[i].MoneyWin.String
		pounters.DayUsed = &listRefferals[i].DayUsed.String
		pounters.UserID = &listRefferals[i].UserID.String
		pounters.ShopID = &listRefferals[i].ShopID.String
		pounters.ShopName = &listRefferals[i].ShopName.String
		listReponse = append(listReponse, pounters)
	}

	response := ResponseRefferals{Referrals: listReponse}

	c.JSON(response)
}

//MyCodeHandler Handler for endpoint
func MyCodeHandler(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var myCodeResponse MyCode

	ErrorCode := sq.Select("code").
		From("code_reference").
		Where("user_id = ?", userID).
		RunWith(database).
		Scan(&myCodeResponse.Code)

	if ErrorCode != nil {
		fmt.Println("Error to get my code", ErrorCode)
		c.JSON(ErrorResponse{MESSAGE: "Error to get code"})
		c.SendStatus(400)
	}

	c.JSON(myCodeResponse)
}

//ReferralsFail Handler for endpoint
func ReferralsFail(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var refferalsSQL RefferalsStruct
	var listRefferals []RefferalsStruct
	var pounters RefferalsPounters
	listReponse := []RefferalsPounters{}

	ref, errorReff := sq.Select(
		"code_used.refund_id",
		"money_win",
		"day_used",
		"code_reference.user_id",
		"plans_pay.shop_id",
		"shop.shop_name",
	).
		From("code_used").
		LeftJoin("code_reference on code_reference.code_reference_id = code_used.code_reference_id").
		LeftJoin("plans_pay on plans_pay.plans_id = code_used.plans_id").
		LeftJoin("shop on shop.shop_id = plans_pay.shop_id").
		Where("code_reference.user_id = ? AND code_used.plans_id IS NOT NULL AND code_used.refund_id IS NOT NULL AND code_used.paid_out IS NULL", userID).
		RunWith(database).
		Query()

	if errorReff != nil {
		fmt.Println(errorReff, "Error get refferals fail")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get refferals fail"})
		c.SendStatus(400)
		return
	}

	for ref.Next() {
		_ = ref.Scan(
			&refferalsSQL.RefundID,
			&refferalsSQL.MoneyWin,
			&refferalsSQL.DayUsed,
			&refferalsSQL.UserID,
			&refferalsSQL.ShopID,
			&refferalsSQL.ShopName,
		)

		listRefferals = append(listRefferals, refferalsSQL)
	}

	for i := 0; i < len(listRefferals); i++ {
		pounters.RefundID = &listRefferals[i].RefundID.String
		pounters.MoneyWin = &listRefferals[i].MoneyWin.String
		pounters.DayUsed = &listRefferals[i].DayUsed.String
		pounters.UserID = &listRefferals[i].UserID.String
		pounters.ShopID = &listRefferals[i].ShopID.String
		pounters.ShopName = &listRefferals[i].ShopName.String
		listReponse = append(listReponse, pounters)
	}

	response := ResponseRefferalsFail{ReferralsFail: listReponse}

	c.JSON(response)
}

//EarnedReferrals Handler for endpoint
func EarnedReferrals(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	var response ResponseEarnedReferrals
	var responsePointer ResponseEarnedReferralsPointer

	ErrorEarned := sq.Select(
		"code_used.code_reference_id",
		"code_reference.user_id",
		"SUM(code_used.money_win) AS money_win",
	).
		From("code_used").
		LeftJoin("code_reference on code_reference.code_reference_id = code_used.code_reference_id").
		Where("code_reference.user_id = ? AND code_used.plans_id IS NOT NULL AND code_used.refund_id IS NULL AND code_used.paid_out IS NOT NULL", userID).
		GroupBy("code_used.code_reference_id").
		RunWith(database).
		Scan(&response.CodeReferenceID, &response.UserID, &response.MoneyWin)

	if ErrorEarned != nil {
		fmt.Println(ErrorEarned, "Error to get earned")
	}

	responsePointer.CodeReferenceID = &response.CodeReferenceID.String
	responsePointer.UserID = &response.UserID.String
	responsePointer.MoneyWin = &response.MoneyWin.String

	if len(*responsePointer.CodeReferenceID) == 0 && len(*responsePointer.UserID) == 0 {
		c.JSON(ResponseEarnedReferralsEmpty{})
	} else {
		fmt.Println("ppp")
		c.JSON(ResponseEarnedReferralsSuccess{EarnedReferrals: responsePointer})
	}

}

//EarnedReferralsMonth Handler for endpoint
func EarnedReferralsMonth(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))
	minDate := 30
	var earnedSQL ResponseEarnedReferrals
	var response ResponseEarnedReferralsPointer

	loc, _ := time.LoadLocation("America/Mexico_City")
	now := time.Now().In(loc)

	y := now.Year()
	m := now.Month().String()
	monthNumber := helper.Month(m)

	if monthNumber == 2 {
		minDate = 28
	}

	dateInit := fmt.Sprintf("%v-%v-1", y, monthNumber)
	dateEnd := fmt.Sprintf("%v-%v-%v", y, monthNumber, minDate)

	ErrorSQL := sq.Select("code_used.code_reference_id", "user_id", "SUM(money_win) AS money_win").
		From("code_used").
		LeftJoin("code_reference on code_reference.code_reference_id=code_used.code_reference_id").
		Where("(day_used BETWEEN ? AND ?) AND code_reference.user_id = ? AND code_used.plans_id IS NOT NULL AND code_used.refund_id IS NULL AND code_used.paid_out IS NULL",
			dateInit,
			dateEnd,
			userID,
		).
		GroupBy("code_used.code_reference_id").
		RunWith(database).
		QueryRow().
		Scan(&earnedSQL.CodeReferenceID, &earnedSQL.UserID, &earnedSQL.MoneyWin)

	if ErrorSQL != nil {
		fmt.Println(ErrorSQL, "Problem to get amount earned")
	}

	response.CodeReferenceID = &earnedSQL.CodeReferenceID.String
	response.UserID = &earnedSQL.UserID.String
	response.MoneyWin = &earnedSQL.MoneyWin.String

	if len(*response.CodeReferenceID) == 0 && len(*response.UserID) == 0 {
		c.JSON(ResponseEarnedReferralsEmpty{})
	} else {
		c.JSON(ResponseEarnedReferralsSuccess{EarnedReferrals: response})
	}
}

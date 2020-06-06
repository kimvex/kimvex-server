package routes

import "database/sql"

//ErrorResponse Basic
//ErrorResponse Structure for response of type error api
type ErrorResponse struct {
	MESSAGE string `json:"message"`
}

//SuccessResponse structure for response of type success api
type SuccessResponse struct {
	MESSAGE string `json:"message"`
}

/*Login*/

//User struct
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserID   string `json:userId`
}

//Profile struct
type ProfileStruct struct {
	UserID   sql.NullString `json:"user_id"`
	Fullname sql.NullString `json:"fullname"`
	Email    sql.NullString `json:"email"`
	Phone    sql.NullString `json:"phone"`
	Age      sql.NullString `json:"age"`
	Gender   sql.NullString `json:"gender"`
	Image    sql.NullString `json:"image"`
	CreateAt sql.NullString `json:"create_at"`
	Code     sql.NullString `json:"code"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

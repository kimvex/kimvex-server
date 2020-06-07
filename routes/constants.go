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

//ProfileStruct struct
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

//RegisterData struct
type RegisterData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Age      string `json:"age"`
	Phone    int    `json:"phone"`
	Gender   string `json:"gender"`
}

//UpdateProfile struct
type UpdateProfile struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Fullname    string `json:"fullname"`
	Age         string `json:"age"`
	Phone       int    `json:"phone"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	URLImage    string `json:"url_image"`
}

//BasicUser struct
type BasicUser struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

//DataRestorePassword struct
type DataRestorePassword struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
	Code        string `json:"code"`
	Email       string `json:"email"`
}

//RefCode struct
type RefCode struct {
	RefferCode string `json:"reffer_code"`
}

//SQLCodeRef struct
type SQLCodeRef struct {
	Code            string `json:"code"`
	CodeReferenceID string `json:"code_reference_id"`
}

//ValidateExistUser struct
type ValidateExistUser struct {
	Email string `json:"email"`
}

//TokenResponse struct
type TokenResponse struct {
	Token string `json:"token"`
}

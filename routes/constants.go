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

//RefferalsStruct struct
type RefferalsStruct struct {
	RefundID sql.NullString `json:"refund_id"`
	MoneyWin sql.NullString `json:"money_win"`
	DayUsed  sql.NullString `json:"day_used"`
	UserID   sql.NullString `json:"user_id"`
	ShopID   sql.NullString `json:"shop_id"`
	ShopName sql.NullString `json:"shop_name"`
}

//RefferalsPounters struct
type RefferalsPounters struct {
	RefundID *string `json:"refund_id"`
	MoneyWin *string `json:"money_win"`
	DayUsed  *string `json:"day_used"`
	UserID   *string `json:"user_id"`
	ShopID   *string `json:"shop_id"`
	ShopName *string `json:"shop_name"`
}

//ResponseRefferals struct
type ResponseRefferals struct {
	Referrals []RefferalsPounters `json:"referrals"`
}

//MyCode struct
type MyCode struct {
	Code string `json:"code"`
}

//ResponseRefferalsFail struct
type ResponseRefferalsFail struct {
	ReferralsFail []RefferalsPounters `json:"reffers_fail"`
}

//ResponseEarnedReferrals struct
type ResponseEarnedReferrals struct {
	CodeReferenceID sql.NullString `json:"code_reference_id"`
	UserID          sql.NullString `json:"user_id"`
	MoneyWin        sql.NullString `json:"money_win"`
}

//ResponseEarnedReferralsPointer struct
type ResponseEarnedReferralsPointer struct {
	CodeReferenceID *string `json:"code_reference_id"`
	UserID          *string `json:"user_id"`
	MoneyWin        *string `json:"money_win"`
}

//ResponseEarnedReferralsSuccess atruct
type ResponseEarnedReferralsSuccess struct {
	EarnedReferrals ResponseEarnedReferralsPointer `json:"earned_referrals"`
}

//Empty struct
type Empty struct {
}

//ResponseEarnedReferralsEmpty atruct
type ResponseEarnedReferralsEmpty struct {
	EarnedReferrals Empty `json:"earned_referrals"`
}

//SQLGetShop struct
type SQLGetShop struct {
	ShopID           sql.NullString `json:"shop_id"`
	ShopName         sql.NullString `json:"shop_name"`
	Address          sql.NullString `json:"address"`
	Phone            sql.NullString `json:"phone"`
	Phone2           sql.NullString `json:"phone2"`
	Description      sql.NullString `json:"description"`
	CoverImage       sql.NullString `json:"cover_image"`
	AcceptCard       sql.NullString `json:"accept_card"`
	ListCards        sql.NullString `json:"list_cards"`
	Lat              sql.NullString `json:"lat"`
	Lon              sql.NullString `json:"lon"`
	ScoreShop        sql.NullString `json:"score_shop"`
	Status           sql.NullString `json:"status"`
	Logo             sql.NullString `json:"logo"`
	ServiceTypeID    sql.NullString `json:"service_type_id"`
	SubServiceTypeID sql.NullString `json:"sub_service_type_id"`
	LUN              sql.NullString `json:"LUN"`
	MAR              sql.NullString `json:"MAR"`
	MIE              sql.NullString `json:"MIE"`
	JUE              sql.NullString `json:"JUE"`
	VIE              sql.NullString `json:"VIE"`
	SAB              sql.NullString `json:"SAB"`
	DOM              sql.NullString `json:"DOM"`
	UserID           sql.NullString `json:"user_id"`
	Images           sql.NullString `json:"images"`
}

//ShopPointerGet struct
type ShopPointerGet struct {
	ShopID           *string  `json:"shop_id"`
	ShopName         *string  `json:"shop_name"`
	Address          *string  `json:"address"`
	Phone            *string  `json:"phone"`
	Phone2           *string  `json:"phone2"`
	Description      *string  `json:"description"`
	CoverImage       *string  `json:"cover_image"`
	AcceptCard       *string  `json:"accept_card"`
	ListCards        []string `json:"list_cards"`
	Lat              *string  `json:"lat"`
	Lon              *string  `json:"lon"`
	ScoreShop        *string  `json:"score_shop"`
	Status           *string  `json:"status"`
	Logo             *string  `json:"logo"`
	ServiceTypeID    *string  `json:"service_type_id"`
	SubServiceTypeID *string  `json:"sub_service_type_id"`
	LUN              *string  `json:"LUN"`
	MAR              *string  `json:"MAR"`
	MIE              *string  `json:"MIE"`
	JUE              *string  `json:"JUE"`
	VIE              *string  `json:"VIE"`
	SAB              *string  `json:"SAB"`
	DOM              *string  `json:"DOM"`
	UserID           *string  `json:"user_id"`
	Images           []string `json:"images"`
}

//Paginations struct for generate paginations.
type Paginations struct {
	Limit string `query:"limit"`
	Page  string `query:"page"`
}

//MyShops structure for build response of shops of a profile
type MyShops struct {
	ShopID         sql.NullString `json:"shop_id"`
	ShopName       sql.NullString `json:"shop_name"`
	Address        sql.NullString `json:"address"`
	Phone          sql.NullString `json:"phone"`
	Phone2         sql.NullString `json:"phone2"`
	Description    sql.NullString `json:"description"`
	CoverImage     sql.NullString `json:"cover_image"`
	AcceptCard     sql.NullString `json:"accept_card"`
	ListCards      sql.NullString `json:"list_cards"`
	Lat            sql.NullString `json:"lat"`
	Lon            sql.NullString `json:"lon"`
	ScoreShop      sql.NullString `json:"score_shop"`
	Status         sql.NullString `json:"status"`
	LockShop       sql.NullString `json:"lock_shop"`
	Canceled       sql.NullString `json:"canceled"`
	ServiceName    sql.NullString `json:"service_name"`
	SubServiceName sql.NullString `json:"sub_service_name"`
	LUN            sql.NullString `json:"LUN"`
	MAR            sql.NullString `json:"MAR"`
	MIE            sql.NullString `json:"MIE"`
	JUE            sql.NullString `json:"JUE"`
	VIE            sql.NullString `json:"VIE"`
	SAB            sql.NullString `json:"SAB"`
	DOM            sql.NullString `json:"DOM"`
	UserID         sql.NullString `json:"user_id"`
	Images         sql.NullString `json:"images"`
	DateInit       sql.NullString `json:"date_init"`
	DateFinish     sql.NullString `json:"date_finish"`
	TypeCharge     sql.NullString `json:"type_charge"`
}

//MyShopsPoints structure for build response of shops of a profile
type MyShopsPoints struct {
	ShopID         *string `json:"shop_id"`
	ShopName       *string `json:"shop_name"`
	Address        *string `json:"address"`
	Phone          *string `json:"phone"`
	Phone2         *string `json:"phone2"`
	Description    *string `json:"description"`
	CoverImage     *string `json:"cover_image"`
	AcceptCard     *string `json:"accept_card"`
	ListCards      *string `json:"list_cards"`
	Lat            *string `json:"lat"`
	Lon            *string `json:"lon"`
	ScoreShop      *string `json:"score_shop"`
	Status         *string `json:"status"`
	LockShop       *string `json:"lock_shop"`
	Canceled       *string `json:"canceled"`
	ServiceName    *string `json:"service_name"`
	SubServiceName *string `json:"sub_service_name"`
	LUN            *string `json:"LUN"`
	MAR            *string `json:"MAR"`
	MIE            *string `json:"MIE"`
	JUE            *string `json:"JUE"`
	VIE            *string `json:"VIE"`
	SAB            *string `json:"SAB"`
	DOM            *string `json:"DOM"`
	UserID         *string `json:"user_id"`
	Images         *string `json:"images"`
	DateInit       *string `json:"date_init"`
	DateFinish     *string `json:"date_finish"`
	TypeCharge     *string `json:"type_charge"`
}

//ResponseResult for response result of arrays
type ResponseResult struct {
	Result []MyShopsPoints `json:"result"`
}

//ParamsShopOffers for parse shopid
type ParamsShopOffers struct {
	Status string `json:"status"`
	Limit  string `json:"limit"`
	Page   string `json:"page"`
}

//ResponseListOffersSQL of SQL
type ResponseListOffersSQL struct {
	OffersID    sql.NullString `json:"offers_id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	DateInit    sql.NullString `json:"date_init"`
	DateEnd     sql.NullString `json:"date_end"`
	ImageURL    sql.NullString `json:"image_url"`
	Active      sql.NullString `json:"active"`
	Lat         sql.NullString `json:"lat"`
	Lon         sql.NullString `json:"lon"`
}

//ResponseListOffers of json
type ResponseListOffers struct {
	OffersID    *string `json:"offers_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DateInit    *string `json:"date_init"`
	DateEnd     *string `json:"date_end"`
	ImageURL    *string `json:"image_url"`
	Active      *string `json:"active"`
	Lat         *string `json:"lat"`
	Lon         *string `json:"lon"`
}

//ResponseResultOffers of json
type ResponseResultOffers struct {
	Offers []ResponseListOffers `json:"offers"`
}

//AOffer return of SQL
type AOffer struct {
	OffersID    sql.NullString `json:"offers_id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	DateEnd     sql.NullString `json:"date_end"`
	ImageURL    sql.NullString `json:"image_url"`
	Active      sql.NullString `json:"active"`
	Lat         sql.NullString `json:"lat"`
	Lon         sql.NullString `json:"lon"`
	ShopID      sql.NullString `json:"shop_id"`
	ShopName    sql.NullString `json:"shop_name"`
	CoverImage  sql.NullString `json:"cover_image"`
}

//AOfferPointer of setOffer response
type AOfferPointer struct {
	OffersID    *string `json:"offers_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DateEnd     *string `json:"date_end"`
	ImageURL    *string `json:"image_url"`
	Active      *string `json:"active"`
	Lat         *string `json:"lat"`
	Lon         *string `json:"lon"`
	ShopID      *string `json:"shop_id"`
	ShopName    *string `json:"shop_name"`
	CoverImage  *string `json:"cover_image"`
}

//ResponseInforOffer for response to the server
type ResponseInforOffer struct {
	Offer AOfferPointer `json:"offer"`
}

//ValidateExistUser struct
type ValidateExistUser struct {
	Email string `json:"email"`
}

//TokenResponse struct
type TokenResponse struct {
	Token string `json:"token"`
}

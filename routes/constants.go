package routes

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Phone       string `json:"phone"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	URLImage    string `json:"image_url"`
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

//ResponseValidate struct
type ResponseValidate struct {
	Validate SQLCodeRef `json:"validate"`
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
	ShopID         *string  `json:"shop_id"`
	ShopName       *string  `json:"shop_name"`
	Address        *string  `json:"address"`
	Phone          *string  `json:"phone"`
	Phone2         *string  `json:"phone2"`
	Description    *string  `json:"description"`
	CoverImage     *string  `json:"cover_image"`
	AcceptCard     *string  `json:"accept_card"`
	ListCards      *string  `json:"list_cards"`
	Lat            *string  `json:"lat"`
	Lon            *string  `json:"lon"`
	ScoreShop      *string  `json:"score_shop"`
	Status         *string  `json:"status"`
	LockShop       *string  `json:"lock_shop"`
	Canceled       *string  `json:"canceled"`
	ServiceName    *string  `json:"service_name"`
	SubServiceName *string  `json:"sub_service_name"`
	LUN            *string  `json:"LUN"`
	MAR            *string  `json:"MAR"`
	MIE            *string  `json:"MIE"`
	JUE            *string  `json:"JUE"`
	VIE            *string  `json:"VIE"`
	SAB            *string  `json:"SAB"`
	DOM            *string  `json:"DOM"`
	UserID         *string  `json:"user_id"`
	Images         []string `json:"images"`
	DateInit       *string  `json:"date_init"`
	DateFinish     *string  `json:"date_finish"`
	TypeCharge     *string  `json:"type_charge"`
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

//ServiceSQL for get SQL
type ServiceSQL struct {
	ServiceTypeID sql.NullString `json:"service_type_id"`
	ServiceName   sql.NullString `json:"service_name"`
}

//ServiceSQLPointer for get pointer
type ServiceSQLPointer struct {
	ServiceTypeID *string `json:"service_type_id"`
	ServiceName   *string `json:"service_name"`
}

//ServiceResponse of results
type ServiceResponse struct {
	Services []ServiceSQLPointer `json:"services"`
}

//SubServiceSQL for get SQL
type SubServiceSQL struct {
	SubServiceTypeID sql.NullString `json:"sub_service_type_id"`
	SubServiceName   sql.NullString `json:"sub_service_name"`
}

//SubServiceSQLPointer for get pointer
type SubServiceSQLPointer struct {
	SubServiceTypeID *string `json:"sub_service_type_id"`
	SubServiceName   *string `json:"sub_service_name"`
}

//SubServiceResponse of results
type SubServiceResponse struct {
	SubService []SubServiceSQLPointer `json:"sub_service"`
}

//CommentsSQL of comments to SQL
type CommentsSQL struct {
	Comment      sql.NullString `json:"comment"`
	CreateDateAt sql.NullString `json:"create_date_at"`
	UserID       sql.NullString `json:"user_id"`
	Fullname     sql.NullString `json:"fullname"`
	Image        sql.NullString `json:"image"`
}

//CommentsSQL of comments to SQL
type CommentsPointerStruct struct {
	Comment      *string `json:"comment"`
	CreateDateAt *string `json:"create_date_at"`
	UserID       *string `json:"user_id"`
	Fullname     *string `json:"fullname"`
	Image        *string `json:"image"`
}

//ScoreSQL of score to SQL
type ScoreSQL struct {
	Score sql.NullString `json:"score"`
}

//ResponseScore response
type ResponseScore struct {
	Score *string `json:"score"`
}

//PageSQL of SQL
type PageSQL struct {
	Active           sql.NullString `json:"active"`
	TemplateType     sql.NullString `json:"template_type"`
	StyleSheets      sql.NullString `json:"style_sheets"`
	ActiveDays       sql.NullString `json:"active_days"`
	ImagesDays       sql.NullString `json:"images_days"`
	OffersActive     sql.NullString `json:"offers_active"`
	AcceptCardActive sql.NullString `json:"accept_card_active"`
	Subdomain        sql.NullString `json:"subdomain"`
	Domain           sql.NullString `json:"domain"`
	ShopID           sql.NullString `json:"shop_id"`
	PagesID          sql.NullString `json:"pages_id"`
	TypeCharge       sql.NullString `json:"type_charge"`
	ShopName         sql.NullString `json:"shop_name"`
	Description      sql.NullString `json:"description"`
	CoverImage       sql.NullString `json:"cover_image"`
	Logo             sql.NullString `json:"logo"`
	UserID           sql.NullString `json:"user_id"`
}

//PagesPointer of pointer
type PagesPointer struct {
	Active           int     `json:"active"`
	TemplateType     int     `json:"template_type"`
	StyleSheets      int     `json:"style_sheets"`
	ActiveDays       int     `json:"active_days"`
	ImagesDays       int     `json:"images_days"`
	OffersActive     int     `json:"offers_active"`
	AcceptCardActive int     `json:"accept_card_active"`
	Subdomain        *string `json:"subdomain"`
	Domain           *string `json:"domain"`
	ShopID           *string `json:"shop_id"`
	PagesID          *string `json:"pages_id"`
	TypeCharge       *string `json:"type_charge"`
	ShopName         *string `json:"shop_name"`
	Description      *string `json:"description"`
	CoverImage       *string `json:"cover_image"`
	Logo             *string `json:"logo"`
	UserID           *string `json:"user_id"`
}

//PageInformation struct
type PageInformation struct {
	Page PagesPointer `json:"page"`
}

//Location struct for locations
type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

//FShops Struct for shops
type FShops struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ShopID      string             `bson:"shop_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Location    Location           `bson:"location,omitempty"`
	Category    string             `bson:"category,omitempty"`
	SubCategory string             `bson:"sub_category,omitempty"`
	status      bool               `bson:"status,omitempty"`
	Distance    float64            `bson:"distance,omitempty"`
}

//FindShopQuery struct
type FindShopQuery struct {
	Limit       string `json:"limit"`
	LastShopID  string `json:"last_shop_id"`
	MinDistance string `json:"minDistance"`
	Category    string `json:"category"`
}

//ShopsID struct
type ShopsID struct {
	ShopID   string  `json:"shop_id"`
	Distance float64 `json:"distance"`
}

//ShopFromSQLFind struct
type ShopFromSQLFind struct {
	ShopID         sql.NullString `json:"shop_id"`
	ShopName       sql.NullString `json:"shop_name"`
	Address        sql.NullString `json:"address"`
	Phone          sql.NullString `json:"phone"`
	ScoreShop      sql.NullString `json:"score_shop"`
	CoverImage     sql.NullString `json:"cover_image"`
	ServiceName    sql.NullString `json:"service_name"`
	SubServiceName sql.NullString `json:"sub_service_name"`
}

//ShopFindPointer struct
type ShopFindPointer struct {
	ShopID         *string `json:"shop_id"`
	ShopName       *string `json:"shop_name"`
	Address        *string `json:"address"`
	Phone          *string `json:"phone"`
	ScoreShop      *string `json:"score_shop"`
	CoverImage     *string `json:"cover_image"`
	ServiceName    *string `json:"service_name"`
	SubServiceName *string `json:"sub_service_name"`
	Distance       float64 `json:"distance"`
}

//ResponseFinalFindShops struct for response list of shops
type ResponseFinalFindShops struct {
	Shop         []ShopFindPointer `json:"shop"`
	LastShopID   *string           `json:"last_shop_id"`
	LastDistance float64           `json:"last_distance"`
}

//FOffers structure for list offers
type FOffers struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	OfferID  string             `bson:"offer_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Location Location           `bson:"location,omitempty"`
	DateInit string             `bson:"date_init,omitempty"`
	DateEnd  string             `bson:"date_end,omitempty"`
	Status   bool               `bson:"status,omitempty"`
	Distance float64            `bson:"distance,omitempty"`
}

// OffersID for get id in mongo
type OffersID struct {
	OfferID  string  `json:"offer_id"`
	Distance float64 `json:"distance"`
}

//OffersFromSQLFind struct for get offers of mysql
type OffersFromSQLFind struct {
	OfferID     sql.NullString `json:"offer_id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	DateInit    sql.NullString `json:"date_init"`
	DateEnd     sql.NullString `json:"date_end"`
	ImageURL    sql.NullString `json:"image_url"`
	ShopID      sql.NullString `json:"shop_id"`
	ShopName    sql.NullString `json:"shop_name"`
	CoverImage  sql.NullString `json:"cover_image"`
}

//OffersFindPointer struct for get offers of mysql
type OffersFindPointer struct {
	OfferID     *string `json:"offer_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DateInit    *string `json:"date_init"`
	DateEnd     *string `json:"date_end"`
	ImageURL    *string `json:"image_url"`
	ShopID      *string `json:"shop_id"`
	ShopName    *string `json:"shop_name"`
	CoverImage  *string `json:"cover_image"`
	Distance    float64 `json:"distance"`
}

//FindOfferQuery struct
type FindOfferQuery struct {
	Limit       string `json:"limit"`
	LastOfferID string `json:"last_offer_id"`
	MinDistance string `json:"minDistance"`
}

//ResponseFinalFindOffers struct for response list of Offers
type ResponseFinalFindOffers struct {
	Offers       []OffersFindPointer `json:"offers"`
	LastOfferID  *string             `json:"last_offer_id"`
	LastDistance float64             `json:"last_distance"`
}

//IsOwnerShop struct for validate if a user is owner of a shop
type IsOwnerShop struct {
	ShopID sql.NullString `json:"shop_id"`
}

//LocationSQL struct for location SQL
type LocationSQL struct {
	Lat sql.NullString `json:"lat"`
	Lon sql.NullString `json:"lon"`
}

//SuccessResponseOffer structure for response created offer
type SuccessResponseOffer struct {
	OfferID string `json:"offer_id"`
	MESSAGE string `json:"message"`
	Status  int    `json:"status"`
}

//SuccessResponseOfferStatus structure for response offer
type SuccessResponseOfferStatus struct {
	MESSAGE string `json:"message"`
	Status  int    `json:"status"`
}

//QueryParamsOffer for get offers
type QueryParamsOffer struct {
	ShopID      string `json:"shop_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateInit    string `json:"date_init"`
	DateEnd     string `json:"date_end"`
	ImageURL    string `json:"image_url"`
}

//QueryParamsOfferUpdate struct for update offer
type QueryParamsOfferUpdate struct {
	ShopID      string `json:"shop_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateInit    string `json:"date_init"`
	DateEnd     string `json:"date_end"`
	ImageURL    string `json:"image_url"`
	Active      int    `json:"active"`
}

//CommentStruct for comment set
type CommentStruct struct {
	Comment string `json:"comment"`
}

//ScoreStruct for score set
type ScoreStruct struct {
	Score int `json:"score"`
}

//PagePut struct for infomation page
type PagePut struct {
	TemplateType     *int   `json:"template_type"`
	StyleSheets      *int   `json:"style_sheets"`
	ActiveDays       *int   `json:"active_days"`
	ImagesDays       *int   `json:"images_days"`
	OffersActive     *int   `json:"offers_active"`
	AcceptCardActive *int   `json:"accept_card_active"`
	Subdomain        string `json:"subdomain"`
	Domain           string `json:"domain"`
}

//ValidateDomains struct for validate domains and subdomains
type ValidateDomains struct {
	ShopID    sql.NullString `json:"shop_id"`
	Subdomain sql.NullString `json:"subdomain"`
	Domain    sql.NullString `json:"domain"`
}

//ResponseSubdomain for response of validate subdomain
type ResponseSubdomain struct {
	Subdomain bool `json:"subdomain"`
}

//ResponseDomain for response of validate domain
type ResponseDomain struct {
	Domain bool `json:"domain"`
}

//ResponseResultSimple struct for response simple
type ResponseResultSimple struct {
	Result string `json:"result"`
}

//DataShop struct for information shop
type DataShop struct {
	ShopName         string   `json:"shop_name"`
	Address          string   `json:"address"`
	Phone            string   `json:"phone"`
	Phone2           string   `json:"phone2"`
	Description      string   `json:"description"`
	CoverImage       string   `json:"cover_image"`
	Logo             string   `json:"logo"`
	AcceptCard       bool     `json:"accept_card"`
	ListCards        []string `json:"list_cards"`
	ShopSchedules    []string `json:"shop_schedules"`
	Lat              float64  `json:"lat"`
	Lon              float64  `json:"lon"`
	ServiceTypeID    string   `json:"service_type_id"`
	SubServiceTypeID string   `json:"sub_service_type_id"`
	ListImages       []string `json:"list_images"`
}

//DataShopString struct for information shop
type DataShopString struct {
	ShopName         string   `json:"shop_name"`
	Address          string   `json:"address"`
	Phone            string   `json:"phone"`
	Phone2           string   `json:"phone2"`
	Description      string   `json:"description"`
	CoverImage       string   `json:"cover_image"`
	Logo             string   `json:"logo"`
	AcceptCard       bool     `json:"accept_card"`
	ListCards        []string `json:"list_cards"`
	ShopSchedules    []string `json:"shop_schedules"`
	Lat              string   `json:"lat"`
	Lon              string   `json:"lon"`
	ServiceTypeID    string   `json:"service_type_id"`
	SubServiceTypeID string   `json:"sub_service_type_id"`
	ListImages       []string `json:"list_images"`
}

//ServiceNames struct for services
type ServiceNames struct {
	SubServiceName sql.NullString `json:"sub_service_name"`
	ServiceName    sql.NullString `json:"service_name"`
}

//ResponseCreateShop struct for response shop created
type ResponseCreateShop struct {
	Message string `json:"message"`
	ShopID  int64  `json:"shop_id"`
	Status  int    `json:"status"`
}

//ResponseStatusCode struct from response
type ResponseStatusCode struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

//DataDeleteImage struct for delete image
type DataDeleteImage struct {
	URLImage string `json:"url_image"`
}

//ResponsePointerLasts struct for conver to *
type ResponsePointerLasts struct {
	LastID   *string `json:"last_id"`
	Distance float64 `json:"distance"`
}

// FromSQLHallways list of hallways of shop
type FromSQLHallways struct {
	HallwaysID  sql.NullString `json:"hallways_id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
}

//SQLArticles list of articles of hallways
type SQLArticles struct {
	Name         sql.NullString `json:"name"`
	Description  sql.NullString `json:"description"`
	Price        sql.NullInt32  `json:"price"`
	CountArticle sql.NullInt32  `json:"count_article"`
}

//ArticlesPointer struct for articles
type ArticlesPointer struct {
	Name         *string `json:"name"`
	Description  *string `json:"description"`
	Price        *int32  `json:"price"`
	CountArticle *int32  `json:"count_article"`
}

//SQLHallwaysArticle struct for response
type SQLHallwaysArticle struct {
	Name        *string           `json:"name"`
	Description *string           `json:"description"`
	Articles    []ArticlesPointer `json:"articles"`
}

//HallwaysResponse struct for response
type HallwaysResponse struct {
	Hallways []SQLHallwaysArticle `json:"hallways"`
}

//HallwaysPostStruct struct for insert
type HallwaysPostStruct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//ValidateExistUser struct
type ValidateExistUser struct {
	Email string `json:"email"`
}

//TokenResponse struct
type TokenResponse struct {
	Token string `json:"token"`
}

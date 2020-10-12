package routes

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"../helper"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Shops namespace
func Shops() {
	apiRouteShop := apiRoute.Group("/shop")
	apiRouteImages := apiRoute.Group("/images")
	apiRouteProfile := apiRoute.Group("/profile")
	apiRouteBase := apiRoute.Group("/")

	//Validations token
	apiRouteProfile.Use("/shops", ValidateRoute)
	apiRouteShop.Use("/:shop_id/offers", ValidateRoute)
	apiRouteShop.Use("/offer/:offer_id", ValidateRoute)
	apiRouteShop.Use("/offers", ValidateRoute)
	apiRouteShop.Use("/offer/:offer_id", ValidateRoute)
	apiRouteShop.Use("/:shop_id/comment", ValidateRoute)
	apiRouteShop.Use("/:shop_id/score", ValidateRoute)
	apiRouteShop.Use("/:shop_id/score/:user_id", ValidateRoute)
	apiRouteShop.Use("/lock/:shop_id", ValidateRoute)
	apiRouteShop.Use("/unlock/:shop_id", ValidateRoute)
	apiRouteShop.Use("/:shop_id/update_page/:page_id", ValidateRoute)
	apiRouteShop.Use("/:shop_id/active_page/:page_id", ValidateRoute)
	apiRouteShop.Use("/:shop_id/deactivate_page/:page_id", ValidateRoute)
	apiRouteImages.Use("/shop", ValidateRoute)

	apiRouteShop.Get("/:shop_id", ShopGet)
	apiRouteShop.Get("/:shop_id/offers", ShopOffers)
	apiRouteShop.Get("/offer/:offer_id", OfferInfo)
	apiRouteShop.Get("/:shop_id/comments", Comments)
	apiRouteShop.Get("/:shop_id/score/:user_id", Score)
	apiRouteShop.Get("/:shop_id/page", Page)
	apiRouteBase.Get("/find/shops/:lat/:lon", FindShops)
	apiRouteBase.Get("/shops/offers/:lat/:lon", FindOffers)

	apiRouteBase.Get("/services", Services)
	apiRouteBase.Get("/sub_service/:service_id", SubServices)

	apiRouteProfile.Get("/shops", ProfileShop)
	apiRouteShop.Post("/:shop_id/comment", Comment)
	apiRouteShop.Post("/:shop_id/score", SetScore)
	apiRouteShop.Put("/:shop_id/score/:user_id", UpdateScore)
	apiRouteShop.Put("/lock/:shop_id", LockShop)
	apiRouteShop.Put("/unlock/:shop_id", UnlockShop)
	apiRouteShop.Put("/:shop_id/update_page/:page_id", UpdatePage)
	apiRouteShop.Put("/:shop_id/active_page/:page_id", ActivePage)
	apiRouteShop.Put("/:shop_id/deactivate_page/:page_id", DeactivePage)
	apiRouteImages.Post("/shop", UploadImages)

	apiRouteShop.Post("/offers", CreateOffer)
	apiRouteShop.Put("/offers/:offer_id", UpdateOffer)

}

//ShopGet Handler for endpoint
func ShopGet(c *fiber.Ctx) {
	shopID := c.Params("shop_id")

	var SQLResponse SQLGetShop
	var response ShopPointerGet

	shopResultsError := sq.Select(
		"shop.shop_id",
		"shop.shop_name",
		"shop.address",
		"shop.phone",
		"shop.phone2",
		"shop.description",
		"shop.cover_image",
		"shop.accept_card",
		"shop.list_cards",
		"shop.lat",
		"shop.lon",
		"shop.score_shop",
		"shop.status",
		"shop.logo",
		"shop.service_type_id",
		"shop.sub_service_type_id",
		"shop_schedules.LUN",
		"shop_schedules.MAR",
		"shop_schedules.MIE",
		"shop_schedules.JUE",
		"shop_schedules.VIE",
		"shop_schedules.SAB",
		"shop_schedules.DOM",
		"usersk.user_id",
		"GROUP_CONCAT(coalesce(images_shop.url_image, '')) AS images",
	).
		From("shop").
		LeftJoin("images_shop on images_shop.shop_id = shop.shop_id").
		LeftJoin("shop_schedules on shop_schedules.shop_id = shop.shop_id").
		LeftJoin("usersk on usersk.user_id = shop.user_id").
		Where("shop.shop_id = ?", shopID).
		GroupBy(
			"shop.shop_id",
			"shop.shop_name",
			"shop.address",
			"shop.phone",
			"shop.phone2",
			"shop.description",
			"shop.cover_image",
			"shop.accept_card",
			"shop.list_cards",
			"shop.lat",
			"shop.lon",
			"shop.score_shop",
			"shop.status",
			"shop.logo",
			"shop.service_type_id",
			"shop.sub_service_type_id",
			"shop_schedules.LUN",
			"shop_schedules.MAR",
			"shop_schedules.MIE",
			"shop_schedules.JUE",
			"shop_schedules.VIE",
			"shop_schedules.SAB",
			"shop_schedules.DOM",
			"usersk.user_id",
		).
		RunWith(database).
		QueryRow().
		Scan(
			&SQLResponse.ShopID,
			&SQLResponse.ShopName,
			&SQLResponse.Address,
			&SQLResponse.Phone,
			&SQLResponse.Phone2,
			&SQLResponse.Description,
			&SQLResponse.CoverImage,
			&SQLResponse.AcceptCard,
			&SQLResponse.ListCards,
			&SQLResponse.Lat,
			&SQLResponse.Lon,
			&SQLResponse.ScoreShop,
			&SQLResponse.Status,
			&SQLResponse.Logo,
			&SQLResponse.ServiceTypeID,
			&SQLResponse.SubServiceTypeID,
			&SQLResponse.LUN,
			&SQLResponse.MAR,
			&SQLResponse.MIE,
			&SQLResponse.JUE,
			&SQLResponse.VIE,
			&SQLResponse.SAB,
			&SQLResponse.DOM,
			&SQLResponse.UserID,
			&SQLResponse.Images,
		)

	if shopResultsError != nil {
		fmt.Println(shopResultsError, "Error get Shops")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get shops"})
		c.SendStatus(400)
		return
	}

	response.ShopID = &SQLResponse.ShopID.String
	response.ShopName = &SQLResponse.ShopName.String
	response.Address = &SQLResponse.Address.String
	response.Phone = &SQLResponse.Phone.String
	response.Phone2 = &SQLResponse.Phone2.String
	response.Description = &SQLResponse.Description.String
	response.CoverImage = &SQLResponse.CoverImage.String
	response.AcceptCard = &SQLResponse.AcceptCard.String
	response.Lat = &SQLResponse.Lat.String
	response.Lon = &SQLResponse.Lon.String
	response.ScoreShop = &SQLResponse.ScoreShop.String
	response.Status = &SQLResponse.Status.String
	response.Logo = &SQLResponse.Logo.String
	response.ServiceTypeID = &SQLResponse.ServiceTypeID.String
	response.SubServiceTypeID = &SQLResponse.SubServiceTypeID.String
	response.LUN = &SQLResponse.LUN.String
	response.MAR = &SQLResponse.MAR.String
	response.MIE = &SQLResponse.MIE.String
	response.JUE = &SQLResponse.JUE.String
	response.VIE = &SQLResponse.VIE.String
	response.SAB = &SQLResponse.SAB.String
	response.DOM = &SQLResponse.DOM.String
	response.UserID = &SQLResponse.UserID.String

	ListCardsConverter := &SQLResponse.ListCards.String
	ListCardsSimple := strings.Replace(*ListCardsConverter, "[", "", -1)
	ListCardsSimple = strings.Replace(ListCardsSimple, "]", "", -1)
	ListCardsSimple = strings.Replace(ListCardsSimple, "\"", "", -1)
	ListCardsSimple = strings.Replace(ListCardsSimple, "]", "", -1)
	ListCards := strings.Split(ListCardsSimple, ",")

	for i := 0; i < len(ListCards); i++ {
		response.ListCards = append(response.ListCards, ListCards[i])
	}

	ListImagesConverter := &SQLResponse.Images.String
	Images := strings.Split(*ListImagesConverter, ",")

	for i := 0; i < len(Images); i++ {
		response.Images = append(response.Images, Images[i])
	}

	c.JSON(response)
}

//ProfileShop handler for get profile shops
func ProfileShop(c *fiber.Ctx) {
	userID := userIDF(c.Get("token"))

	var myShops MyShops
	MyShopUnStructured := []MyShops{}
	var finalMyShops MyShopsPoints
	listReponse := []MyShopsPoints{}

	Pagination := new(Paginations)
	if err := c.QueryParser(Pagination); err != nil {
		fmt.Println(err, "Error parsing shops")
	}

	page, _ := strconv.Atoi(Pagination.Page)
	limit, _ := strconv.Atoi(Pagination.Limit)
	offset := limit * page

	shops, err := sq.Select(
		"shop.shop_id",
		"shop.shop_name",
		"shop.address",
		"shop.phone",
		"shop.phone2",
		"shop.description",
		"shop.cover_image",
		"shop.accept_card",
		"shop.list_cards",
		"lat",
		"lon",
		"score_shop",
		"shop.status",
		"lock_shop",
		"canceled",
		"service_name",
		"sub_service_name",
		"shop_schedules.LUN",
		"shop_schedules.MAR",
		"shop_schedules.MIE",
		"shop_schedules.JUE",
		"shop_schedules.VIE",
		"shop_schedules.SAB",
		"shop_schedules.DOM",
		"usersk.user_id",
		"GROUP_CONCAT(coalesce(images_shop.url_image, '')) AS images",
		"date_init",
		"date_finish",
		"type_charge",
	).
		From("shop").
		LeftJoin("images_shop on images_shop.shop_id = shop.shop_id").
		LeftJoin("service_type on service_type.service_type_id = shop.service_type_id").
		LeftJoin("sub_service_type on sub_service_type.sub_service_type_id = shop.sub_service_type_id").
		LeftJoin("shop_schedules on shop_schedules.shop_id = shop.shop_id").
		LeftJoin("plans_pay on plans_pay.shop_id = shop.shop_id").
		LeftJoin("usersk on usersk.user_id = shop.user_id").
		Where("shop.user_id = ? AND (plans_pay.expired = ? OR plans_pay.expired IS NULL)", userID, 0).
		OrderBy("shop_id DESC").
		GroupBy("shop.shop_id, shop_schedules.LUN, shop_schedules.MAR, shop_schedules.MAR, shop_schedules.MIE, shop_schedules.JUE, shop_schedules.VIE, shop_schedules.SAB, shop_schedules.DOM, plans_pay.date_init, plans_pay.date_finish, plans_pay.type_charge").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		RunWith(database).
		Query()

	if err != nil {
		fmt.Println(err, "Error to get shops")
		ErrorProfile := ErrorResponse{MESSAGE: "Error to get shops"}
		c.JSON(ErrorProfile)
		c.Status(400)
		return
	}

	for shops.Next() {
		_ = shops.Scan(
			&myShops.ShopID,
			&myShops.ShopName,
			&myShops.Address,
			&myShops.Phone,
			&myShops.Phone2,
			&myShops.Description,
			&myShops.CoverImage,
			&myShops.AcceptCard,
			&myShops.ListCards,
			&myShops.Lat,
			&myShops.Lon,
			&myShops.ScoreShop,
			&myShops.Status,
			&myShops.LockShop,
			&myShops.Canceled,
			&myShops.ServiceName,
			&myShops.SubServiceName,
			&myShops.LUN,
			&myShops.MAR,
			&myShops.MIE,
			&myShops.JUE,
			&myShops.VIE,
			&myShops.SAB,
			&myShops.DOM,
			&myShops.UserID,
			&myShops.Images,
			&myShops.DateInit,
			&myShops.DateFinish,
			&myShops.TypeCharge,
		)

		MyShopUnStructured = append(MyShopUnStructured, myShops)
	}

	for i := 0; i < len(MyShopUnStructured); i++ {
		finalMyShops.ShopID = &MyShopUnStructured[i].ShopID.String
		finalMyShops.ShopName = &MyShopUnStructured[i].ShopName.String
		finalMyShops.Address = &MyShopUnStructured[i].Address.String
		finalMyShops.Phone = &MyShopUnStructured[i].Phone.String
		finalMyShops.Phone2 = &MyShopUnStructured[i].Phone2.String
		finalMyShops.Description = &MyShopUnStructured[i].Description.String
		finalMyShops.CoverImage = &MyShopUnStructured[i].CoverImage.String
		finalMyShops.AcceptCard = &MyShopUnStructured[i].AcceptCard.String
		finalMyShops.ListCards = &MyShopUnStructured[i].ListCards.String
		finalMyShops.Lat = &MyShopUnStructured[i].Lat.String
		finalMyShops.Lon = &MyShopUnStructured[i].Lon.String
		finalMyShops.ScoreShop = &MyShopUnStructured[i].ScoreShop.String
		finalMyShops.Status = &MyShopUnStructured[i].Status.String
		finalMyShops.LockShop = &MyShopUnStructured[i].LockShop.String
		finalMyShops.Canceled = &MyShopUnStructured[i].Canceled.String
		finalMyShops.ServiceName = &MyShopUnStructured[i].ServiceName.String
		finalMyShops.SubServiceName = &MyShopUnStructured[i].SubServiceName.String
		finalMyShops.LUN = &MyShopUnStructured[i].LUN.String
		finalMyShops.MAR = &MyShopUnStructured[i].MAR.String
		finalMyShops.MIE = &MyShopUnStructured[i].MIE.String
		finalMyShops.JUE = &MyShopUnStructured[i].JUE.String
		finalMyShops.VIE = &MyShopUnStructured[i].VIE.String
		finalMyShops.SAB = &MyShopUnStructured[i].SAB.String
		finalMyShops.DOM = &MyShopUnStructured[i].DOM.String
		finalMyShops.UserID = &MyShopUnStructured[i].UserID.String
		finalMyShops.Images = strings.Split(MyShopUnStructured[i].Images.String, ",")
		finalMyShops.DateInit = &MyShopUnStructured[i].DateInit.String
		finalMyShops.DateFinish = &MyShopUnStructured[i].DateFinish.String
		finalMyShops.TypeCharge = &MyShopUnStructured[i].TypeCharge.String

		listReponse = append(listReponse, finalMyShops)
	}

	c.JSON(ResponseResult{Result: listReponse})
}

//ShopOffers handler for get offers
func ShopOffers(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	ShopOffers := new(ParamsShopOffers)

	if err := c.QueryParser(ShopOffers); err != nil {
		fmt.Println(err, "Error parsing shop id")
	}

	page, _ := strconv.Atoi(ShopOffers.Page)
	limit, _ := strconv.Atoi(ShopOffers.Limit)
	offset := limit * page

	var ListOfferSQL ResponseListOffersSQL
	var ListOffers []ResponseListOffersSQL
	var OffersPointer ResponseListOffers
	Offers := []ResponseListOffers{}

	ChainOfferSQL := sq.Select(
		"offers_id",
		"title",
		"description",
		"date_init",
		"date_end",
		"image_url",
		"active",
		"lat",
		"lon",
	).
		From("offers")

	ChainWhere := "shop_id = ?"

	if ShopOffers.Status == "actives" {
		ChainWhere = ChainWhere + " and active = 1"
	}

	if ShopOffers.Status == "inactive" {
		ChainWhere = ChainWhere + " and active = 0"
	}

	OfferSQL, err := ChainOfferSQL.
		Where(ChainWhere, ShopID).
		OrderBy("create_at_offer DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		RunWith(database).
		Query()

	if err != nil {
		fmt.Println(err, "Error to get offers")
		ErrorProfile := ErrorResponse{MESSAGE: "Error to get offers"}
		c.JSON(ErrorProfile)
		c.Status(400)
		return
	}

	for OfferSQL.Next() {
		_ = OfferSQL.Scan(
			&ListOfferSQL.OffersID,
			&ListOfferSQL.Title,
			&ListOfferSQL.Description,
			&ListOfferSQL.DateInit,
			&ListOfferSQL.DateEnd,
			&ListOfferSQL.ImageURL,
			&ListOfferSQL.Active,
			&ListOfferSQL.Lat,
			&ListOfferSQL.Lon,
		)

		ListOffers = append(ListOffers, ListOfferSQL)
	}

	for i := 0; i < len(ListOffers); i++ {
		OffersPointer.OffersID = &ListOffers[i].OffersID.String
		OffersPointer.Title = &ListOffers[i].Title.String
		OffersPointer.Description = &ListOffers[i].Description.String
		OffersPointer.DateInit = &ListOffers[i].DateInit.String
		OffersPointer.DateEnd = &ListOffers[i].DateEnd.String
		OffersPointer.ImageURL = &ListOffers[i].ImageURL.String
		OffersPointer.Active = &ListOffers[i].Active.String
		OffersPointer.Lat = &ListOffers[i].Lat.String
		OffersPointer.Lon = &ListOffers[i].Lon.String

		Offers = append(Offers, OffersPointer)
	}

	response := ResponseResultOffers{Offers: Offers}
	c.JSON(response)
}

//OfferInfo handler for get information of a offer
func OfferInfo(c *fiber.Ctx) {
	OfferID := c.Params("offer_id")
	fmt.Println(OfferID)

	var AOfferSQL AOffer
	var ResponseToOffer AOfferPointer
	ToResponse := ResponseInforOffer{}

	OfferResultsError := sq.Select(
		"offers_id",
		"title",
		"offers.description",
		"date_end",
		"image_url",
		"active",
		"offers.lat",
		"offers.lon",
		"shop.shop_id",
		"shop_name",
		"cover_image",
	).
		From("offers").
		LeftJoin("shop on offers.shop_id = shop.shop_id").
		Where("offers_id = ?", OfferID).
		RunWith(database).
		QueryRow().
		Scan(
			&AOfferSQL.OffersID,
			&AOfferSQL.Title,
			&AOfferSQL.Description,
			&AOfferSQL.DateEnd,
			&AOfferSQL.ImageURL,
			&AOfferSQL.Active,
			&AOfferSQL.Lat,
			&AOfferSQL.Lon,
			&AOfferSQL.ShopID,
			&AOfferSQL.ShopName,
			&AOfferSQL.CoverImage,
		)

	if OfferResultsError != nil {
		fmt.Println(OfferResultsError, "Error get offers")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get offers"})
		c.SendStatus(400)
		return
	}

	ResponseToOffer.OffersID = &AOfferSQL.OffersID.String
	ResponseToOffer.Title = &AOfferSQL.Title.String
	ResponseToOffer.Description = &AOfferSQL.Description.String
	ResponseToOffer.DateEnd = &AOfferSQL.DateEnd.String
	ResponseToOffer.ImageURL = &AOfferSQL.ImageURL.String
	ResponseToOffer.Active = &AOfferSQL.Active.String
	ResponseToOffer.Lat = &AOfferSQL.Lat.String
	ResponseToOffer.Lon = &AOfferSQL.Lon.String
	ResponseToOffer.ShopID = &AOfferSQL.ShopID.String
	ResponseToOffer.ShopName = &AOfferSQL.ShopName.String
	ResponseToOffer.CoverImage = &AOfferSQL.CoverImage.String

	ToResponse.Offer = ResponseToOffer

	c.JSON(ToResponse)
}

//Services handler for get services
func Services(c *fiber.Ctx) {
	var ServiceFromSQL ServiceSQL
	var ServicePointer []ServiceSQL
	var Pointer ServiceSQLPointer
	Pointers := []ServiceSQLPointer{}

	Services, ErrorGetService := sq.Select(
		"service_type_id",
		"service_name",
	).
		From("service_type").
		RunWith(database).
		Query()

	if ErrorGetService != nil {
		fmt.Println(ErrorGetService, "Error get services")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get services"})
		c.SendStatus(400)
		return
	}

	for Services.Next() {
		_ = Services.Scan(
			&ServiceFromSQL.ServiceTypeID,
			&ServiceFromSQL.ServiceName,
		)

		ServicePointer = append(ServicePointer, ServiceFromSQL)
	}

	for i := 0; i < len(ServicePointer); i++ {
		Pointer.ServiceTypeID = &ServicePointer[i].ServiceTypeID.String
		Pointer.ServiceName = &ServicePointer[i].ServiceName.String

		Pointers = append(Pointers, Pointer)
	}

	response := ServiceResponse{Services: Pointers}

	c.JSON(response)
}

//SubServices Handler for get sub services of a service
func SubServices(c *fiber.Ctx) {
	ServiceID := c.Params("service_id")

	var SubServiceFromSQL SubServiceSQL
	var SubServicePointer []SubServiceSQL
	var Pointer SubServiceSQLPointer
	Pointers := []SubServiceSQLPointer{}

	SubServicesSQL, ErrorSubservices := sq.Select(
		"sub_service_type_id",
		"sub_service_name",
	).
		From("sub_service_type").
		Where("service_type_id = ?", ServiceID).
		RunWith(database).
		Query()

	if ErrorSubservices != nil {
		fmt.Println(ErrorSubservices, "Error get sub services")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get sub services"})
		c.SendStatus(400)
		return
	}

	for SubServicesSQL.Next() {
		_ = SubServicesSQL.Scan(
			&SubServiceFromSQL.SubServiceTypeID,
			&SubServiceFromSQL.SubServiceName,
		)

		SubServicePointer = append(SubServicePointer, SubServiceFromSQL)
	}

	for i := 0; i < len(SubServicePointer); i++ {
		Pointer.SubServiceTypeID = &SubServicePointer[i].SubServiceTypeID.String
		Pointer.SubServiceName = &SubServicePointer[i].SubServiceName.String

		Pointers = append(Pointers, Pointer)
	}

	response := SubServiceResponse{SubService: Pointers}

	c.JSON(response)
}

//Comments Handler for get comments of a shop
func Comments(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")

	var CommentFromSQL CommentsSQL
	CommentsFromSQL := []CommentsSQL{}
	var CommentPointer CommentsPointerStruct
	CommentsPointer := []CommentsPointerStruct{}

	Pagination := new(Paginations)

	if err := c.QueryParser(Pagination); err != nil {
		fmt.Println(err, "Error parsing shops")
	}

	Page, _ := strconv.Atoi(Pagination.Page)
	Limit, _ := strconv.Atoi(Pagination.Limit)

	Offset := Limit * Page

	FromSQL, ErrorComments := sq.Select(
		"comment",
		"create_date_at",
		"usersk.user_id",
		"fullname",
		"image",
	).
		From("shop_comments").
		LeftJoin("usersk on shop_comments.user_id = usersk.user_id").
		Where("shop_id = ?", ShopID).
		OrderBy("create_date_at DESC").
		Limit(uint64(Limit)).
		Offset(uint64(Offset)).
		RunWith(database).
		Query()

	if ErrorComments != nil {
		fmt.Println(ErrorComments, "Error to get comments")
		ErrorProfile := ErrorResponse{MESSAGE: "Error to get comments"}
		c.JSON(ErrorProfile)
		c.Status(400)
		return
	}

	for FromSQL.Next() {
		_ = FromSQL.Scan(
			&CommentFromSQL.Comment,
			&CommentFromSQL.CreateDateAt,
			&CommentFromSQL.UserID,
			&CommentFromSQL.Fullname,
			&CommentFromSQL.Image,
		)

		CommentsFromSQL = append(CommentsFromSQL, CommentFromSQL)
	}

	for i := 0; i < len(CommentsFromSQL); i++ {
		CommentPointer.Comment = &CommentsFromSQL[i].Comment.String
		CommentPointer.CreateDateAt = &CommentsFromSQL[i].CreateDateAt.String
		CommentPointer.UserID = &CommentsFromSQL[i].UserID.String
		CommentPointer.Fullname = &CommentsFromSQL[i].Fullname.String
		CommentPointer.Image = &CommentsFromSQL[i].Image.String

		CommentsPointer = append(CommentsPointer, CommentPointer)
	}

	c.JSON(CommentsPointer)
}

//Score Handler for get score
func Score(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := c.Params("user_id")
	var Score ScoreSQL
	var Response ResponseScore

	ErrorScore := sq.Select(
		"AVG(score) as score",
	).
		From("shop_score_users").
		Where("shop_id = ? AND user_id = ?", ShopID, UserID).
		RunWith(database).
		QueryRow().
		Scan(
			&Score.Score,
		)

	if ErrorScore != nil {
		fmt.Println(ErrorScore, "Error get score")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get score"})
		c.SendStatus(400)
		return
	}

	Response.Score = &Score.Score.String

	c.JSON(Response)
}

//Page Handler for get page
func Page(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	userID := userIDF(c.Get("token"))
	var PageFromSQL PageSQL
	var PagePointer PagesPointer
	var response PageInformation

	ErrorPage := sq.Select(
		"active",
		"template_type",
		"style_sheets",
		"active_days",
		"images_days",
		"offers_active",
		"accept_card_active",
		"subdomain",
		"domain",
		"shop.shop_id",
		"pages_id",
		"type_charge",
		"shop_name",
		"description",
		"cover_image",
		"logo",
		"shop.user_id",
	).
		From("pages").
		LeftJoin("plans_pay on pages.shop_id = plans_pay.shop_id").
		LeftJoin("shop on pages.shop_id = shop.shop_id").
		Where("pages.shop_id = ? AND shop.user_id = ?", ShopID, userID).
		RunWith(database).
		QueryRow().
		Scan(
			&PageFromSQL.Active,
			&PageFromSQL.TemplateType,
			&PageFromSQL.StyleSheets,
			&PageFromSQL.ActiveDays,
			&PageFromSQL.ImagesDays,
			&PageFromSQL.OffersActive,
			&PageFromSQL.AcceptCardActive,
			&PageFromSQL.Subdomain,
			&PageFromSQL.Domain,
			&PageFromSQL.ShopID,
			&PageFromSQL.PagesID,
			&PageFromSQL.TypeCharge,
			&PageFromSQL.ShopName,
			&PageFromSQL.Description,
			&PageFromSQL.CoverImage,
			&PageFromSQL.Logo,
			&PageFromSQL.UserID,
		)

	if ErrorPage != nil {
		fmt.Println(ErrorPage, "Error get page")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get page"})
		c.SendStatus(400)
		return
	}

	Active, _ := strconv.Atoi(PageFromSQL.Active.String)
	PagePointer.Active = Active
	TemplateType, _ := strconv.Atoi(PageFromSQL.TemplateType.String)
	PagePointer.TemplateType = TemplateType
	StyleSheets, _ := strconv.Atoi(PageFromSQL.StyleSheets.String)
	PagePointer.StyleSheets = StyleSheets
	ActiveDays, _ := strconv.Atoi(PageFromSQL.ActiveDays.String)
	PagePointer.ActiveDays = ActiveDays
	ImagesDays, _ := strconv.Atoi(PageFromSQL.ImagesDays.String)
	PagePointer.ImagesDays = ImagesDays
	OffersActive, _ := strconv.Atoi(PageFromSQL.OffersActive.String)
	PagePointer.OffersActive = OffersActive
	AcceptCardActive, _ := strconv.Atoi(PageFromSQL.AcceptCardActive.String)
	PagePointer.AcceptCardActive = AcceptCardActive
	PagePointer.Subdomain = &PageFromSQL.Subdomain.String
	PagePointer.Domain = &PageFromSQL.Domain.String
	PagePointer.ShopID = &PageFromSQL.ShopID.String
	PagePointer.PagesID = &PageFromSQL.PagesID.String
	PagePointer.TypeCharge = &PageFromSQL.TypeCharge.String
	PagePointer.ShopName = &PageFromSQL.ShopName.String
	PagePointer.Description = &PageFromSQL.Description.String
	PagePointer.CoverImage = &PageFromSQL.CoverImage.String
	PagePointer.Logo = &PageFromSQL.Logo.String
	PagePointer.UserID = &PageFromSQL.UserID.String

	response.Page = PagePointer
	UserID := *PagePointer.UserID
	if UserID == userID {
		c.JSON(response)
	} else {
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
	}
}

//FindShops Handler for get Shop
func FindShops(c *fiber.Ctx) {
	Lat := c.Params("lat")
	Lon := c.Params("lon")

	var shop FShops
	var array []FShops
	var ShopIDStruct ShopsID
	var arrShopID []ShopsID
	var ShopFind ShopFromSQLFind
	var ShopArr []ShopFromSQLFind
	var ShopPointer ShopFindPointer
	var ShopArrPointer []ShopFindPointer

	LatI, _ := strconv.ParseFloat(Lat, 64)
	LonI, _ := strconv.ParseFloat(Lon, 64)

	Query := new(FindShopQuery)

	if err := c.QueryParser(Query); err != nil {
		fmt.Println(err, "Error parsing Query")
	}

	Query.LastShopID = c.Query("last_shop_id")

	fmt.Println(Query.LastShopID, "last")

	Limit, _ := strconv.Atoi(Query.Limit)
	MinDistance, _ := strconv.ParseFloat(Query.MinDistance, 64)
	Aggregate := []bson.M{
		bson.M{
			"$geoNear": bson.M{
				"near": bson.M{
					"type":        "Point",
					"coordinates": []float64{LonI, LatI},
				},
				"minDistance":   MinDistance,
				"spherical":     true,
				"distanceField": "distance",
			},
		},
		bson.M{
			"$match": bson.M{
				"status":   true,
				"category": Query.Category,
				"shop_id": bson.M{
					"$nin": []string{Query.LastShopID},
				},
			},
		},
		bson.M{
			"$limit": Limit,
		},
	}

	curl, error := mongodb.Collection("shop").Aggregate(context.TODO(), Aggregate, options.Aggregate())

	if error != nil {
		fmt.Println(error)
	}

	for curl.Next(context.TODO()) {
		_ = curl.Decode(&shop)
		ShopIDStruct.ShopID = shop.ShopID
		ShopIDStruct.Distance = shop.Distance
		arrShopID = append(arrShopID, ShopIDStruct)

		array = append(array, shop)
	}
	// fmt.Println(array)
	IDs := ""

	for i := 0; i < len(arrShopID); i++ {
		if len(arrShopID)-1 == i {
			IDs = fmt.Sprintf("%s%v", IDs, arrShopID[i].ShopID)
		} else {
			IDs = fmt.Sprintf("%s%v,", IDs, arrShopID[i].ShopID)
		}
	}

	where := fmt.Sprintf("shop_id IN (%s)", IDs)

	fmt.Println(IDs)

	ShopsFromSQL, ErrorShops := sq.Select(
		"shop_id",
		"shop_name",
		"address",
		"phone",
		"score_shop",
		"cover_image",
		"service_name",
		"sub_service_name",
	).
		From("shop").
		LeftJoin("service_type on shop.service_type_id = service_type.service_type_id").
		LeftJoin("sub_service_type on shop.sub_service_type_id = sub_service_type.sub_service_type_id").
		Where(where).
		RunWith(database).
		Query()

	if ErrorShops != nil {
		fmt.Println(ErrorShops, "Error get Shops find")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get Shops find"})
		c.SendStatus(400)
		return
	}

	for ShopsFromSQL.Next() {
		_ = ShopsFromSQL.Scan(
			&ShopFind.ShopID,
			&ShopFind.ShopName,
			&ShopFind.Address,
			&ShopFind.Phone,
			&ShopFind.ScoreShop,
			&ShopFind.CoverImage,
			&ShopFind.ServiceName,
			&ShopFind.SubServiceName,
		)

		ShopArr = append(ShopArr, ShopFind)
	}

	for i := 0; i < len(ShopArr); i++ {
		ShopPointer.ShopID = &ShopArr[i].ShopID.String
		ShopPointer.ShopName = &ShopArr[i].ShopName.String
		ShopPointer.Address = &ShopArr[i].Address.String
		ShopPointer.Phone = &ShopArr[i].Phone.String
		ShopPointer.ScoreShop = &ShopArr[i].ScoreShop.String
		ShopPointer.CoverImage = &ShopArr[i].CoverImage.String
		ShopPointer.ServiceName = &ShopArr[i].ServiceName.String
		ShopPointer.SubServiceName = &ShopArr[i].SubServiceName.String

		for e := 0; e < len(arrShopID); e++ {
			if arrShopID[e].ShopID == ShopArr[i].ShopID.String {
				ShopPointer.Distance = arrShopID[e].Distance
			}
		}

		ShopArrPointer = append(ShopArrPointer, ShopPointer)
	}

	sort.Slice(ShopArrPointer, func(i, j int) bool {
		return ShopArrPointer[i].Distance < ShopArrPointer[j].Distance
	})

	c.JSON(ResponseFinalFindShops{
		Shop:         ShopArrPointer,
		LastShopID:   ShopArrPointer[len(ShopArrPointer)-1].ShopID,
		LastDistance: ShopArrPointer[len(ShopArrPointer)-1].Distance,
	})
}

//Comment Handle for insert comments
func Comment(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := userIDF(c.Get("token"))

	var Data CommentStruct

	if errorParse := c.BodyParser(&Data); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	id, errorInsert := sq.Insert("shop_comments").
		Columns(
			"user_id",
			"shop_id",
			"comment",
		).
		Values(
			UserID,
			ShopID,
			Data.Comment,
		).
		RunWith(database).
		Exec()

	if errorInsert != nil {
		fmt.Println("Error to save shop", errorInsert)
	}

	IDLast, _ := id.LastInsertId()
	IDS := strconv.FormatInt(IDLast, 10)

	c.JSON(SuccessResponse{MESSAGE: IDS})
}

//FindOffers Handler for get Offers
func FindOffers(c *fiber.Ctx) {
	Lat := c.Params("lat")
	Lon := c.Params("lon")

	var offers FOffers
	var array []FOffers
	var OffersIDStruct OffersID
	var arrOffersID []OffersID
	var OffersFind OffersFromSQLFind
	var OffersArr []OffersFromSQLFind
	var OffersPointer OffersFindPointer
	var OffersArrPointer []OffersFindPointer

	LatI, _ := strconv.ParseFloat(Lat, 64)
	LonI, _ := strconv.ParseFloat(Lon, 64)

	Query := new(FindOfferQuery)

	if err := c.QueryParser(Query); err != nil {
		fmt.Println(err, "Error parsing Query")
	}

	Query.LastOfferID = c.Query("last_offer_id")

	fmt.Println(Query.LastOfferID, "last")

	Limit, _ := strconv.Atoi(Query.Limit)
	MinDistance, _ := strconv.ParseFloat(Query.MinDistance, 64)

	_, errTime := time.LoadLocation("America/Mexico_City")

	if errTime != nil {
		fmt.Println(errTime)
	}

	TimeUnParsing := time.Now()
	TimeString := fmt.Sprintf("%s", TimeUnParsing)
	TimePaser := strings.Split(TimeString, " ")
	TimeRemoveSpace := strings.TrimRight(TimePaser[0], " ")

	Aggregate := []bson.M{
		bson.M{
			"$geoNear": bson.M{
				"near": bson.M{
					"type":        "Point",
					"coordinates": []float64{LonI, LatI},
				},
				"minDistance":   MinDistance,
				"spherical":     true,
				"distanceField": "distance",
			},
		},
		bson.M{
			"$match": bson.M{
				"active": true,
				"offer_id": bson.M{
					"$nin": []string{Query.LastOfferID},
				},
				"date_end": bson.M{
					"$gt": TimeRemoveSpace,
				},
			},
		},
		bson.M{
			"$limit": Limit,
		},
	}

	curl, error := mongodb.Collection("offers").Aggregate(context.TODO(), Aggregate, options.Aggregate())

	if error != nil {
		fmt.Println(error)
	}

	for curl.Next(context.TODO()) {
		_ = curl.Decode(&offers)
		OffersIDStruct.OfferID = offers.OfferID
		OffersIDStruct.Distance = offers.Distance
		arrOffersID = append(arrOffersID, OffersIDStruct)

		array = append(array, offers)
	}
	// fmt.Println(array)
	IDs := ""

	for i := 0; i < len(arrOffersID); i++ {
		if len(arrOffersID)-1 == i {
			IDs = fmt.Sprintf("%s%v", IDs, arrOffersID[i].OfferID)
		} else {
			IDs = fmt.Sprintf("%s%v,", IDs, arrOffersID[i].OfferID)
		}
	}

	where := fmt.Sprintf("offers_id IN (%s)", IDs)

	offersSQL := sq.Select(
		"offers_id",
		"title",
		"offers.description",
		"date_init",
		"date_end",
		"image_url",
		"shop.shop_id",
		"shop_name",
		"cover_image",
	).
		From("offers").
		LeftJoin("shop on shop.shop_id = offers.shop_id")

	if len(IDs) > 0 {
		offersSQL = offersSQL.Where(where)
	}

	OffersFromSQL, ErrorOffers := offersSQL.
		RunWith(database).
		Query()

	if ErrorOffers != nil {
		fmt.Println(ErrorOffers, "Error get Offers find")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get Offers find"})
		c.SendStatus(400)
		return
	}

	for OffersFromSQL.Next() {
		_ = OffersFromSQL.Scan(
			&OffersFind.OfferID,
			&OffersFind.Title,
			&OffersFind.Description,
			&OffersFind.DateInit,
			&OffersFind.DateEnd,
			&OffersFind.ImageURL,
			&OffersFind.ShopID,
			&OffersFind.ShopName,
			&OffersFind.CoverImage,
		)

		OffersArr = append(OffersArr, OffersFind)
	}

	for i := 0; i < len(OffersArr); i++ {
		OffersPointer.OfferID = &OffersArr[i].OfferID.String
		OffersPointer.Title = &OffersArr[i].Title.String
		OffersPointer.Description = &OffersArr[i].Description.String
		OffersPointer.DateInit = &OffersArr[i].DateInit.String
		OffersPointer.DateEnd = &OffersArr[i].DateEnd.String
		OffersPointer.ImageURL = &OffersArr[i].ImageURL.String
		OffersPointer.ShopID = &OffersArr[i].ShopID.String
		OffersPointer.ShopName = &OffersArr[i].ShopName.String
		OffersPointer.CoverImage = &OffersArr[i].CoverImage.String

		for e := 0; e < len(arrOffersID); e++ {
			if arrOffersID[e].OfferID == OffersArr[i].OfferID.String {
				OffersPointer.Distance = arrOffersID[e].Distance
			}
		}

		OffersArrPointer = append(OffersArrPointer, OffersPointer)
	}

	sort.Slice(OffersArrPointer, func(i, j int) bool {
		return OffersArrPointer[i].Distance < OffersArrPointer[j].Distance
	})

	c.JSON(ResponseFinalFindOffers{
		Offers:       OffersArrPointer,
		LastOfferID:  OffersArrPointer[len(OffersArrPointer)-1].ShopID,
		LastDistance: OffersArrPointer[len(OffersArrPointer)-1].Distance,
	})
}

//CreateOffer Handler for create offer
func CreateOffer(c *fiber.Ctx) {
	UserID := userIDF(c.Get("token"))

	var Query QueryParamsOffer

	if errorParse := c.BodyParser(&Query); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	var IsOwner IsOwnerShop
	var Position LocationSQL

	fmt.Println(UserID, Query.ShopID, Query.Title, "?")

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = true",
			UserID,
			Query.ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	ErrorShop := sq.Select(
		"lat",
		"lon",
	).
		From("shop").
		Where("shop_id = ? AND user_id = ?", Query.ShopID, UserID).
		RunWith(database).
		QueryRow().
		Scan(
			&Position.Lat,
			&Position.Lon,
		)

	if ErrorShop != nil {
		fmt.Println("Not found shop")
		c.JSON(ErrorResponse{MESSAGE: "Not found shop"})
		c.SendStatus(400)
		return
	}

	id, errorInsert := sq.Insert("offers").
		Columns(
			"user_id",
			"shop_id",
			"title",
			"description",
			"date_init",
			"date_end",
			"image_url",
			"lat",
			"lon",
			"active",
		).
		Values(
			UserID,
			Query.ShopID,
			Query.Title,
			Query.Description,
			Query.DateInit,
			Query.DateEnd,
			Query.ImageURL,
			&Position.Lat,
			&Position.Lon,
			0,
		).
		RunWith(database).
		Exec()

	IDLast, _ := id.LastInsertId()
	fmt.Println("This offer id", IDLast)

	if errorInsert != nil {
		fmt.Println("Error to save shop", errorInsert)
	}

	PosLat := Position.Lat.String
	PosLon := Position.Lon.String

	Lat, errLat := strconv.ParseFloat(PosLat, 64)
	if errLat != nil {
		fmt.Println("Error to conver lat", errLat)
	}

	Lon, errLon := strconv.ParseFloat(PosLon, 64)
	if errLon != nil {
		fmt.Println("Error to conver lon", errLon)
	}

	// IDString := fmt.Sprintf("%s", IDLast)
	IDString := strconv.FormatInt(IDLast, 10)

	resInsertMongo, errInsertMongo := mongodb.Collection("offers").InsertOne(context.TODO(), bson.M{
		"offer_id": IDString,
		"shop_id":  Query.ShopID,
		"title":    Query.Title,
		"location": bson.M{
			"type":        "Point",
			"coordinates": []float64{Lon, Lat},
		},
		"date_init": Query.DateInit,
		"date_end":  Query.DateEnd,
		"active":    false,
	})

	if errInsertMongo != nil {
		fmt.Println(errInsertMongo, "Error to Insert mongo")
	}

	IDMongo := resInsertMongo.InsertedID

	fmt.Println(IDMongo, "Id of offer in mongodb")

	c.JSON(SuccessResponseOffer{MESSAGE: "Created offers", OfferID: IDString, Status: 200})

}

//UpdateOffer Handler for update offer
// https://github.com/mongodb/mongo-go-driver/blob/master/mongo/crud_examples_test.go
func UpdateOffer(c *fiber.Ctx) {
	UserID := userIDF(c.Get("token"))
	OfferID := c.Params("offer_id")

	var Query QueryParamsOfferUpdate
	var IsOwner IsOwnerShop
	var OffersMongo bson.D

	if errorParse := c.BodyParser(&Query); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = true",
			UserID,
			Query.ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	queryUpdateValue := sq.Update("offers")

	if len(Query.Title) > 0 {
		queryUpdateValue = queryUpdateValue.Set("title", Query.Title)
		OffersMongo = append(OffersMongo, bson.E{"title", Query.Title})
	}

	if len(Query.Description) > 0 {
		queryUpdateValue = queryUpdateValue.Set("description", Query.Description)
		OffersMongo = append(OffersMongo, bson.E{"description", Query.Description})
	}

	if len(Query.DateInit) > 0 {
		queryUpdateValue = queryUpdateValue.Set("date_init", Query.DateInit)
		OffersMongo = append(OffersMongo, bson.E{"date_init", Query.DateInit})
	}

	if len(Query.DateEnd) > 0 {
		queryUpdateValue = queryUpdateValue.Set("date_end", Query.DateEnd)
		OffersMongo = append(OffersMongo, bson.E{"date_end", Query.DateEnd})
	}

	if len(Query.ImageURL) > 0 {
		queryUpdateValue = queryUpdateValue.Set("image_url", Query.ImageURL)
		OffersMongo = append(OffersMongo, bson.E{"image_url", Query.ImageURL})
	}

	if Query.Active >= 0 && Query.Active <= 1 {
		queryUpdateValue = queryUpdateValue.Set("active", Query.Active)

		Active := false
		if Query.Active == 0 {
			Active = false
		} else {
			Active = true
		}

		OffersMongo = append(OffersMongo, bson.E{"active", Active})
	}

	_, ErrorUpdateOffer := queryUpdateValue.
		Where("offers_id = ? AND user_id = ? AND shop_id = ?", OfferID, UserID, Query.ShopID).
		RunWith(database).
		Exec()

	if ErrorUpdateOffer != nil {
		fmt.Println(ErrorUpdateOffer, "Problem with update offer")
		c.JSON(ErrorResponse{MESSAGE: "Problem with update offer"})
		c.SendStatus(500)
		return
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"offer_id", OfferID}}
	update := bson.D{
		{"$set", OffersMongo},
	}

	resultMongoOffers, errOfferMongo := mongodb.Collection("offers").UpdateOne(context.TODO(), filter, update, opts)

	if errOfferMongo != nil {
		fmt.Println("promblem with update offer in mongodb")
	}

	if resultMongoOffers.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	if resultMongoOffers.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", resultMongoOffers.UpsertedID)
	}

	fmt.Println(OffersMongo, queryUpdateValue)

	c.JSON(SuccessResponseOfferStatus{MESSAGE: "Success update", Status: 200})
}

//SetScore set my score to shop
func SetScore(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := userIDF(c.Get("token"))

	var Data ScoreStruct

	if errorParse := c.BodyParser(&Data); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	if Data.Score <= 0 {
		fmt.Println("score can't be empty")
		c.JSON(ErrorResponse{MESSAGE: "Score can't be empty"})
		c.Status(400)
		return
	}

	var Score ScoreSQL
	var Response ResponseScore

	_, errorInsert := sq.Insert("shop_score_users").
		Columns(
			"user_id",
			"shop_id",
			"score",
		).
		Values(
			UserID,
			ShopID,
			Data.Score,
		).
		RunWith(database).
		Exec()

	if errorInsert != nil {
		fmt.Println("Error to save score", errorInsert)
	}

	ErrorScore := sq.Select(
		"AVG(score) as score",
	).
		From("shop_score_users").
		Where("shop_id = ?", ShopID).
		RunWith(database).
		QueryRow().
		Scan(
			&Score.Score,
		)

	if ErrorScore != nil {
		fmt.Println(ErrorScore, "Error get score")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get score"})
		c.SendStatus(400)
		return
	}

	queryUpdateValue := sq.Update("shop")

	ScoreFloat, _ := strconv.ParseFloat(Score.Score.String, 64)

	NewScore := fmt.Sprintf("%.0f", ScoreFloat)

	_, ErrorUpdateOffer := queryUpdateValue.
		Set("score_shop", NewScore).
		Where("shop_id = ? ", ShopID).
		RunWith(database).
		Exec()

	if ErrorUpdateOffer != nil {
		fmt.Println(ErrorUpdateOffer, "Problem with update score")
		c.JSON(ErrorResponse{MESSAGE: "Problem with update score"})
		c.SendStatus(500)
		return
	}

	Response.Score = &NewScore

	c.JSON(Response)
}

//UpdateScore Update my score to shop
func UpdateScore(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := userIDF(c.Get("token"))

	var Data ScoreStruct

	if errorParse := c.BodyParser(&Data); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	if Data.Score <= 0 {
		fmt.Println("score can't be empty")
		c.JSON(ErrorResponse{MESSAGE: "Score can't be empty"})
		c.Status(400)
		return
	}

	var Score ScoreSQL

	_, ErrorUpdateScoreUser := sq.Update("shop_score_users").
		Set("score", Data.Score).
		Where("user_id = ? AND shop_id = ?", UserID, ShopID).
		RunWith(database).
		Exec()

	if ErrorUpdateScoreUser != nil {
		fmt.Println("Error to update score", ErrorUpdateScoreUser)
	}

	ErrorScore := sq.Select(
		"AVG(score) as score",
	).
		From("shop_score_users").
		Where("shop_id = ?", ShopID).
		RunWith(database).
		QueryRow().
		Scan(
			&Score.Score,
		)

	if ErrorScore != nil {
		fmt.Println(ErrorScore, "Error get score")
		c.JSON(ErrorResponse{MESSAGE: "Problem with get score"})
		c.SendStatus(400)
		return
	}

	queryUpdateValue := sq.Update("shop")

	ScoreFloat, _ := strconv.ParseFloat(Score.Score.String, 64)

	NewScore := fmt.Sprintf("%.0f", ScoreFloat)

	_, ErrorUpdateScoreShop := queryUpdateValue.
		Set("score_shop", NewScore).
		Where("shop_id = ? ", ShopID).
		RunWith(database).
		Exec()

	if ErrorUpdateScoreShop != nil {
		fmt.Println(ErrorUpdateScoreShop, "Problem with update score")
		c.JSON(ErrorResponse{MESSAGE: "Problem with update score"})
		c.SendStatus(500)
		return
	}

	c.JSON(SuccessResponse{MESSAGE: "Calificación actualizada"})
}

//LockShop Handler for lock shop
func LockShop(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := userIDF(c.Get("token"))

	var IsOwner IsOwnerShop
	var ShopMongo bson.D
	var OffersMongo bson.D

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = 1",
			UserID,
			ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorLock := sq.Update("shop").
		Set("status", 0).
		Set("lock_shop", 1).
		Where("user_id = ? AND shop_id = ?", UserID, ShopID).
		RunWith(database).
		Exec()

	if ErrorLock != nil {
		fmt.Println("Problem with lock shop", ErrorLock)
		c.JSON(ErrorResponse{MESSAGE: "Problem with lock shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorLockOffer := sq.Update("offers").
		Set("active", 0).
		Where("user_id = ? AND shop_id = ?", UserID, ShopID).
		RunWith(database).
		Exec()

	if ErrorLockOffer != nil {
		fmt.Println("Problem with lock offer", ErrorLockOffer)
		c.JSON(ErrorResponse{MESSAGE: "Problem with lock offer"})
		c.SendStatus(400)
		return
	}

	_, ErrorLockPage := sq.Update("pages").
		Set("active", 0).
		Where("shop_id = ?", ShopID).
		RunWith(database).
		Exec()

	if ErrorLockPage != nil {
		fmt.Println("Problem with lock page", ErrorLockPage)
		c.JSON(ErrorResponse{MESSAGE: "Problem with lock page"})
		c.SendStatus(400)
		return
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"shop_id", ShopID}}
	ShopMongo = append(ShopMongo, bson.E{"status", false})

	update := bson.D{
		{"$set", ShopMongo},
	}

	ShopMongoLock, ErrorLockMongo := mongodb.Collection("shop").UpdateOne(context.TODO(), filter, update, opts)

	if ErrorLockMongo != nil {
		fmt.Println("promblem with lock shop in mongodb")
	}

	if ShopMongoLock.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	}
	if ShopMongoLock.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", ShopMongoLock.UpsertedID)
	}

	filterOffers := bson.D{{"shop_id", ShopID}}
	OffersMongo = append(OffersMongo, bson.E{"active", false})
	updateOffers := bson.D{
		{"$set", OffersMongo},
	}

	MongoLockOffers, MongoOffersLockError := mongodb.Collection("offers").UpdateMany(context.TODO(), filterOffers, updateOffers)
	if MongoOffersLockError != nil {
		log.Println("Problem with update lock offers", MongoOffersLockError)
	}

	if MongoLockOffers.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	}

	c.JSON(SuccessResponse{MESSAGE: "Tienda desabilitada"})
}

//UnlockShop Handler for lock shop
func UnlockShop(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	UserID := userIDF(c.Get("token"))

	var IsOwner IsOwnerShop
	var ShopMongo bson.D
	var OffersMongo bson.D

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = 0",
			UserID,
			ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorUnlok := sq.Update("shop").
		Set("status", 1).
		Set("lock_shop", 0).
		Where("user_id = ? AND shop_id = ?", UserID, ShopID).
		RunWith(database).
		Exec()

	if ErrorUnlok != nil {
		fmt.Println("Problem with Unlok shop", ErrorUnlok)
		c.JSON(ErrorResponse{MESSAGE: "Problem with Unlok shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorUnlokOffer := sq.Update("offers").
		Set("active", 1).
		Where("user_id = ? AND shop_id = ?", UserID, ShopID).
		RunWith(database).
		Exec()

	if ErrorUnlokOffer != nil {
		fmt.Println("Problem with Unlok offer", ErrorUnlokOffer)
		c.JSON(ErrorResponse{MESSAGE: "Problem with Unlok offer"})
		c.SendStatus(400)
		return
	}

	_, ErrorUnlokPage := sq.Update("pages").
		Set("active", 1).
		Where("shop_id = ?", ShopID).
		RunWith(database).
		Exec()

	if ErrorUnlokPage != nil {
		fmt.Println("Problem with Unlok page", ErrorUnlokPage)
		c.JSON(ErrorResponse{MESSAGE: "Problem with Unlok page"})
		c.SendStatus(400)
		return
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"shop_id", ShopID}}
	ShopMongo = append(ShopMongo, bson.E{"status", true})
	update := bson.D{
		{"$set", ShopMongo},
	}

	ShopMongoUnlok, ErrorUnlokMongo := mongodb.Collection("shop").UpdateOne(context.TODO(), filter, update, opts)

	if ErrorUnlokMongo != nil {
		fmt.Println("promblem with Unlok shop in mongodb")
	}

	if ShopMongoUnlok.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	}
	if ShopMongoUnlok.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", ShopMongoUnlok.UpsertedID)
	}

	filterOffers := bson.D{{"shop_id", ShopID}}
	OffersMongo = append(OffersMongo, bson.E{"active", true})
	updateOffers := bson.D{
		{"$set", OffersMongo},
	}

	MongoUnlokOffers, MongoOffersUnlokError := mongodb.Collection("offers").UpdateMany(context.TODO(), filterOffers, updateOffers)
	if MongoOffersUnlokError != nil {
		log.Println("Problem with update Unlok offers", MongoOffersUnlokError)
	}

	if MongoUnlokOffers.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	}

	c.JSON(SuccessResponse{MESSAGE: "Tienda desabilitada"})
}

//UpdatePage Handler for change page information
func UpdatePage(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	PageID := c.Params("page_id")
	UserID := userIDF(c.Get("token"))

	var DataPage PagePut
	var IsOwner IsOwnerShop
	var Subdomain ValidateDomains
	var Domain ValidateDomains

	if errorParse := c.BodyParser(&DataPage); errorParse != nil {
		fmt.Println("Error parsing data", errorParse)
		c.JSON(ErrorResponse{MESSAGE: "Error al parsear información"})
		c.Status(400)
		return
	}

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = 1",
			UserID,
			ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	UpdatePageSQL := sq.Update("pages")

	if DataPage.TemplateType != nil {
		UpdatePageSQL = UpdatePageSQL.Set("template_type", DataPage.TemplateType)
	}

	if DataPage.StyleSheets != nil {
		UpdatePageSQL = UpdatePageSQL.Set("style_sheets", DataPage.StyleSheets)
	}

	if DataPage.ActiveDays != nil {
		UpdatePageSQL = UpdatePageSQL.Set("active_days", DataPage.ActiveDays)
	}

	if DataPage.ImagesDays != nil {
		UpdatePageSQL = UpdatePageSQL.Set("images_days", DataPage.ImagesDays)
	}

	if DataPage.OffersActive != nil {
		UpdatePageSQL = UpdatePageSQL.Set("offers_active", DataPage.OffersActive)
	}

	if DataPage.AcceptCardActive != nil {
		UpdatePageSQL = UpdatePageSQL.Set("accept_card_active", DataPage.AcceptCardActive)
	}

	if len(DataPage.Subdomain) > 0 {
		ErrorSubdomain := sq.Select(
			"shop_id",
			"subdomain",
		).
			From("pages").
			Where("subdomain = ?", DataPage.Subdomain).
			RunWith(database).
			QueryRow().
			Scan(
				&Subdomain.ShopID,
				&Subdomain.Subdomain,
			)

		if ErrorSubdomain != nil {
			if len(Subdomain.ShopID.String) > 0 {
				fmt.Println("Error to validate subdomian or is used", ErrorSubdomain)
				c.JSON(ResponseSubdomain{Subdomain: false})
				c.SendStatus(400)
				return
			}

			UpdatePageSQL = UpdatePageSQL.Set("subdomain", DataPage.Subdomain)
		} else {
			fmt.Println("Error to validate subdomian or is used", ErrorSubdomain)
			c.JSON(ResponseSubdomain{Subdomain: false})
			c.SendStatus(400)
			return
		}
	}

	if len(DataPage.Domain) > 0 {
		ErrorDomain := sq.Select(
			"shop_id",
			"domain",
		).
			From("pages").
			Where("domain = ?", DataPage.Domain).
			RunWith(database).
			QueryRow().
			Scan(
				&Domain.ShopID,
				&Domain.Domain,
			)

		if ErrorDomain != nil {
			if len(Domain.ShopID.String) > 0 {
				fmt.Println("Error to validate subdomian or is used", ErrorDomain)
				c.JSON(ResponseDomain{Domain: false})
				c.SendStatus(400)
				return
			}

			UpdatePageSQL = UpdatePageSQL.Set("domain", DataPage.Domain)
		} else {
			fmt.Println("Error to validate domian or is used", ErrorDomain)
			c.JSON(ResponseDomain{Domain: false})
			c.SendStatus(400)
			return
		}
	}

	_, ErrorUpdatePage := UpdatePageSQL.
		Where("pages_id = ? ", PageID).
		RunWith(database).
		Exec()

	if ErrorUpdatePage != nil {
		fmt.Println(ErrorUpdatePage, "Problem with update information page")
		c.JSON(ErrorResponse{MESSAGE: "Problem with update information page"})
		c.SendStatus(500)
		return
	}

	c.JSON(SuccessResponse{MESSAGE: "Actualizado"})
}

//ActivePage Handler for active pages
func ActivePage(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	PageID := c.Params("page_id")
	UserID := userIDF(c.Get("token"))

	var IsOwner IsOwnerShop

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = 1",
			UserID,
			ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorActivePage := sq.Update("pages").
		Set("active", true).
		Where("pages_id = ? ", PageID).
		RunWith(database).
		Exec()

	if ErrorActivePage != nil {
		fmt.Println(ErrorActivePage, "Problem with Active page")
		c.JSON(ErrorResponse{MESSAGE: "Problem with Active page"})
		c.SendStatus(500)
		return
	}

	c.JSON(SuccessResponse{MESSAGE: "Actualizado"})
}

//DeactivePage Handler for active pages
func DeactivePage(c *fiber.Ctx) {
	ShopID := c.Params("shop_id")
	PageID := c.Params("page_id")
	UserID := userIDF(c.Get("token"))

	var IsOwner IsOwnerShop

	ErrorOwner := sq.Select(
		"shop_id",
	).
		From("shop").
		Where(
			"user_id = ? AND shop_id = ? AND status = 1",
			UserID,
			ShopID,
		).
		RunWith(database).
		QueryRow().
		Scan(
			&IsOwner.ShopID,
		)

	if ErrorOwner != nil {
		fmt.Println("Not is owner or active shop", ErrorOwner)
		c.JSON(ErrorResponse{MESSAGE: "Not is owner or active shop"})
		c.SendStatus(400)
		return
	}

	_, ErrorActivePage := sq.Update("pages").
		Set("active", false).
		Where("pages_id = ? ", PageID).
		RunWith(database).
		Exec()

	if ErrorActivePage != nil {
		fmt.Println(ErrorActivePage, "Problem with Active page")
		c.JSON(ErrorResponse{MESSAGE: "Problem with Active page"})
		c.SendStatus(500)
		return
	}

	c.JSON(SuccessResponse{MESSAGE: "Actualizado"})
}

//UploadImages Handler for upload images
func UploadImages(c *fiber.Ctx) {
	file, err := c.FormFile("file")

	if err != err {
		fmt.Println(err)
	}

	image := helper.UploadImg(file)

	c.JSON(ResponseResultSimple{Result: image.URL})
}

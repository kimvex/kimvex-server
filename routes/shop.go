package routes

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Shops namespace
func Shops() {
	apiRouteShop := apiRoute.Group("/shop")
	apiRouteProfile := apiRoute.Group("/profile")
	apiRouteBase := apiRoute.Group("/")

	//Validations token
	apiRouteProfile.Use("/shops", ValidateRoute)
	apiRouteShop.Use("/:shop_id/offers", ValidateRoute)
	apiRouteShop.Use("/offer/:offer_id", ValidateRoute)

	apiRouteShop.Get("/:shop_id", ShopGet)
	apiRouteShop.Get("/:shop_id/offers", ShopOffers)
	apiRouteShop.Get("/offer/:offer_id", OfferInfo)
	apiRouteShop.Get("/:shop_id/comments", Comments)
	apiRouteShop.Get("/:shop_id/score/:user_id", Score)
	apiRouteShop.Get("/:shop_id/page", Page)
	apiRouteBase.Get("/find/shops/:lat/:lon", FindShops)

	apiRouteBase.Get("/services", Services)
	apiRouteBase.Get("/sub_service/:service_id", SubServices)

	apiRouteProfile.Get("/shops", ProfileShop)
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
		CommentPointer.Comment = &CommentFromSQL.Comment.String
		CommentPointer.CreateDateAt = &CommentFromSQL.CreateDateAt.String
		CommentPointer.UserID = &CommentFromSQL.UserID.String
		CommentPointer.Fullname = &CommentFromSQL.Fullname.String
		CommentPointer.Image = &CommentFromSQL.Image.String

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

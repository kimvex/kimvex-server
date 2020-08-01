package routes

import (
	"fmt"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber"
)

//Shops namespace
func Shops() {
	apiRouteShop := apiRoute.Group("/shop")
	apiRouteProfile := apiRoute.Group("/profile")

	//Validations token
	// apiRouteShop.Use("/:shop_id", ValidateRoute)

	apiRouteShop.Get("/:shop_id", ShopGet)
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
		fmt.Println("error to get shop", shopResultsError)
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
	fmt.Println(Pagination)
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
		"date_init",
		"date_finish",
		"type_charge",
		"GROUP_CONCAT(coalesce(images_shop.url_image, '')) AS images",
	).
		From("shop").
		LeftJoin("images_shop on images_shop.shop_id = shop.shop_id").
		LeftJoin("service_type on service_type.service_type_id = shop.service_type_id").
		LeftJoin("sub_service_type on sub_service_type.sub_service_type_id = shop.sub_service_type_id").
		LeftJoin("shop_schedules on shop_schedules.shop_id = shop.shop_id").
		LeftJoin("plans_pay on plans_pay.shop_id = shop.shop_id").
		LeftJoin("usersk on usersk.user_id = shop.user_id").
		Where("shop.user_id = ? AND plans_pay.expired = ? OR plans_pay.expired IS NULL", userID, 0).
		OrderBy("create_at_shop DESC").
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
		finalMyShops.Images = &MyShopUnStructured[i].Images.String
		finalMyShops.DateInit = &MyShopUnStructured[i].DateInit.String
		finalMyShops.DateFinish = &MyShopUnStructured[i].DateFinish.String
		finalMyShops.TypeCharge = &MyShopUnStructured[i].TypeCharge.String

		listReponse = append(listReponse, finalMyShops)
	}

	c.JSON(ResponseResult{Result: listReponse})
}

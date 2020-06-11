package routes

import (
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber"
)

//Shops namespace
func Shops() {
	apiRouteUser := apiRoute.Group("/shop")

	//Validations token
	// apiRouteUser.Use("/:shop_id", ValidateRoute)

	apiRouteUser.Get("/:shop_id", ShopGet)
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

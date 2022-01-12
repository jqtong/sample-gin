package card

import (
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/app/controller"
	"nucarf.com/store_service/api/app/service/cardService"
	"nucarf.com/store_service/api/request/card"

	//"nucarf.com/store_service/api/request/card"
)

func List(c *gin.Context) {
	var inputParams = card.CardInputParams{}
	err := inputParams.CheckInputParams(c)
	if err != nil {
		controller.Resp(c, nil, err)
		return
	}
	res, err := cardService.List(inputParams)
	controller.Resp(c, res, err)
}

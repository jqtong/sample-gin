package cardService

import (
	"nucarf.com/store_service/api/app/model/card"
	card2 "nucarf.com/store_service/api/request/card"
)

func List(inputParams card2.CardInputParams) (list interface{}, err error) {
	var cardModel = card.Card{}
	list, err = cardModel.List(inputParams.Page, inputParams.Size)
	if err != nil {
		return
	}
	return
}

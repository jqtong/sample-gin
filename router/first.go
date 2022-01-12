package router

import (
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/app/controller/Index"
	"nucarf.com/store_service/api/app/controller/card"
)

func InitBaseRoutes(e *gin.Engine) {
	//b := e.Group("/base")

	e.GET("/ping", index.Pong)
	e.GET("/card_list", card.List)

}

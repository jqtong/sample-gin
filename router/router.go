package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/middleware"
)


func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Cross())
	r.Use(gin.Recovery())

	pprof.Register(r)

	InitBaseRoutes(r)

	return r
}

package main

import (
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/conf"
	"nucarf.com/store_service/api/conf/initialize"
	"nucarf.com/store_service/api/router"
	"strconv"
	"strings"
)

func main() {

	// load config
	conf.Load()

	// set run mode
	gin.SetMode(initialize.ServerConf.RunMode)

	router := router.InitRouter()

	// address
	addr := strings.Builder{}
	addr.Grow(32)

	addr.WriteString(initialize.ServerConf.Host)
	addr.WriteString(":")
	addr.WriteString(strconv.Itoa(initialize.ServerConf.Port))

	/*go func() {
		log.Println(http.ListenAndServe("0.0.0.0:17000", nil))
	}()*/

	if err := router.Run(addr.String()); err != nil {
		panic(err)
	}
}

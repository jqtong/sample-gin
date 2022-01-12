package index

import (
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/app/controller"
)

// Pong returns server status
func Pong(c *gin.Context) {
	controller.Resp(c, "pong", nil)
}

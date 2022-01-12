package card

import (
	"github.com/gin-gonic/gin"
	"nucarf.com/store_service/api/errorcode"
	"nucarf.com/store_service/api/request"
	"strings"
)

//CardInputParams struct
type CardInputParams struct {
	CardId int    `json:"card_id" validate:"min=0,max=1000000" form:"card_id"  comment:"油卡ID"`
	Name   string `json:"name" form:"name" comment:"油卡名称"`
	request.PageInput
	request.SizeInput
}

// CheckInputParams 检查词组添加参数是否正确
func (ci *CardInputParams) CheckInputParams(c *gin.Context) error {
	//参数赋值
	if err := c.ShouldBind(ci); err != nil {
		return err
	}


	//参数基本校验
	if errs, err := request.Validate(ci); err != nil {
		return errorcode.New(errorcode.ErrParams, strings.Join(errs, ","), nil)
	}
	return nil
}
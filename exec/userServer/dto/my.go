package dto

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type OperateAddress struct {
	Mark     int    `form:"mark" json:"mark" validate:"omitempty,min=1" zh:"标识"`
	Province string `form:"province" json:"province" validate:"min=1" zh:"省"`
	City     string `form:"city" json:"city" validate:"min=1" zh:"市"`
	Region   string `form:"region" json:"region" validate:"min=1" zh:"区"`
	Address  string `form:"address" json:"address" validate:"min=1" zh:"地址"`
	Name     string `form:"name" json:"name" validate:"min=1" zh:"名称"`
	Phone    string `form:"phone" json:"phone" validate:"len=11" zh:"手机号"`
	Def      bool   `form:"def" json:"def" validate:"" zh:"是否是默认"`
}

func (o *OperateAddress) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return BindingValidParams(c, o, reflect.TypeOf(*o))
}

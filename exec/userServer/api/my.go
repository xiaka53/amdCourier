package api

import (
	"amdCourier/dao"
	"amdCourier/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type my struct{}

func MyRouterInit(r *gin.RouterGroup) {
	var m my
	r.GET("info", m.info)
	r.GET("address_list", m.addressList)
	r.POST("operateAddress", m.operateAddress)
}

// @Summary 我的信息
// @Tags 我的页面
// @Produce  json
// @Success 200 {object} my_info
// @Router /my/info [Get]
func (m my) info(c *gin.Context) {
	var (
		user dao.User
		info my_info
	)

	user.Id = c.GetUint("uuid")
	if user.Id < 1 {
		middleware.ResponseError(c, middleware.NoLoginCode, errors.New("未登陆"))
	}
	if err := (&user).First(); err != nil {
		middleware.ResponseError(c, middleware.NoLoginCode, errors.New("不存在的账户"))
	}
	info = my_info{
		Name:   user.Name,
		Avatar: user.Avatar,
		Number: user.Number,
	}
	middleware.ResponseSuccess(c, info)
}

type my_info struct {
	Name   string `json:"name"`   // 名称
	Avatar string `json:"avatar"` // 头像
	Number string `json:"number"` // 编号
}

// @Summary 地址簿
// @Tags 我的页面
// @Produce  json
// @Success 200 {object} my_addressList
// @Router /my/addressList [Get]
func (m my) addressList(c *gin.Context) {
	var (
		user        dao.User
		addressList []dao.Address
		info        []my_addressList
	)

	user.Id = c.GetUint("uuid")
	if user.Id < 1 {
		middleware.ResponseError(c, middleware.NoLoginCode, errors.New("未登陆"))
	}
	if err := (&user).First(); err != nil {
		middleware.ResponseError(c, middleware.NoLoginCode, errors.New("不存在的账户"))
	}
	addressList, _ = (&dao.Address{Uid: user.Id, Del: false}).Find()
	for _, v := range addressList {
		info = append(info, my_addressList{
			Province: v.Province,
			City:     v.City,
			Region:   v.Region,
			Address:  v.Address,
			Name:     v.Name,
			Phone:    v.Phone[:3] + "****" + v.Phone[7:],
			Def:      v.Def,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type my_addressList struct {
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	Region   string `json:"region"`   // 区
	Address  string `json:"address"`  // 地址
	Name     string `json:"name"`     // 名称
	Phone    string `json:"phone"`    // 手机号
	Def      bool   `json:"def"`      // 是否默认
}

// @Summary 创建/修改地址
// @Tags 我的页面
// @Produce  json
// @Param mark query int false "标识"
// @Param province query string required "省"
// @Param city query string required "市"
// @Param region query string required "区"
// @Param address query string required "address"
// @Param name query string required "名称"
// @Param phone query string required "手机号"
// @Param def query bool required "是否默认"
// @Success 200 {string} "success"
// @Router /my/operateAddress [Post]
func (m my) operateAddress(c *gin.Context) {

}

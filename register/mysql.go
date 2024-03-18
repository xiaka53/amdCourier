package register

import (
	"amdCourier/dao"
	"amdCourier/server"
)

// 数据库表格注册
func RegisterMysql() {
	RegisterMysqlUser()
	RegisterMysqlRider()
	RegisterMysqlTableConfig()
}

// 所有用户相关表格注册
func RegisterMysqlUser() {
	server.RegisterMysqlTable(
		&dao.Address{},
		&dao.User{},
		&dao.Regional{},
	)
}

// 所有骑手相关表格注册
func RegisterMysqlRider() {
	server.RegisterMysqlTable()
}

// 配置表数据初始化
func RegisterMysqlTableConfig() {
	RegisterRegional()
}

// 地区快递价格初始化
func RegisterRegional() {
	var (
		all         []dao.Regional
		zones       [5][]string
		firstWeight [5]float64
		renewal     [5]float64
	)
	zones[0] = []string{"浙江", "安徽", "江苏", "上海"}
	zones[1] = []string{"福建", "江西", "湖北", "湖南", "河南", "河北", "广东", "山东", "北京", "天津"}
	zones[2] = []string{"广西", "陕西", "山西", "辽宁", "四川", "重庆", "重庆", "贵州", "黑龙江", "吉林", "云南"}
	zones[3] = []string{"海南", "甘肃", "青海", "内蒙古", "宁夏"}
	zones[4] = []string{"西藏", "新疆"}

	firstWeight = [5]float64{5, 6, 8, 10, 12}
	renewal = [5]float64{2, 2, 3, 5, 18}

	for zone, provinces := range zones {
		for _, province := range provinces {
			all = append(all, dao.Regional{
				Province:    province,
				FirstWeight: firstWeight[zone],
				Renewal:     renewal[zone],
			})
		}
	}

	(&dao.Regional{}).Register(all)
}

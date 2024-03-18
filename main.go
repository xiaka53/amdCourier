package main

import (
	"amdCourier/dao"
	"amdCourier/register"
	"fmt"
)

func main() {
	register.RegisterMysql()
	addr := dao.Address{
		Uid:     1,
		Name:    "李华",
		Phone:   "19999999999",
		Region:  "北京市-北京市-望京区",
		Address: "soho 33号楼2109",
		Def:     true,
		Del:     false,
	}
	if err := (&addr).Create(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("a")
}

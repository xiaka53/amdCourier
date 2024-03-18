package dao

import (
	"amdCourier/server"
	"gorm.io/gorm"
	"log"
)

type Regional struct {
	Id          uint    `gorm:"column:id;primaryKey;autoIncrement;size:12;comment:'ID'"`
	Province    string  `gorm:"column:province;not null;size:20;comment:'省份'"`
	FirstWeight float64 `gorm:"column:first_weight;not null;size:5;comment:'首重价格'"`
	Renewal     float64 `gorm:"column:renewal;not null;size:5;comment:'续重(公斤/元)'"`
}

func (Regional) TableName() string {
	return "regional"
}

func (r *Regional) Conn() *gorm.DB {
	return server.MSlConn().Table(r.TableName())
}

func (r *Regional) First() error {
	return r.Conn().Where(r).First(r).Error
}

func (r *Regional) Register(data []Regional) {
	if err := r.First(); err == nil {
		return
	}
	if err := r.Conn().Create(data).Error; err != nil {
		log.Fatalf("地区收费标准初始化失败：%v", err)
	}
}

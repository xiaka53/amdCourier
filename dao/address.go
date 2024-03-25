package dao

import (
	"amdCourier/server"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	Id        uint       `gorm:"column:id;primaryKey;autoIncrement;size:12;comment:'ID'"`
	Uid       uint       `gorm:"column:uid;not null;comment:'用户ID'"`
	Name      string     `gorm:"column:name;not null;size:20;comment:'姓名'"`
	Phone     string     `gorm:"column:phone;not null;size:11;comment:'手机号'"`
	Province  string     `gorm:"column:province;not null;size:80;comment:'省'"`
	City      string     `gorm:"column:city;not null;size:80;comment:'市'"`
	Region    string     `gorm:"column:region;not null;size:80;comment:'区'"`
	Address   string     `gorm:"column:address;not null;size:100;comment:'地址'"`
	Def       bool       `gorm:"column:def;not null;comment:'是否是默认'"`
	Del       bool       `gorm:"column:del;not null;comment:'是否删除'"`
	CreatedAt *time.Time `gorm:"column:created_at;not null;comment:'创建时间'"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;comment:'更新时间'"`
	DeletedAt *time.Time `gorm:"column:deleted_at;comment:'删除时间'"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) Conn() *gorm.DB {
	return server.MSlConn().Table(a.TableName())
}

func (a *Address) First() error {
	return a.Conn().Where(a).First(a).Error
}

func (a *Address) Find() (data []Address, err error) {
	err = a.Conn().Where(a).Find(&data).Error
	return
}

func (a *Address) Create() error {
	// TODO 地址追踪查询
	if a.Id > 0 {
		if err := a.Conn().Where("id=?", a.Id).Updates(a).Error; err != nil {
			return err
		}
	}
	a.DeletedAt = nil
	return a.Conn().Create(a).Error
}

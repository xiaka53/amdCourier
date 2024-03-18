package dao

import (
	"amdCourier/server"
	"gorm.io/gorm"
	"time"
)

type StatusEnum string

const (
	StatusNormal  StatusEnum = "normal"  // 正常
	StatusDisable StatusEnum = "disable" // 禁用
)

type User struct {
	Id        uint       `gorm:"column:id;primaryKey;autoIncrement;size:12;comment:'ID'"`
	OpenId    string     `gorm:"column:name;not null;unique;size:100;comment:'open id'"`
	Avatar    string     `gorm:"column:name;not null;size:80;comment:'头像'"`
	Name      string     `gorm:"column:name;not null;size:20;comment:'姓名'"`
	Phone     string     `gorm:"column:phone;not null;size:11;comment:'手机号'"`
	Number    string     `gorm:"column:number;not null;size:11;comment:'编号'"`
	Status    StatusEnum `gorm:"column:status;not null;comment:'状态'"`
	CreatedAt *time.Time `gorm:"column:created_at;not null;comment:'创建时间'"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;comment:'更新时间'"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Conn() *gorm.DB {
	return server.MSlConn().Table(u.TableName())
}

func (u *User) First() error {
	return u.Conn().First(u).Error
}

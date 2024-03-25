package dao

import "gorm.io/gorm"

type City struct {
	Id uint `gorm:"column:id;primaryKey;autoIncrement;size:12;comment:'ID'"`
}

func (City) TableName() string {
	//TODO implement me
	panic("implement me")
}

func (City) Conn() *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (City) First() error {
	//TODO implement me
	panic("implement me")
}

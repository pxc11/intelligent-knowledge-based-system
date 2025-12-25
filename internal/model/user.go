package model

import "time"

type User struct {
	ID         uint
	Username   string
	Password   string
	Createtime time.Time `gorm:"autoCreateTime"`
	Updatetime time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

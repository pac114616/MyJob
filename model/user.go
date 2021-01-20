package model

import "github.com/jinzhu/gorm"

type User struct {
	//User内嵌了gorm.Model，内置了ID、CreatedAt、UpdatedAt、DeletedAt属性，
	//同时Create的时候会自动设置CreatedAt、UpdatedAt，Update的时候会自动更新UpdatedAt
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

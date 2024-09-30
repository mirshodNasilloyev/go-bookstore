package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "mnasilloyev:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", "mnasilloyev:%40Nm24811913@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "mnasilloyev:@Nm24811913@/simplerest?charset=utf8&parseTime=True&loc=Local")

	// d, err := gorm.Open("mysql", "root@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

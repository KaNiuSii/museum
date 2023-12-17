package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:3213@tcp(localhost:9999)/goauth?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate()
}
package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn_go/learnGo/model"
)

var DB *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "go_learn"
	username := "alphasu"
	password := "990211"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username,
		password, host, port, database, charset)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	} else {
		_ = db.AutoMigrate(&model.User{})
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

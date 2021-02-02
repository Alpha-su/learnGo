package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null; unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		// 如果名称没有传，赋值一个随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Print(name, telephone, password)
		// 判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(200, gin.H{"code": 200, "msg": "注册成功"})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "go_learn"
	username := "alphasu"
	password := "Sgl@123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username,
		password, host, port, database, charset)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	} else {
		_ = db.AutoMigrate(&User{})
	}
	return db
}

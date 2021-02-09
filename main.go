package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"learn_go/learnGo/common"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
	panic(r.Run())
}

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		//panic()
	}
}

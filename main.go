package main

import (
	"github.com/gin-gonic/gin"
	"learn_go/learnGo/common"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

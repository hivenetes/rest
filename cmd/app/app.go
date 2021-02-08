package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp ...
func StartApp() {
	fmt.Println("Starting WebServer")
	MapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

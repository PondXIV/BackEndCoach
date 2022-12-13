package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewDemoController(router *gin.Engine) {
	ping := router.Group("/demo")
	{
		ping.GET("/", demoHello)

	}
}
func demoHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ฟ่ดนะจ๊ะ",
	})
}

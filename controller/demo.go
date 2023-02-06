package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func NewDemoController(router *gin.Engine) {
	ping := router.Group("/pp")
	{
		ping.GET("/", demoHello)
		ping.POST("/", getCourseByID)
	}
}
func demoHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ฟ่ดนะจ๊ะ",
	})
}

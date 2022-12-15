package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserController(router *gin.Engine) {
	ping := router.Group("/user")
	{
		ping.POST("/login", loginPostBody)
		ping.POST("/registerCus", registerCus)
	}
}
func loginPostBody(ctx *gin.Context) {
	
}
func registerCus(ctx *gin.Context) {
	
}

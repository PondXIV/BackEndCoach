package main

import (
	"backEndGo/controller"
	"net/http"

	"github.com/gin-gonic/gin"

)

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Api is working...")
	})
	controller.NewDemoController(router)
	controller.NewUserController(router)
	router.Run()
}

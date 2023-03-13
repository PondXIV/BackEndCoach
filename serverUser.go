package main

import (
	usercontroller "backEndGo/controller/UserController"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserStartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Api is working...")
	})
	//controller.NewDemoController(router)
	usercontroller.NewCourseController(router)

	router.Run()
}

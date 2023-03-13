package main

import (
	"backEndGo/controller"
	coursecontroller "backEndGo/controller/CoachController/CourseController"
	usercontroller "backEndGo/controller/UserController"

	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Api is working...")

	})
	//controller.NewDemoController(router)
	controller.NewUserController(router)
	coursecontroller.NewCourseController(router)
	usercontroller.NewCourseController(router)
	router.Run()
}

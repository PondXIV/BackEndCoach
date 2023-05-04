package main

import (
	"backEndGo/controller"
	clipcontroller "backEndGo/controller/CoachController/ClipController"
	coursecontroller "backEndGo/controller/CoachController/CourseController"
	foodincoursecontooler "backEndGo/controller/CoachController/CourseController/Food_in_CourseContooler"
	foodcontroller "backEndGo/controller/CoachController/FoodContoller"
	usercontroller "backEndGo/controller/UserController"
	coursectl "backEndGo/controller/UserController/Course_CTL"
	reviewctl "backEndGo/controller/UserController/Review_CTL"
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
	usercontroller.NewUserController(router)
	foodcontroller.NewListFoodController(router)
	foodincoursecontooler.NewFoodController(router)
	clipcontroller.NewListClipController(router)
	reviewctl.NewReviewController(router)
	coursectl.NewCourseController(router)
	router.Run()
}

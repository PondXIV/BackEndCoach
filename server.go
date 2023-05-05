package main

import (
	clipcontroller "backEndGo/controller/CoachController/ClipController"
	coursecontroller "backEndGo/controller/CoachController/CourseController"
	foodincoursecontooler "backEndGo/controller/CoachController/CourseController/Food_in_CourseContooler"
	foodcontroller "backEndGo/controller/CoachController/FoodContoller"
	usercontroller "backEndGo/controller/UserController"
	coursectl "backEndGo/controller/UserController/Course_CTL"
	reviewctl "backEndGo/controller/UserController/Review_CTL"
	auth "backEndGo/controller/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Api is working...")

	})
	//controller.NewDemoController(router)
	auth.NewAuthController(router)
	coursecontroller.NewCourseController(router)
	usercontroller.NewUserController(router)
	foodcontroller.NewListFoodController(router)
	foodincoursecontooler.NewFoodController(router)
	clipcontroller.NewListClipController(router)
	reviewctl.NewReviewController(router)
	coursectl.NewCourseController(router)
	router.Run()
}

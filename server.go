package main

import (

	//Course
	ctrlClip "backEndGo/controller/CoachController/ClipController/Clip_in_Course"
	ctrlCourse "backEndGo/controller/CoachController/CourseController"
	ctrlDays "backEndGo/controller/CoachController/DaysController"
	ctrlFood "backEndGo/controller/CoachController/FoodContoller/Food_in_Course"

	//Coach
	ctrlListClip "backEndGo/controller/CoachController/ClipController/Clip_in_Coach"
	ctrlListFood "backEndGo/controller/CoachController/FoodContoller/Food_in_Coach"

	//User
	usercontroller "backEndGo/controller/UserController"
	coachctl "backEndGo/controller/UserController/Coach_CTL"
	coursectl "backEndGo/controller/UserController/Course_CTL"
	reviewctl "backEndGo/controller/UserController/Review_CTL"

	auth "backEndGo/controller/auth"

	//Buying
	buyctl "backEndGo/controller/UserController/Buy_CTL"

	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Api is working...")

	})

	///Coach
	ctrlListFood.NewListFoodController(router)
	ctrlListClip.NewListClipController(router)
	coachctl.NewCoachController(router)

	//Course
	ctrlCourse.NewCourseController(router)
	ctrlFood.NewFoodController(router)
	ctrlDays.NewDayController(router)
	ctrlClip.NewClipController(router)

	//User
	auth.NewAuthController(router)
	usercontroller.NewUserController(router)
	reviewctl.NewReviewController(router)
	coursectl.NewCourseController(router)

	//Buy
	buyctl.NewBuyingCourseController(router)

	router.Run()
}

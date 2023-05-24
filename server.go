package main

import (
	ctrlBuycourse "backEndGo/controller/BuyCourseController"
	ctrlListClip "backEndGo/controller/CoachController/ClipController/Clip_in_Coach"
	ctrlCourse "backEndGo/controller/CoachController/CourseController"
	ctrlDays "backEndGo/controller/CoachController/DaysController"
	ctrlListFood "backEndGo/controller/CoachController/FoodContoller/Food_in_Coach"
	ctrlFood "backEndGo/controller/CoachController/FoodContoller/Food_in_Course"

	usercontroller "backEndGo/controller/UserController"
	coachctl "backEndGo/controller/UserController/Coach_CTL"
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

	///Coach
	ctrlListFood.NewListFoodController(router)
	ctrlListClip.NewListClipController(router)
	coachctl.NewCoachController(router)

	//Course
	ctrlCourse.NewCourseController(router)
	ctrlFood.NewFoodController(router)
	ctrlDays.NewDayController(router)

	//User
	auth.NewAuthController(router)
	usercontroller.NewUserController(router)
	reviewctl.NewReviewController(router)
	coursectl.NewCourseController(router)

	//Buy
	ctrlBuycourse.NewBuyCourseController(router)

	router.Run()
}

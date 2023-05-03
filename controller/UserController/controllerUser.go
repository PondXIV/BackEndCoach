package usercontroller

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nameCoachDateService = userservice.NewCoachByNameDataService()
var reviewDataService = userservice.NewReviewDataService()
var customerService = userservice.NewUserDataService()
var updatecustomerService = userservice.NewUserDataService()
var modelsCustomer = models.Customer{}
var mycourseService = userservice.NewMyCourseDataService()

func NewCourseController(router *gin.Engine) {
	nameCoach := router.Group("/user")
	{
		nameCoach.GET("/nameCourse/:name", GetCoachByName)
		nameCoach.GET("/review/courseID/:coID", GetReviewByCoID)
		nameCoach.GET("/customerID/:uid", Customer)
		nameCoach.PUT("/customerID/:uid", updateCustomer)
		nameCoach.GET("/course/customerID/:uid", GetMycourse)

	}

}
func GetMycourse(ctx *gin.Context) {
	cusID := ctx.Param("uid")
	uid, err := strconv.Atoi(cusID)
	course, err := mycourseService.ServiceGetMycourse(uid)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, course)
}
func updateCustomer(ctx *gin.Context) {
	cusID := ctx.Param("uid")
	uid, _ := strconv.Atoi(cusID)
	err := ctx.ShouldBindJSON(&modelsCustomer)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	// // process
	rowsAffected := updatecustomerService.ServiceUpdateCustomer(uid, &modelsCustomer)

	if rowsAffected == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "1",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "0",
		})
	}
}

func GetCoachByName(ctx *gin.Context) {
	name := ctx.Param("name")
	coachs, err := nameCoachDateService.ServiceGetNameCoach(name)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, coachs)

}
func Customer(ctx *gin.Context) {
	cusID := ctx.Param("uid")
	uid, err := strconv.Atoi(cusID)
	customer, err := customerService.ServiceGetUserByUid(uid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, customer)
}
func GetReviewByCoID(ctx *gin.Context) {

	courseID := ctx.Param("coID")
	coID, err := strconv.Atoi(courseID)
	review, err := reviewDataService.ServiceGetReviewByCoId(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, review)

}

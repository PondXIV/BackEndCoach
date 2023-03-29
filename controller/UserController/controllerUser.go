package usercontroller

import (
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nameCoachDateService = userservice.NewCoachByNameDataService()
var reviewDataService = userservice.NewReviewDataService()
var customerService = userservice.NewUserDataService()

func NewCourseController(router *gin.Engine) {
	nameCoach := router.Group("/user2")
	{
		nameCoach.GET("/getCoachByName/:name", GetCoachByName)
		nameCoach.GET("/getReviewByCoID/:coID", GetReviewByCoID)
		nameCoach.GET("/customer/:uid", Customer)

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

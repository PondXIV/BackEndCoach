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
	nameCoach := router.Group("/user2")
	{
		nameCoach.GET("/getCoachByName/:name", GetCoachByName)
		nameCoach.GET("/getReviewByCoID/:coID", GetReviewByCoID)
		nameCoach.GET("/customer/:uid", Customer)
		nameCoach.PUT("/updateCus", updateCustomer)
		nameCoach.GET("/mycourse", GetMycourse)

	}

}
func GetMycourse(ctx *gin.Context) {
	// Receive Query Params
	quid := ctx.Query("uid")
	qcoid := ctx.Query("coid")
	qname := ctx.Query("name")
	uid, err := strconv.Atoi(quid)
	coid, err := strconv.Atoi(qcoid)
	// If params not exist string = "", int = 0

	course, err := mycourseService.ServiceGetMycourse(uid, coid, qname)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, course)

}
func updateCustomer(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCustomer)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	// // process
	rowsAffected := updatecustomerService.ServiceUpdateCustomer(&modelsCustomer)

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

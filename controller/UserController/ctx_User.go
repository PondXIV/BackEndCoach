package usercontroller

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var customerService = userservice.NewUserDataService()
var updatecustomerService = userservice.NewUserDataService()
var modelsCustomer = models.Customer{}

func NewUserController(router *gin.Engine) {
	nameCoach := router.Group("/user")
	{
		//เอา repo GetCourse ของเค้าได้
		// nameCoach.GET("/nameCourse/:name", GetCoachByName)
		//nameCoach.GET("/review/courseID/:coID", GetReviewByCoID)
		nameCoach.GET("/customerID/:uid", Customer)
		nameCoach.PUT("/customerID/:uid", updateCustomer)
		//ฟ่ด repo GetCourse ของเค้า bid IS NULL นะ ต้องเอาอันอืน
		// nameCoach.GET("/course/customerID/:uid", GetMycourse)

	}

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

func Customer(ctx *gin.Context) {
	cusID := ctx.Param("uid")
	uid, err := strconv.Atoi(cusID)
	customer, err := customerService.ServiceGetUserByUid(uid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, customer)
}

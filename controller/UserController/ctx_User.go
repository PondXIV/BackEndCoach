package walletctl

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
		nameCoach.GET("", Customer)
		nameCoach.PUT(":uid", updateCustomer)
		//ฟ่ด repo GetCourse ของเค้า bid IS NULL นะ ต้องเอาอันอืน
		// nameCoach.GET("/course/customerID/:uid", GetMycourse)

	}

}

func updateCustomer(ctx *gin.Context) {
	cusID := ctx.Param("uid")
	uid, errs := strconv.Atoi(cusID)
	//err := ctx.ShouldBindJSON(&modelsCustomer)
	// fmt.Printf("%v", cus)
	if errs != nil {
		panic(errs)
	}
	// // process
	err := ctx.ShouldBindJSON(&modelsCustomer)
	// fmt.Printf("%v", cus)
	if err != nil {
		error400(ctx)
	}
	rowsAffected, err := updatecustomerService.ServiceUpdateCustomer(uid, &modelsCustomer)

	if err != nil {
		error400(ctx)

	} else {
		if rowsAffected == 1 {
			outputOne(ctx)

		} else {
			outputSoon(ctx)
		}
	}
}

func Customer(ctx *gin.Context) {
	cusID := ctx.Query("uid")
	uid, err := strconv.Atoi(cusID)
	customer, err := customerService.ServiceGetUserByUid(uid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, customer)
}

func error400(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":   "400",
		"result": "null",
	})
}
func error500(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":   "500",
		"result": "null",
	})
}

func outputOne(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "1",
	})
}
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}

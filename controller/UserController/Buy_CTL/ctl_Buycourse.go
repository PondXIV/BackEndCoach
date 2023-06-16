package buyctl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var buycourseService = userservice.NewBuyCourseDataService()
var modelsBuying = models.Buying{}
var modelGBprime = models.Gbprimpay{}

func NewBuyingCourseController(router *gin.Engine) {
	buy := router.Group("/buycourse")
	{
		buy.POST("/:coID", buycourse)
		buy.POST("", getgbprime)
	}

}

func getgbprime(ctx *gin.Context) {
	jsonDto := modelGBprime
	err := ctx.ShouldBindJSON(&jsonDto)

	if err != nil {

		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{

		"ResultCode":     jsonDto.ResultCode,
		"ReferenceNo":    jsonDto.ReferenceNo,
		"GbpReferenceNo": jsonDto.GbpReferenceNo,
	})
}

func buycourse(ctx *gin.Context) {
	courseID := ctx.Param("coID")
	coID, errs := strconv.Atoi(courseID)

	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsBuying)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err,
			})
		}
		// else {
		// 	error500(ctx)
		// }
	}
	rowsAffected, err := userservice.NewBuyingCourseDataService().ServiceBuyCourse(coID, &modelsBuying)
	if err != nil {
		if http.StatusBadRequest == 400 {
			error400(ctx)
		} else {
			error500(ctx)
		}

	} else {
		if rowsAffected >= 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   "200",
				"result": rowsAffected,
			})

		} else {
			outputSoon(ctx)
		}
	}

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

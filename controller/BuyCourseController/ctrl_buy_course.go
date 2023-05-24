package buycoursecontroller

import (
	"backEndGo/models"
	buycourse "backEndGo/service/BuyCourse"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var modelsBuying = models.Buying{}

func NewBuyCourseController(router *gin.Engine) {
	buy := router.Group("/buy")
	{
		buy.POST("/:coID", buyPostBody)
	}

}
func buyPostBody(ctx *gin.Context) {
	courseID := ctx.Param("coID")

	coID, errs := strconv.Atoi(courseID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsBuying)
	// fmt.Printf("%v", cus)
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

	rowsAffected, err := buycourse.NewBuyCourseDataService().ServiceBuyCourse(coID, &modelsBuying)
	if err != nil {
		if http.StatusBadRequest == 400 {
			error400(ctx)
		} else {
			error500(ctx)
		}

	} else {
		if rowsAffected == 1 {
			outputOne(ctx)

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

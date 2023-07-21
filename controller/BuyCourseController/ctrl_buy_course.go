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
var show = buycourse.NewBuyingDataService()

func NewBuyCourseController(router *gin.Engine) {
	buy := router.Group("/buy")
	{
		buy.GET("", getBuying)
		buy.GET("/user/:cid", getCourseByUser)
		buy.POST("/:coID", buyPostBody)
	}

}
func getCourseByUser(ctx *gin.Context) {

	coachId := ctx.Param("cid")

	cid, _ := strconv.Atoi(coachId)

	course, err := show.SeviceGetCourseByUser(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func getBuying(ctx *gin.Context) {
	quid := ctx.Query("uid")
	qcoID := ctx.Query("coID")
	qcid := ctx.Query("cid")

	uid, err := strconv.Atoi(quid)
	coID, err := strconv.Atoi(qcoID)
	cid, err := strconv.Atoi(qcid)

	Buys, err := show.GetBuying(uid, coID, cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, Buys)

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

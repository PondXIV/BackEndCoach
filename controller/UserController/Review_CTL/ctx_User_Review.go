package reviewctl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var reviewDataService = userservice.NewReviewDataService()

func NewReviewController(router *gin.Engine) {
	review := router.Group("/review")
	{
		review.GET("", GetReviewByCoID)
		review.GET(":cid", GetReviewByCid)
		review.POST(":bid", InsertReview)

	}

}
func GetReviewByCid(ctx *gin.Context) {
	fmt.Println()
	coachID := ctx.Param("cid")
	cid, err := strconv.Atoi(coachID)
	review, err := reviewDataService.ServiceGetReviewByCid(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, review)

}
func GetReviewByCoID(ctx *gin.Context) {

	courseID := ctx.Query("coID")
	coID, err := strconv.Atoi(courseID)
	review, err := reviewDataService.ServiceGetReviewByCoId(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, review)

}
func InsertReview(ctx *gin.Context) {
	billID := ctx.Param("bid")

	bid, errs := strconv.Atoi(billID)
	if errs != nil {
		panic(errs)
	}
	modelReview := models.Review{}
	err := ctx.ShouldBindJSON(&modelReview)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := reviewDataService.ServiceInsertReview(bid, &modelReview)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}
	} else {
		if rowsAffected >= 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   "200",
				"result": strconv.Itoa(int(rowsAffected)),
			})

		} else {
			outputSoon(ctx)
		}
	}

}
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}

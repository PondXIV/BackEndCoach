package reviewctl

import (
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var reviewDataService = userservice.NewReviewDataService()

func NewReviewController(router *gin.Engine) {
	review := router.Group("/review")
	{
		review.GET("", GetReviewByCoID)

	}

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

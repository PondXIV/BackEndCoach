package foodincoursecontooler

import (
	foodincourse "backEndGo/service/CoachService/CourseSV/Food_in_Course"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var foodSV = foodincourse.NewFoodDataService()

func NewFoodController(router *gin.Engine) {
	food := router.Group("/food")
	{
		food.GET("/did/:did", getFoodByDid)
		food.GET("/coID/:coID", getFoodByIDCourse)
	}

}
func getFoodByDid(ctx *gin.Context) {
	daydid := ctx.Param("did")

	did, err := strconv.Atoi(daydid)
	course, err := foodSV.SeviceGetFoodByDid(did)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func getFoodByIDCourse(ctx *gin.Context) {
	daydid := ctx.Param("coID")

	coID, err := strconv.Atoi(daydid)
	course, err := foodSV.SeviceGetFoodByIDCourse(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}

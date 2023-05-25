package foodcontroller

import (
	foodincourse "backEndGo/service/CoachService/FoodSV/Food_in_Course"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var foodSV = foodincourse.NewFoodDataService()

func NewFoodController(router *gin.Engine) {
	food := router.Group("/food")
	{
		food.GET("", getFood)

	}

}
func getFood(ctx *gin.Context) {
	qfid := ctx.Query("fid")
	qifid := ctx.Query("ifid")
	qdid := ctx.Query("did")

	fid, err := strconv.Atoi(qfid)
	ifid, err := strconv.Atoi(qifid)
	did, err := strconv.Atoi(qdid)

	course, err := foodSV.SeviceGetFood(fid, ifid, did)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}

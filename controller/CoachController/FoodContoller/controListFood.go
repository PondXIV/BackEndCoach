package foodcontroller

import (
	foodsv "backEndGo/service/CoachService/FoodSV"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var listFoodDateService = foodsv.NewListFoodDataService()

func NewListFoodController(router *gin.Engine) {
	listFood := router.Group("/listFood")
	{
		listFood.GET("/:cid", getListFoodeByID)

	}

}
func getListFoodeByID(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, err := strconv.Atoi(coachID)
	foods, err := listFoodDateService.SeviceGetFoodByIDCoach(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, foods)

}

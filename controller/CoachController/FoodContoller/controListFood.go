package foodcontroller

import (
	"backEndGo/models"
	foodsv "backEndGo/service/CoachService/FoodSV"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var listFoodDateService = foodsv.NewListFoodDataService()
var insertListFoodDataService = foodsv.NewInsertListFoodDataService()
var modelsListFood = models.ListFood{}

func NewListFoodController(router *gin.Engine) {
	listFood := router.Group("/listFood")
	{
		listFood.GET("/:cid", getListFoodeByID)
		listFood.POST("/insertListFood", insertListFood)

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
func insertListFood(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsListFood)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	rowsAffected := insertListFoodDataService.SeviceInsertListFoodByID(&modelsListFood)
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

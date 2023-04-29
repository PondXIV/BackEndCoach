package foodcontroller

import (
	"backEndGo/models"
	foodsv "backEndGo/service/CoachService/FoodSV"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var listFoodDateService = foodsv.NewListFoodDataService()
var insertListFoodDataService = foodsv.NewInsertListFoodDataService()
var updateListFoodDataService = foodsv.NewUpdateListFoodDataService()
var modelsListFood = models.ListFood{}

func NewListFoodController(router *gin.Engine) {
	listFood := router.Group("/listFood")
	{
		listFood.GET("/:cid", getListFoodeByID)
		listFood.POST("/insertListFood", insertListFood)
		listFood.PUT("/updateListFood", updateListFood)
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
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			error400(ctx)
		} else {
			error500(ctx)
		}
	} else {
		rowsAffected := insertListFoodDataService.SeviceInsertListFoodByID(&modelsListFood)
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

}
func updateListFood(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsListFood)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			error400(ctx)
		} else {
			error500(ctx)
		}
	} else {
		rowsAffected, err := updateListFoodDataService.ServiceUpdateListFood(&modelsListFood)
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

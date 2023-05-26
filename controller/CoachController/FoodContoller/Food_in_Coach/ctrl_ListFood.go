package foodcontroller

import (
	"backEndGo/models"
	foodsv "backEndGo/service/CoachService/FoodSV/Food_in_Coach"
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
		listFood.GET("", getListFoode)
		listFood.POST("/coachID/:cid", insertListFood)
		listFood.PUT("/foodID/:ifid", updateListFood)
		listFood.DELETE("/foodID/:ifid", deleteListFood)
	}

}
func getListFoode(ctx *gin.Context) {

	qcid := ctx.Query("cid")
	qifid := ctx.Query("ifid")
	qname := ctx.Query("name")

	ifid, err := strconv.Atoi(qifid)
	cid, err := strconv.Atoi(qcid)

	foods, err := listFoodDateService.SeviceGetFood(ifid, cid, qname)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, foods)

}

func insertListFood(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, errs := strconv.Atoi(coachID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsListFood)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

	}
	rowsAffected, err := insertListFoodDataService.SeviceInsertListFoodByID(cid, &modelsListFood)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})

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
func updateListFood(ctx *gin.Context) {
	foodID := ctx.Param("ifid")

	ifid, errs := strconv.Atoi(foodID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsListFood)
	// fmt.Printf("%v", cus)
	if err != nil {
	} else {
		rowsAffected, err := updateListFoodDataService.ServiceUpdateListFood(ifid, &modelsListFood)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})

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

}

func deleteListFood(ctx *gin.Context) {
	foodID := ctx.Param("ifid")

	ifid, errs := strconv.Atoi(foodID)

	if errs != nil {
	} else {
		rowsAffected, err := foodsv.NewDeleteListFoodDataService().SeviceDeleteListFood(ifid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
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

}

func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}

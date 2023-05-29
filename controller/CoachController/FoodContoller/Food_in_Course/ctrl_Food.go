package foodcontroller

import (
	"backEndGo/models"
	foodincourse "backEndGo/service/CoachService/FoodSV/Food_in_Course"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var foodSV = foodincourse.NewFoodDataService()
var insertFoodDataService = foodincourse.NewInsertFoodDataService()
var updateFoodDataService = foodincourse.NewUpdateFoodDataService()
var deleteFoodDataService = foodincourse.NewDeleteFoodDataService()
var modelsFood = models.Food{}

func NewFoodController(router *gin.Engine) {
	food := router.Group("/food")
	{
		food.GET("", getFood)
		food.POST("/dayID/:did", insertFood)
		food.PUT("/foodID/:fid", updateFood)
		food.DELETE("/foodID/:fid", deleteFood)

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
func insertFood(ctx *gin.Context) {
	dayID := ctx.Param("did")

	did, errs := strconv.Atoi(dayID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsFood)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := insertFoodDataService.SeviceInsertFood(did, &modelsFood)
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

func updateFood(ctx *gin.Context) {
	foodID := ctx.Param("fid")

	fid, errs := strconv.Atoi(foodID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsFood)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	} else {
		rowsAffected, err := updateFoodDataService.ServiceUpdateFood(fid, &modelsFood)
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
func deleteFood(ctx *gin.Context) {
	foodID := ctx.Param("fid")

	fid, errs := strconv.Atoi(foodID)

	if errs != nil {
	} else {
		rowsAffected, err := deleteFoodDataService.ServiceDeleteFood(fid)
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

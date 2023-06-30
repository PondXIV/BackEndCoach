package dayscontroller

import (
	"backEndGo/models"
	daysv "backEndGo/service/CoachService/DaySV"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var modelsDay = models.DayOfCouse{}
var updaterDayDataService = daysv.NewUpdateDayDataService()
var deleteDayDataService = daysv.NewDeleteDayDataService()
var getDayDataService = daysv.NewDayDataService()

func NewDayController(router *gin.Engine) {
	day := router.Group("/day")
	{
		day.GET("", getDay)
		day.PUT("/dayID/:did", updateDay)
		day.DELETE("/dayID/:did", deleteDay)
	}

}
func getDay(ctx *gin.Context) {

	qdid := ctx.Query("did")
	qcoID := ctx.Query("coID")
	qsequence := ctx.Query("sequence")

	did, err := strconv.Atoi(qdid)
	coID, err := strconv.Atoi(qcoID)
	sequence, err := strconv.Atoi(qsequence)

	days, err := getDayDataService.SeviceGetDay(did, coID, sequence)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, days)

}

func updateDay(ctx *gin.Context) {
	dayID := ctx.Param("did")

	did, errs := strconv.Atoi(dayID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsDay)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	} else {
		rowsAffected, err := updaterDayDataService.ServiceUpdateDay(did, &modelsDay)
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

func deleteDay(ctx *gin.Context) {

	dayID := ctx.Param("did")

	did, errs := strconv.Atoi(dayID)

	if errs != nil {
		fmt.Print(http.StatusBadRequest)
	} else {
		rowsAffected, err := deleteDayDataService.SeviceDeleteDay(did)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})

		} else {
			if rowsAffected > 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code":   "200",
					"result": strconv.Itoa(int(rowsAffected)),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":   "200",
					"result": strconv.Itoa(int(rowsAffected)),
				})
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

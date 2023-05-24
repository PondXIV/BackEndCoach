package dayscontroller

import (
	daysv "backEndGo/service/CoachService/DaySV"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDayController(router *gin.Engine) {
	day := router.Group("/day")
	{
		day.GET("", getDay)
	}

}
func getDay(ctx *gin.Context) {

	qdid := ctx.Query("did")
	qcoID := ctx.Query("coID")
	qsequence := ctx.Query("sequence")

	did, err := strconv.Atoi(qdid)
	coID, err := strconv.Atoi(qcoID)
	sequence, err := strconv.Atoi(qsequence)

	days, err := daysv.NewDayDataService().SeviceGetDay(did, coID, sequence)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, days)

}

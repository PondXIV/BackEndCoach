package coachctl

import (
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var coachDateService = userservice.NewCoachByNameDataService()

func NewCoachController(router *gin.Engine) {
	coach := router.Group("/coach")
	{
		coach.GET("", GetCoach)

	}

}
func GetCoach(ctx *gin.Context) {
	qcid := ctx.Query("cid")
	name := ctx.Query("username")
	cid, err := strconv.Atoi(qcid)
	coach, err := coachDateService.ServiceGetNameCoachs(cid, name)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, coach)
}

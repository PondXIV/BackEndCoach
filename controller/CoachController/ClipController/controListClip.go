package clipcontroller

import (
	clipsv "backEndGo/service/CoachService/ClipSV"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewListClipController(router *gin.Engine) {
	listClip := router.Group("/listClip")
	{
		listClip.GET("/:cid", getListClipByID)

	}

}
func getListClipByID(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, err := strconv.Atoi(coachID)
	clips, err := clipsv.NewListClipDataService().SeviceGetListClipByIDCoach(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, clips)

}

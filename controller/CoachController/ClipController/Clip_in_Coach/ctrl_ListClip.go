package clipcontroller

import (
	"backEndGo/models"
	clipsv "backEndGo/service/CoachService/ClipSV/Clip_in_Coach"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var modelsListClip = models.ListClip{}
var insertListClipDataService = clipsv.NewInsertListClipDataService()
var showListClip = clipsv.NewListClipDataService()

func NewListClipController(router *gin.Engine) {
	listClip := router.Group("/listClip")
	{
		listClip.GET("", getListClip)
		listClip.POST("/coachID/:cid", insertListClip)
		listClip.PUT("/clipID/:icpID", updateListClip)
		listClip.DELETE("/clipID/:icpID", deleteListClip)
	}

}
func getListClip(ctx *gin.Context) {
	qicpID := ctx.Query("icpID")
	qcid := ctx.Query("cid")
	qname := ctx.Query("name")

	icpID, err := strconv.Atoi(qicpID)
	cid, err := strconv.Atoi(qcid)

	clips, err := showListClip.SeviceGetListClip(icpID, cid, qname)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, clips)

}
func insertListClip(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, errs := strconv.Atoi(coachID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsListClip)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			error400(ctx)
		}
		// else {
		// 	error500(ctx)
		// }
	}
	rowsAffected, err := insertListClipDataService.SeviceInsertListClip(cid, &modelsListClip)
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

func updateListClip(ctx *gin.Context) {
	clipID := ctx.Param("icpID")

	icpID, errs := strconv.Atoi(clipID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsListClip)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			error400(ctx)
		}
		//else {
		// 	error500(ctx)
		// }
	} else {
		rowsAffected, err := clipsv.NewUpdateListClipDataService().ServiceUpdateListClip(icpID, &modelsListClip)
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
func deleteListClip(ctx *gin.Context) {
	clipID := ctx.Param("icpID")

	icpID, errs := strconv.Atoi(clipID)

	if errs != nil {
		fmt.Print(http.StatusBadRequest)

		if http.StatusBadRequest == 400 {
			error400(ctx)
		}
		//else {
		// 	error500(ctx)
		// }
	} else {
		rowsAffected, err := clipsv.NewDeleteListClipDataService().SeviceDeleteListClip(icpID)
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

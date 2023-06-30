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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := insertListClipDataService.SeviceInsertListClip(cid, &modelsListClip)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}
	} else {
		if rowsAffected == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   "200",
				"result": strconv.Itoa(int(rowsAffected)),
			})

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})

	} else {
		rowsAffected, err := clipsv.NewUpdateListClipDataService().ServiceUpdateListClip(icpID, &modelsListClip)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})

		} else {
			if rowsAffected == 1 {
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
func deleteListClip(ctx *gin.Context) {
	clipID := ctx.Param("icpID")

	icpID, errs := strconv.Atoi(clipID)

	if errs != nil {
		fmt.Print(http.StatusBadRequest)
	} else {
		rowsAffected, err := clipsv.NewDeleteListClipDataService().SeviceDeleteListClip(icpID)
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

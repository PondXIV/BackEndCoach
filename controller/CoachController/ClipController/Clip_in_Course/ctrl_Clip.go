package clipincourse

import (
	"backEndGo/models"
	clipincourse "backEndGo/service/CoachService/ClipSV/Clip_in_Course"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var clipSV = clipincourse.NewClipDataService()
var insertClipDataService = clipincourse.NewInsertClipDataService()
var updateClipDataService = clipincourse.NewUpdateClipDataService()
var deleteClipDataService = clipincourse.NewDeleteClipDataService()
var modelsClip = models.Clip{}

func NewClipController(router *gin.Engine) {
	clip := router.Group("/clip")
	{
		clip.GET("", getClip)
		clip.POST("/dayID/:did", insertClip)
		clip.PUT("/clipID/:cpID", updateClip)
		clip.DELETE("/clipID/:cpID", deleteClip)
	}
}
func getClip(ctx *gin.Context) {
	qcpID := ctx.Query("cpID")
	qicpid := ctx.Query("icpid")
	qdid := ctx.Query("did")

	cpID, err := strconv.Atoi(qcpID)
	icpid, err := strconv.Atoi(qicpid)
	did, err := strconv.Atoi(qdid)

	clips, err := clipSV.GetClip(cpID, icpid, did)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, clips)

}
func insertClip(ctx *gin.Context) {
	dayID := ctx.Param("did")

	did, errs := strconv.Atoi(dayID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsClip)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)
	}
	rowsAffected, err := insertClipDataService.SeviceInsertClip(did, &modelsClip)
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

func updateClip(ctx *gin.Context) {
	clipID := ctx.Param("cpID")

	cpID, errs := strconv.Atoi(clipID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsClip)
	// fmt.Printf("%v", cus)
	if err != nil {
		fmt.Print(http.StatusBadRequest)

	} else {
		rowsAffected, err := updateClipDataService.ServiceUpdateClip(cpID, &modelsClip)
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
				outputSoon(ctx)
			}
		}

	}

}

func deleteClip(ctx *gin.Context) {
	clipID := ctx.Param("cpID")

	cpID, errs := strconv.Atoi(clipID)

	if errs != nil {
		fmt.Print(http.StatusBadRequest)
	} else {
		rowsAffected, err := deleteClipDataService.SeviceDeleteClip(cpID)
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

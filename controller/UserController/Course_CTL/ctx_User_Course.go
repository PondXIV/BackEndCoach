package coursectl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mycourseService = userservice.NewMyCourseDataService()

func NewCourseController(router *gin.Engine) {
	courses := router.Group("/courses")
	{
		courses.PUT("/clip/:cpID", UpdateStatus)
		courses.GET("", GetCourseByUid)
	}

}

func GetCourseByUid(ctx *gin.Context) {
	cusID := ctx.Query("uid")
	uid, err := strconv.Atoi(cusID)
	course, err := mycourseService.ServiceGetMycourse(uid)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, course)
}
func UpdateStatus(ctx *gin.Context) {
	clipID := ctx.Param("cpID")
	cpID, _ := strconv.Atoi(clipID)
	jsonDto := models.Clip{}
	err := ctx.ShouldBindJSON(jsonDto)
	fmt.Printf("ID3", "%f", cpID)
	rowsAffected, err := userservice.NewClipUpdateStatusDataService().ServiceUpdateStatus(cpID, jsonDto.Status)
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
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}

package coursecontroller

import (
	//"backEndGo/models"

	coursesv "backEndGo/service/CoachService/CourseSV"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var courseDateService = coursesv.NewCourseDataService()

func NewCourseController(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.GET("/getCourseByIDCoach/:cid", getCourseByID)

	}

}

func getCourseByID(ctx *gin.Context) {
	// jsonDto := coachdto.IDCoachDTO{}
	// err := ctx.ShouldBindJSON(&jsonDto)
	coachID := ctx.Param("cid")
	//fmt.Printf("%v", coachID)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"mass": coachID,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	cid, err := strconv.Atoi(coachID)
	course, err := courseDateService.ServiceGetCourseByIDCoach(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}

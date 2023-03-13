package coursecontroller

import (
	//"backEndGo/models"
	"backEndGo/models"
	coursesv "backEndGo/service/CoachService/CourseSV"
	"fmt"
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
	for _, v := range *course {
		fmt.Printf("%v\n", v)
		ctx.JSON(http.StatusOK, v)
	}

	if cid > 0 {
		ctx.JSON(http.StatusOK, models.Course{})

	}

}

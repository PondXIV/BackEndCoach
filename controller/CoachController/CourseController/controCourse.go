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
		course.GET("/getCourseByName/:name", GetCousehByName)
		course.GET("/getCourseByCoID/:coID", GetCousehByCoID)

	}

}

func getCourseByID(ctx *gin.Context) {
	// jsonDto := coachdto.IDCoachDTO{}
	// err := ctx.ShouldBindJSON(&jsonDto)
	coachID := ctx.Param("cid")

	cid, err := strconv.Atoi(coachID)
	course, err := courseDateService.ServiceGetCourseByIDCoach(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func GetCousehByCoID(ctx *gin.Context) {
	// jsonDto := coachdto.IDCoachDTO{}
	// err := ctx.ShouldBindJSON(&jsonDto)
	courseID := ctx.Param("coID")

	coID, err := strconv.Atoi(courseID)
	course, err := courseDateService.SeviceGetCourseByCoID(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func GetCousehByName(ctx *gin.Context) {
	name := ctx.Param("name")

	course, err := courseDateService.SeviceGetCourseByName(name)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}

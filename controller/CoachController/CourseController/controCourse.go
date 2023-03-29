package coursecontroller

import (
	//"backEndGo/models"

	coachdto "backEndGo/DTO/CoachDTO"
	"backEndGo/models"
	coursesv "backEndGo/service/CoachService/CourseSV"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var courseDateService = coursesv.NewCourseDataService()
var updatecourseDateService = coursesv.NewUpdateCourseDataService()
var insertCourseDataService = coursesv.NewInsertCourseDataService()
var modelsCourse = models.Course{}

func NewCourseController(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.GET("/getCourseByIDCoach/:cid", getCourseByID)
		course.GET("/getCourseByName/:name", GetCousehByName)
		course.GET("/getCourseByCoID/:coID", GetCousehByCoID)
		course.PUT("/updateStatusCourse", updateStatusCourse)
		course.PUT("/updateCourse", updateCourse)
		course.POST("/insertCourse", insertCourse)
	}

}

func getCourseByID(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, err := strconv.Atoi(coachID)
	course, err := courseDateService.ServiceGetCourseByIDCoach(cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func updateStatusCourse(ctx *gin.Context) {
	jsonDto := coachdto.UpdateStatusCoachDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.Status)

	if err != nil {
		panic(err)
	}
	// // process
	rowsAffected := updatecourseDateService.ServiceUpdateStatusCourse(jsonDto.CoID, jsonDto.Status)
	if rowsAffected == 1 {
		ctx.JSON(http.StatusOK, models.Course{Status: jsonDto.Status})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "191",
		})
	}
}
func updateCourse(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	// // process
	rowsAffected := updatecourseDateService.ServiceUpdateCourse(&modelsCourse)

	if rowsAffected == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "1",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "0",
		})
	}
}
func insertCourse(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	rowsAffected := insertCourseDataService.ServiceInsertCourse(&modelsCourse)
	if rowsAffected == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "1",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"rowsAffected": "0",
		})
	}
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

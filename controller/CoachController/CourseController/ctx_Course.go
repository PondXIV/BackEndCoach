package coursecontroller

import (
	//"backEndGo/models"

	"backEndGo/models"
	coachdto "backEndGo/models/request"
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
		course.GET("", getCourse)
		// course.PUT("/updateStatusCourse", updateStatusCourse)
		course.PUT("/courseID/:coID", updateCourse)
		course.POST("/coachID/:cid", insertCourse)

	}

}
func getCourse(ctx *gin.Context) {
	qcoid := ctx.Query("coID")
	qcid := ctx.Query("cid")
	qname := ctx.Query("name")

	coid, err := strconv.Atoi(qcoid)
	cid, err := strconv.Atoi(qcid)

	// If params not exist string = "", int = 0

	course, err := courseDateService.SeviceGetCourse(coid, cid, qname)
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
	courseID := ctx.Param("coID")

	coID, errs := strconv.Atoi(courseID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		error400(ctx)
	}
	// // process
	rowsAffected, err := updatecourseDateService.ServiceUpdateCourse(coID, &modelsCourse)

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
func insertCourse(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, errs := strconv.Atoi(coachID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		error400(ctx)
	}
	rowsAffected, err := insertCourseDataService.ServiceInsertCourse(cid, &modelsCourse)
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

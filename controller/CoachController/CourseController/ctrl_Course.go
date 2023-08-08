package coursecontroller

import (
	//"backEndGo/models"

	"backEndGo/models"
	coachdto "backEndGo/models/request"
	coursesv "backEndGo/service/CoachService/CourseSV"
	daysv "backEndGo/service/CoachService/DaySV"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var courseDateService = coursesv.NewCourseDataService()
var updatecourseDateService = coursesv.NewUpdateCourseDataService()
var insertCourseDataService = coursesv.NewInsertCourseDataService()
var deleteCourseDataService = coursesv.NewDeleteCourseDataService()
var getDayDataService = daysv.NewDayDataService()
var modelsCourse = models.Course{}

func NewCourseController(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.GET("", getCourse)
		course.GET("/opensell", getCourseSell)
		course.GET("/sell/:coID", getCoursebidnotnull)
		// course.PUT("/updateStatusCourse", updateStatusCourse)
		course.PUT("/courseID/:coID", updateCourse)
		course.POST("/coachID/:cid", insertCourse)
		course.DELETE("/courseID/:coID", deleteCourse)
		course.PUT("/expiration/:coID", updateExpiration)
		course.GET("/progess", GetProgessbar)
		course.GET("/amount", Getamountclip)

	}

}
func getCourseSell(ctx *gin.Context) {
	qcoid := ctx.Query("coID")
	qcid := ctx.Query("cid")
	qname := ctx.Query("name")

	coid, err := strconv.Atoi(qcoid)
	cid, err := strconv.Atoi(qcid)

	// If params not exist string = "", int = 0

	course, err := courseDateService.SeviceGetCourseSell(coid, cid, qname)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

}
func Getamountclip(ctx *gin.Context) {
	courseID := ctx.Query("coID")

	coID, err := strconv.Atoi(courseID)
	course := getDayDataService.SeviceGetAmoutclip(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"amount": course,
	})

}
func GetProgessbar(ctx *gin.Context) {
	courseID := ctx.Query("coID")

	coID, err := strconv.Atoi(courseID)
	course := getDayDataService.SeviceShowProgess(coID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"percent": course,
	})

}
func updateExpiration(ctx *gin.Context) {
	courseID := ctx.Param("coID")

	coID, errs := strconv.Atoi(courseID)
	if errs != nil {
		panic(errs)
	}
	jsonDto := models.Course{}
	err := ctx.ShouldBindJSON(&jsonDto)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := updatecourseDateService.ServiceUpdateExpiration(coID, jsonDto.Days)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}

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
func getCoursebidnotnull(ctx *gin.Context) {
	courseID := ctx.Param("coID")
	coid, err := strconv.Atoi(courseID)

	// If params not exist string = "", int = 0

	course, err := courseDateService.SeviceGetCourseByCoID(coid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, course)

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
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}

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
func updateCourse(ctx *gin.Context) {
	courseID := ctx.Param("coID")

	coID, errs := strconv.Atoi(courseID)
	if errs != nil {
		panic(errs)
	}

	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	// // process
	rowsAffected, err := updatecourseDateService.ServiceUpdateCourse(coID, &modelsCourse)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}

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
func insertCourse(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, errs := strconv.Atoi(coachID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsCourse)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := insertCourseDataService.ServiceInsertCourse(cid, &modelsCourse)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}

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
func deleteCourse(ctx *gin.Context) {
	courseID := ctx.Param("coID")

	coID, errs := strconv.Atoi(courseID)

	if errs != nil {
	} else {
		rowsAffected, err := deleteCourseDataService.SeviceDeleteCourse(coID)
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

package coursecontroller

import (
	coachdto "backEndGo/DTO/CoachDTO"
	coursesv "backEndGo/service/CoachService/CourseSV"
	
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var courseDateService = coursesv.NewCourseDataService()

func NewCourseController(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.POST("/getCourseByIDCoach", getCourseByID)

	}

}

func getCourseByID(ctx *gin.Context) {
	jsonDto := coachdto.IDCoachDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf("%v", jsonDto.Cid)
	if err != nil {
		panic(err)
	}

	course, err := courseDateService.ServiceGetCourseByIDCoach(jsonDto.Cid)
	if err != nil {
		panic(err)
	}
	for _, v := range *course {
		fmt.Printf("%v\n", v)
		ctx.JSON(http.StatusOK, v)
	}

	// if course.CoID > 0 {
	// 	ctx.JSON(http.StatusOK, models.Course{})

	// }

}

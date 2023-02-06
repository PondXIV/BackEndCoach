package controller

import (
	coachdto "backEndGo/DTO/CoachDTO"
	coachservice "backEndGo/service/CoachService"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var courseDateService = coachservice.NewCourseDataService()

func NewCourseController(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.POST("/", getCourseByID)
		course.GET("/", demoHelloo)
	}

}
func demoHelloo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ฟ่ดนะจ๊ะ",
	})
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

package usercontroller

import (
	userdto "backEndGo/DTO/UserDto"
	userservice "backEndGo/service/userService"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

var nameCoachDateService = userservice.NewCoachByNameDataService()

func NewCourseController(router *gin.Engine) {
	course := router.Group("/user2")
	{
		course.POST("/getCoachByName", GetCoachByName)

	}

}

func GetCoachByName(ctx *gin.Context) {
	jsonDto := userdto.UsernameCoachdto{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf("%v", jsonDto.Username)
	if err != nil {
		panic(err)
	}

	coachs, err := nameCoachDateService.ServiceGetNameCoach(jsonDto.Username)
	if err != nil {
		panic(err)
	}
	for _, v := range *coachs {
		fmt.Printf("%v\n", v)
		ctx.JSON(http.StatusOK, v)
	}

	// if course.CoID > 0 {
	// 	ctx.JSON(http.StatusOK, models.Course{})

	// }

}

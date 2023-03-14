package usercontroller

import (
	userservice "backEndGo/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

var nameCoachDateService = userservice.NewCoachByNameDataService()

func NewCourseController(router *gin.Engine) {
	nameCoach := router.Group("/user2")
	{
		nameCoach.GET("/getCoachByName/:name", GetCoachByName)

	}

}

func GetCoachByName(ctx *gin.Context) {
	name := ctx.Param("name")
	//jsonDto := userdto.UsernameCoachdto{}
	//err := ctx.ShouldBindJSON(&jsonDto)
	//fmt.Printf(name, jsonDto.Username)
	/*if err != nil {
		panic(err)
	}*/

	coachs, err := nameCoachDateService.ServiceGetNameCoach(name)
	if err != nil {
		panic(err)
	}
	// for _, v := range *coachs {
	// 	fmt.Printf("%v\n", v)
	ctx.JSON(http.StatusOK, coachs)
	// }

	// if course.CoID > 0 {
	// 	ctx.JSON(http.StatusOK, models.Course{})

	// }

}

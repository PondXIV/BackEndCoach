package coursectl

import (
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mycourseService = userservice.NewMyCourseDataService()

func NewCourseController(router *gin.Engine) {
	courses := router.Group("/courses")
	{

		courses.GET("/customer", GetCourseByUid)
	}

}

func GetCourseByUid(ctx *gin.Context) {
	cusID := ctx.Query("uid")
	uid, err := strconv.Atoi(cusID)
	course, err := mycourseService.ServiceGetMycourse(uid)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, course)
}

package controller

import (
	dto "backEndGo/DTO"
	"backEndGo/models"
	"backEndGo/service"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

var userDateService = service.NewUserDataService()
var modelsCus = models.Customer{}
var modelsCoach = models.Coach{}

func NewUserController(router *gin.Engine) {
	ping := router.Group("/user")
	{
		ping.POST("/login", loginPostBody)
		ping.POST("/registerCus", registerCus)
		ping.POST("/registerCoach", registerCoach)
	}

}
func loginPostBody(ctx *gin.Context) {
	// input json
	jsonDto := dto.LoginDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.Email, jsonDto.Password)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, err)
		panic(err)
	}
	// process
	coach, cus, err := userDateService.ServiceLogin(jsonDto.Email, jsonDto.Password, jsonDto.Type)
	if err != nil {
		panic(err)
	}

	if coach.Cid > 0 {
		ctx.JSON(http.StatusOK, models.Coach{Cid: coach.Cid})
	} else if cus.Uid > 0 {
		ctx.JSON(http.StatusOK, models.Customer{Uid: cus.Uid})
	} else if cus.Uid == 0 {
		if jsonDto.Type == 1 {
			ctx.JSON(http.StatusOK, models.Customer{})
		} else {
			ctx.JSON(http.StatusOK, models.Coach{})
		}

	} else if coach.Cid == 0 {
		if jsonDto.Type == 0 {
			ctx.JSON(http.StatusOK, models.Coach{})
		} else {
			ctx.JSON(http.StatusOK, models.Customer{})
		}
	}
	// else {
	// 	ctx.JSON(http.StatusBadRequest, err)
	// }
	// else if len(*cus) == 0 {
	// 	ctx.JSON(http.StatusOK, models.Customer{})
	// }
	// println("=================")
}
func registerCus(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCus)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCus(&modelsCus)

	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, models.Customer{Uid: modelsCus.Uid})
	} else {
		ctx.JSON(http.StatusOK, models.Customer{Uid: modelsCus.Uid})
	}
}
func registerCoach(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCoach)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCoach(&modelsCoach)
	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, models.Coach{Cid: modelsCoach.Cid})
	} else {
		ctx.JSON(http.StatusOK, models.Coach{Cid: modelsCoach.Cid})
	}
}

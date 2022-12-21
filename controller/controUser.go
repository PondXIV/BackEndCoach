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
		ctx.JSON(http.StatusBadRequest, err)
	}
	// process
	coach, cus, err := userDateService.ServiceLogin(jsonDto.Email, jsonDto.Password, jsonDto.Type)

	// output json
	// println("===============")
	if err != nil {
		panic(err)
	}

	if len(*coach) == 1 {
		ctx.JSON(http.StatusOK, coach)
	} else if len(*cus) == 1 {
		ctx.JSON(http.StatusOK, cus)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	// println("=================")
}
func registerCus(ctx *gin.Context) {
	cus := models.Customer{}
	err := ctx.ShouldBindJSON(&cus)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCus(&cus)

	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "Register Success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "Register failed",
		})
	}
}
func registerCoach(ctx *gin.Context) {
	coach := models.Coach{}
	err := ctx.ShouldBindJSON(&coach)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCoach(&coach)
	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "Register Success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "Register failed",
		})
	}
}

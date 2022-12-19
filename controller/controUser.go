package controller

import (
	dto "backEndGo/DTO"
	"backEndGo/service"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var showDateService = service.NewShowDataService()

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
	coach, cus, err := showDateService.Login(jsonDto.Email, jsonDto.Password, jsonDto.Type)

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
	jsonDto := dto.RegisterDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.AliasName, jsonDto.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
func registerCoach(ctx *gin.Context) {

}
